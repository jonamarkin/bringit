<script setup lang="ts">
import { ArrowUpRight, Check, Package, Plus, Trash2 } from '@lucide/vue'
import type { EventItem, EventSnapshot, Guest, ItemClaim } from '~/types/bringit'

const api = useApi()
const route = useRoute()
const { showToast } = useToast()

const shareCode = computed(() => String(route.params.shareCode || ''))
const snapshot = ref<EventSnapshot | null>(null)
const filter = ref('All')
const guest = ref<Partial<Guest>>({})
const guestName = ref('')
const guestEmail = ref('')
const rsvpStatus = ref('attending')
const partySize = ref(1)
const stream = ref<EventSource | null>(null)

const tokenKey = computed(() => `bringit_guest_token_${shareCode.value}`)
const guestIDKey = computed(() => `bringit_guest_id_${shareCode.value}`)
const guestNameKey = computed(() => `bringit_guest_name_${shareCode.value}`)

const categories = computed(() => {
  if (!snapshot.value) return ['All']
  return ['All', ...Array.from(new Set(snapshot.value.items.map((item) => item.category)))]
})

const filteredItems = computed(() => {
  if (!snapshot.value) return []
  if (filter.value === 'All') return snapshot.value.items
  return snapshot.value.items.filter((item) => item.category === filter.value)
})

const guestClaims = computed(() => {
  if (!snapshot.value || !guest.value.id) return []
  return snapshot.value.claims.filter((claim) => claim.guest_id === guest.value.id)
})

const progress = computed(() => {
  const stats = snapshot.value?.stats
  if (!stats) return 0
  return Math.min(100, Math.round((stats.items_claimed / Math.max(stats.items_needed, 1)) * 100))
})

const hasRSVP = computed(() => Boolean(guest.value.session_token))

const apiBase = computed(() => useRuntimeConfig().public.apiBase as string)

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

const loadSnapshot = async () => {
  const res = await api<{ snapshot: EventSnapshot }>(`/api/v1/public/events/${shareCode.value}`)
  snapshot.value = res.snapshot
}

const saveRSVP = async () => {
  try {
    const res = await api<{ guest: Guest }>(`/api/v1/public/events/${shareCode.value}/rsvp`, {
      method: 'POST',
      body: {
        session_token: localStorage.getItem(tokenKey.value) || '',
        name: guestName.value,
        email: guestEmail.value,
        rsvp_status: rsvpStatus.value,
        party_size: partySize.value,
      },
    })
    guest.value = res.guest
    localStorage.setItem(tokenKey.value, res.guest.session_token || '')
    localStorage.setItem(guestIDKey.value, res.guest.id)
    localStorage.setItem(guestNameKey.value, res.guest.name)
    showToast('RSVP saved')
    await loadSnapshot()
  } catch (error: any) {
    showToast(error.data?.error || error.message || 'Could not save RSVP')
  }
}

const claimItem = async (item: EventItem, event: Event) => {
  const form = new FormData(event.currentTarget as HTMLFormElement)
  try {
    await api(`/api/v1/public/events/${shareCode.value}/claims`, {
      method: 'POST',
      body: {
        session_token: localStorage.getItem(tokenKey.value) || '',
        item_id: item.id,
        quantity: Number(form.get('quantity') || 1),
        note: String(form.get('note') || ''),
      },
    })
    ;(event.currentTarget as HTMLFormElement).reset()
    showToast('Added to your list')
    await loadSnapshot()
  } catch (error: any) {
    showToast(error.data?.error || error.message || 'Could not claim item')
  }
}

const removeClaim = async (claim: ItemClaim) => {
  try {
    await api(`/api/v1/public/events/${shareCode.value}/claims/${claim.id}`, {
      method: 'DELETE',
      body: {
        session_token: localStorage.getItem(tokenKey.value) || '',
      },
    })
    showToast('Removed')
    await loadSnapshot()
  } catch (error: any) {
    showToast(error.data?.error || error.message || 'Could not remove item')
  }
}

const connectStream = () => {
  if (stream.value) stream.value.close()
  stream.value = new EventSource(`${apiBase.value}/api/v1/public/events/${shareCode.value}/stream`)
  stream.value.addEventListener('snapshot', (event) => {
    snapshot.value = JSON.parse(event.data)
  })
}

onMounted(async () => {
  guest.value = {
    id: localStorage.getItem(guestIDKey.value) || '',
    session_token: localStorage.getItem(tokenKey.value) || '',
  }
  guestName.value = localStorage.getItem(guestNameKey.value) || ''
  await loadSnapshot()
  connectStream()
})

onBeforeUnmount(() => {
  if (stream.value) stream.value.close()
})
</script>

<template>
  <div class="app-shell">
    <AppTopbar>
      <template #action>
        <NuxtLink class="secondary-button" to="/host">Host</NuxtLink>
      </template>
    </AppTopbar>

    <main v-if="snapshot" class="main">
      <section class="panel">
        <div class="event-hero">
          <span class="badge dark">You're invited</span>
          <h1 class="event-title">{{ snapshot.event.title }}</h1>
          <div class="event-meta">
            <span>{{ formatDate(snapshot.event.starts_at) }}</span>
            <span>{{ snapshot.event.location_name || 'Location pending' }}</span>
          </div>
        </div>
      </section>

      <section class="panel">
        <EventMetrics :stats="snapshot.stats" />
        <div class="progress-track" style="margin-top: 12px">
          <span class="progress-fill" :style="{ width: `${progress}%` }" />
        </div>
      </section>

      <section class="dashboard-grid">
        <section class="panel" id="rsvp">
          <div class="panel-header">
            <h2 class="panel-title">RSVP</h2>
            <span class="small muted">{{ hasRSVP ? 'Saved' : 'Open' }}</span>
          </div>

          <form class="form-grid" @submit.prevent="saveRSVP">
            <label class="field">
              <span class="label">Name</span>
              <Input v-model="guestName" class="input" required placeholder="Your name" />
            </label>
            <div class="two-col">
              <label class="field">
                <span class="label">Status</span>
                <select v-model="rsvpStatus" class="select">
                  <option value="attending">Attending</option>
                  <option value="maybe">Maybe</option>
                  <option value="not_attending">Not attending</option>
                </select>
              </label>
              <label class="field">
                <span class="label">Party size</span>
                <Input v-model.number="partySize" class="input" type="number" min="1" max="20" />
              </label>
            </div>
            <label class="field">
              <span class="label">Email</span>
              <Input v-model="guestEmail" class="input" type="email" placeholder="Optional" />
            </label>
            <Button class="primary-button" type="submit">
              <Check :size="16" />
              Save RSVP
            </Button>
          </form>
        </section>

        <section class="panel">
          <div class="panel-header">
            <h2 class="panel-title">What can you bring?</h2>
            <span class="small muted">{{ filteredItems.length }} shown</span>
          </div>

          <div class="tabs">
            <button
              v-for="category in categories"
              :key="category"
              class="tab"
              :class="{ active: category === filter }"
              @click="filter = category"
            >
              {{ category }}
            </button>
          </div>

          <div class="item-list">
            <article
              v-for="item in filteredItems"
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
              <form
                v-if="item.claimed_qty < item.needed_qty"
                class="form-grid"
                @submit.prevent="claimItem(item, $event)"
              >
                <div class="two-col">
                  <label class="field">
                    <span class="label">Qty</span>
                    <Input class="input" name="quantity" type="number" min="1" :max="item.needed_qty - item.claimed_qty" value="1" :disabled="!hasRSVP" />
                  </label>
                  <label class="field">
                    <span class="label">Note</span>
                    <Input class="input" name="note" placeholder="Optional" :disabled="!hasRSVP" />
                  </label>
                </div>
                <Button class="secondary-button" type="submit" :disabled="!hasRSVP">
                  <Plus :size="16" />
                  Claim
                </Button>
              </form>
            </article>
          </div>
        </section>
      </section>

      <section v-if="guestClaims.length" class="panel">
        <div class="panel-header">
          <h2 class="panel-title">Your list</h2>
          <span class="small muted">{{ guestClaims.length }} claims</span>
        </div>
        <div class="claim-list">
          <div v-for="claim in guestClaims" :key="claim.id" class="claim-row">
            <div>
              <strong>{{ claim.quantity }} {{ claim.item_name }}</strong>
              <div class="item-meta">{{ claim.note || 'Ready' }}</div>
            </div>
            <Button class="danger-button" @click="removeClaim(claim)">
              <Trash2 :size="15" />
              Remove
            </Button>
          </div>
        </div>
      </section>
    </main>

    <main v-else class="main">
      <section class="panel">
        <div class="empty">
          <Package :size="22" />
          <strong>Loading event</strong>
        </div>
      </section>
    </main>

    <div class="bottom-bar">
      <a class="primary-button" href="#rsvp">
        <ArrowUpRight :size="16" />
        {{ hasRSVP ? 'Bring something' : 'RSVP' }}
      </a>
    </div>

    <AppToast />
  </div>
</template>
