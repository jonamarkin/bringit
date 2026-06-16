<script setup lang="ts">
import { ArrowUpRight, CalendarDays, Copy, LogOut, MessageCircle, Package, Plus, PlusCircle } from '@lucide/vue'
import type { EventSnapshot, EventSummary, Notification } from '~/types/bringit'

const api = useApi()
const route = useRoute()
const router = useRouter()
const { user, loadUser, requestOTP, verifyOTP, logout } = useAuth()
const { showToast } = useToast()

const events = ref<EventSummary[]>([])
const snapshot = ref<EventSnapshot | null>(null)
const notifications = ref<Notification[]>([])
const shareUrl = ref('')
const loading = ref(true)
const authStep = ref<'email' | 'code'>('email')
const loginEmail = ref('')
const loginName = ref('')
const loginCode = ref('')

const progress = computed(() => {
  const stats = snapshot.value?.stats
  if (!stats) return 0
  return Math.min(100, Math.round((stats.items_claimed / Math.max(stats.items_needed, 1)) * 100))
})

const selectedEventId = computed(() => String(route.query.event || events.value[0]?.id || ''))

const formatDate = (value: string) => {
  if (!value) return ''
  return new Intl.DateTimeFormat(undefined, {
    weekday: 'short',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(new Date(value))
}

const loadEvents = async () => {
  const res = await api<{ events: EventSummary[] }>('/api/v1/events')
  events.value = res.events || []
  if (!events.value.length) {
    snapshot.value = null
    notifications.value = []
    return
  }
  if (!route.query.event) {
    await router.replace({ path: '/host', query: { event: events.value[0].id } })
    return
  }
  await loadEvent(selectedEventId.value)
}

const loadEvent = async (eventId: string) => {
  if (!eventId) return
  const res = await api<{ snapshot: EventSnapshot, share_url: string }>(`/api/v1/events/${eventId}`)
  snapshot.value = res.snapshot
  shareUrl.value = res.share_url
  const notificationRes = await api<{ notifications: Notification[] }>(`/api/v1/events/${eventId}/notifications`)
  notifications.value = notificationRes.notifications || []
}

const sendCode = async () => {
  try {
    const res = await requestOTP(loginEmail.value, loginName.value)
    if (res.dev_code) loginCode.value = res.dev_code
    authStep.value = 'code'
    showToast('Code sent')
  } catch (error: any) {
    showToast(error.data?.error || error.message || 'Could not send code')
  }
}

const verifyCode = async () => {
  try {
    await verifyOTP(loginEmail.value, loginName.value, loginCode.value)
    await loadEvents()
  } catch (error: any) {
    showToast(error.data?.error || error.message || 'Invalid code')
  }
}

const copyShareLink = async () => {
  await navigator.clipboard.writeText(shareUrl.value)
  showToast('Share link copied')
}

const markRead = async () => {
  if (!snapshot.value) return
  await api(`/api/v1/events/${snapshot.value.event.id}/notifications/mark-read`, { method: 'POST' })
  await loadEvent(snapshot.value.event.id)
  showToast('Notifications marked read')
}

const switchEvent = async (event: Event) => {
  const target = event.target as HTMLSelectElement
  await router.push({ path: '/host', query: { event: target.value } })
}

const addItem = async (event: Event) => {
  if (!snapshot.value) return
  const form = new FormData(event.currentTarget as HTMLFormElement)
  try {
    await api(`/api/v1/events/${snapshot.value.event.id}/items`, {
      method: 'POST',
      body: {
        name: String(form.get('name') || ''),
        category: String(form.get('category') || 'Other'),
        needed_qty: Number(form.get('needed_qty') || 1),
        unit: String(form.get('unit') || ''),
      },
    })
    ;(event.currentTarget as HTMLFormElement).reset()
    await loadEvent(snapshot.value.event.id)
    showToast('Item added')
  } catch (error: any) {
    showToast(error.data?.error || error.message || 'Could not add item')
  }
}

watch(() => route.query.event, async (eventId) => {
  if (user.value && eventId) {
    await loadEvent(String(eventId))
  }
})

onMounted(async () => {
  loading.value = true
  await loadUser()
  if (user.value) {
    await loadEvents()
  }
  loading.value = false
})
</script>

<template>
  <div class="app-shell">
    <AppTopbar>
      <template #action>
        <div class="button-row">
          <NuxtLink v-if="user" class="primary-button" to="/create">
            <PlusCircle :size="16" />
            New Event
          </NuxtLink>
          <Button v-if="user" variant="ghost" size="icon" title="Log out" @click="logout">
            <LogOut :size="17" />
          </Button>
        </div>
      </template>
    </AppTopbar>

    <main class="main">
      <section v-if="loading" class="panel">
        <div class="empty">
          <strong>Loading</strong>
        </div>
      </section>

      <section v-else-if="!user" class="dashboard-grid">
        <div class="panel">
          <div class="event-hero">
            <span class="badge dark">Organizer</span>
            <h1 class="event-title">BringIt</h1>
            <div class="event-meta">
              <span>Plan the hangout, share the link, watch the checklist fill in.</span>
            </div>
          </div>
        </div>

        <div class="panel">
          <div class="panel-header">
            <h2 class="panel-title">{{ authStep === 'email' ? 'Sign in' : 'Enter code' }}</h2>
            <span class="small muted">Host dashboard</span>
          </div>

          <form v-if="authStep === 'email'" class="form-grid" @submit.prevent="sendCode">
            <label class="field">
              <span class="label">Email</span>
              <Input v-model="loginEmail" class="input" type="email" required placeholder="you@example.com" />
            </label>
            <label class="field">
              <span class="label">Name</span>
              <Input v-model="loginName" class="input" placeholder="Your name" />
            </label>
            <Button class="primary-button" type="submit">
              <MessageCircle :size="16" />
              Send code
            </Button>
          </form>

          <form v-else class="form-grid" @submit.prevent="verifyCode">
            <label class="field">
              <span class="label">Code</span>
              <Input v-model="loginCode" class="input" inputmode="numeric" required placeholder="000000" />
            </label>
            <Button class="primary-button" type="submit">Continue</Button>
            <Button class="ghost-button" type="button" @click="authStep = 'email'">Use another email</Button>
          </form>
        </div>
      </section>

      <section v-else-if="!events.length" class="panel">
        <div class="empty">
          <strong>No events yet</strong>
          <span>Create the BBQ plan and share the guest link.</span>
          <NuxtLink class="primary-button" to="/create">
            <Plus :size="16" />
            New Event
          </NuxtLink>
        </div>
      </section>

      <template v-else-if="snapshot">
        <section class="panel">
          <div class="panel-header">
            <h2 class="panel-title">Notifications ({{ notifications.filter((item) => !item.is_read).length }})</h2>
            <button class="panel-link" @click="markRead">See all</button>
          </div>
          <div class="notification-grid">
            <article class="notice-card peach">
              <span class="corner-icon"><ArrowUpRight :size="12" /></span>
              <div class="notice-kicker"><MessageCircle :size="15" /> RSVP</div>
              <h3 class="notice-title">{{ snapshot.stats.attending }} attending</h3>
              <p class="notice-text">{{ snapshot.stats.maybe }} maybe. Party size is {{ snapshot.stats.party_size_total || snapshot.stats.attending }}.</p>
            </article>
            <article class="notice-card blue">
              <span class="corner-icon"><ArrowUpRight :size="12" /></span>
              <div class="notice-kicker"><CalendarDays :size="15" /> Event</div>
              <h3 class="notice-title">{{ snapshot.event.title }}</h3>
              <p class="notice-text">{{ formatDate(snapshot.event.starts_at) }} at {{ snapshot.event.location_name || 'location pending' }}.</p>
            </article>
            <article class="notice-card mint">
              <span class="corner-icon"><ArrowUpRight :size="12" /></span>
              <div class="notice-kicker"><Package :size="15" /> Items</div>
              <h3 class="notice-title">{{ snapshot.stats.items_remaining }} remaining</h3>
              <p class="notice-text">{{ notifications[0]?.message || 'No guest activity yet.' }}</p>
            </article>
          </div>
        </section>

        <section class="dashboard-grid">
          <section class="panel stack">
            <div class="panel-header">
              <h2 class="panel-title">Event</h2>
              <select
                class="select"
                :value="snapshot.event.id"
                @change="switchEvent"
              >
                <option v-for="event in events" :key="event.id" :value="event.id">
                  {{ event.title }}
                </option>
              </select>
            </div>

            <div class="event-hero">
              <span class="badge dark">Live plan</span>
              <h1 class="event-title">{{ snapshot.event.title }}</h1>
              <div class="event-meta">
                <span>{{ formatDate(snapshot.event.starts_at) }}</span>
                <span>{{ snapshot.event.location_name || 'Location pending' }}</span>
              </div>
            </div>

            <EventMetrics :stats="snapshot.stats" />
            <div class="progress-track">
              <span class="progress-fill" :style="{ width: `${progress}%` }" />
            </div>

            <div class="button-row">
              <Button class="primary-button" @click="copyShareLink">
                <Copy :size="16" />
                Share link
              </Button>
              <NuxtLink class="secondary-button" :to="`/event/${snapshot.event.share_code}`">
                <ArrowUpRight :size="16" />
                Open guest view
              </NuxtLink>
            </div>

            <form class="form-grid" @submit.prevent="addItem">
              <div class="two-col">
                <label class="field">
                  <span class="label">Item</span>
                  <Input class="input" name="name" required placeholder="Plantain" />
                </label>
                <label class="field">
                  <span class="label">Category</span>
                  <Input class="input" name="category" placeholder="Sides" />
                </label>
              </div>
              <div class="two-col">
                <label class="field">
                  <span class="label">Qty</span>
                  <Input class="input" name="needed_qty" type="number" min="1" value="1" />
                </label>
                <label class="field">
                  <span class="label">Unit</span>
                  <Input class="input" name="unit" placeholder="packs" />
                </label>
              </div>
              <Button class="secondary-button" type="submit">
                <Plus :size="16" />
                Add item
              </Button>
            </form>
          </section>

          <section class="stack">
            <div class="panel">
              <div class="panel-header">
                <h2 class="panel-title">Checklist</h2>
                <span class="small muted">{{ snapshot.items.length }} items</span>
              </div>
              <div class="item-list">
                <article
                  v-for="item in snapshot.items"
                  :key="item.id"
                  class="item-card"
                  :class="{ done: item.claimed_qty >= item.needed_qty }"
                >
                  <div class="item-top">
                    <div>
                      <div class="item-name">{{ item.name }}</div>
                      <div class="item-meta">{{ item.category }} &middot; {{ item.claimed_qty }}/{{ item.needed_qty }} {{ item.unit || 'needed' }}</div>
                    </div>
                    <span class="badge" :class="{ soft: item.claimed_qty < item.needed_qty }">
                      {{ item.claimed_qty >= item.needed_qty ? 'Covered' : `${item.needed_qty - item.claimed_qty} left` }}
                    </span>
                  </div>
                  <div class="progress-track">
                    <span class="progress-fill" :style="{ width: `${Math.min(100, (item.claimed_qty / Math.max(item.needed_qty, 1)) * 100)}%` }" />
                  </div>
                </article>
              </div>
            </div>

            <div class="panel">
              <div class="panel-header">
                <h2 class="panel-title">Guests</h2>
                <span class="small muted">{{ snapshot.guests.length }} responses</span>
              </div>
              <div v-if="snapshot.guests.length" class="guest-list">
                <div v-for="guest in snapshot.guests" :key="guest.id" class="guest-row">
                  <div>
                    <strong>{{ guest.name }}</strong>
                    <div class="guest-text">{{ guest.rsvp_status }} &middot; {{ guest.party_size }} {{ guest.party_size === 1 ? 'person' : 'people' }}</div>
                  </div>
                  <span class="badge soft">{{ guest.email || 'guest' }}</span>
                </div>
              </div>
              <div v-else class="empty">
                <strong>No guests yet</strong>
              </div>
            </div>

            <div class="panel">
              <div class="panel-header">
                <h2 class="panel-title">Activity</h2>
                <span class="small muted">Latest</span>
              </div>
              <div v-if="snapshot.activities.length" class="activity-list">
                <div v-for="activity in snapshot.activities" :key="activity.id" class="activity-row">
                  <div>
                    <strong>{{ activity.actor_name }}</strong>
                    <div class="activity-text">{{ activity.message }}</div>
                  </div>
                  <span class="small muted">{{ new Date(activity.created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) }}</span>
                </div>
              </div>
              <div v-else class="empty">
                <strong>No activity yet</strong>
              </div>
            </div>
          </section>
        </section>
      </template>
    </main>

    <AppToast />
  </div>
</template>
