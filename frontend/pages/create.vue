<script setup lang="ts">
import { Plus, PlusCircle } from '@lucide/vue'

const api = useApi()
const router = useRouter()
const { user, loadUser } = useAuth()
const { showToast } = useToast()

const title = ref('Lulea BBQ Hangout')
const startsAt = ref('')
const locationName = ref('London Sky Garden')
const description = ref('Food, music, and a relaxed BBQ hangout with friends.')
const itemsText = ref(`Grill meat | Grill | 4 | packs
Jollof rice | Sides | 2 | trays
Soft drinks | Drinks | 8 | bottles
Water | Drinks | 12 | bottles
Charcoal | Supplies | 2 | bags
Disposable plates | Supplies | 1 | pack`)

const parseItems = () => {
  return itemsText.value
    .split('\n')
    .map((line, index) => {
      const [name, category, qty, unit] = line.split('|').map((part) => part?.trim())
      return {
        name,
        category: category || 'Other',
        needed_qty: Number(qty || 1),
        unit: unit || '',
        sort_order: index,
      }
    })
    .filter((item) => item.name)
}

const createEvent = async () => {
  try {
    const created = await api<{ event: { id: string } }>('/api/v1/events', {
      method: 'POST',
      body: {
        title: title.value,
        starts_at: new Date(startsAt.value).toISOString(),
        location_name: locationName.value,
        description: description.value,
        timezone: Intl.DateTimeFormat().resolvedOptions().timeZone || 'Europe/Berlin',
        items: parseItems(),
      },
    })
    showToast('Event created')
    await router.push({ path: '/host', query: { event: created.event.id } })
  } catch (error: any) {
    showToast(error.data?.error || error.message || 'Could not create event')
  }
}

onMounted(async () => {
  await loadUser()
  if (!user.value) {
    await router.push('/host')
    return
  }
  const future = new Date(Date.now() + 3 * 24 * 60 * 60 * 1000)
  future.setMinutes(future.getMinutes() - future.getTimezoneOffset())
  startsAt.value = future.toISOString().slice(0, 16)
})
</script>

<template>
  <div class="app-shell">
    <AppTopbar>
      <template #action>
        <NuxtLink class="secondary-button" to="/host">Dashboard</NuxtLink>
      </template>
    </AppTopbar>

    <main class="main">
      <section class="panel">
        <div class="panel-header">
          <h2 class="panel-title">New event</h2>
          <span class="small muted">Checklist builder</span>
        </div>

        <form class="form-grid" @submit.prevent="createEvent">
          <label class="field">
            <span class="label">Title</span>
            <Input v-model="title" class="input" required />
          </label>

          <div class="two-col">
            <label class="field">
              <span class="label">Date and time</span>
              <Input v-model="startsAt" class="input" type="datetime-local" required />
            </label>
            <label class="field">
              <span class="label">Location</span>
              <Input v-model="locationName" class="input" />
            </label>
          </div>

          <label class="field">
            <span class="label">Description</span>
            <textarea v-model="description" class="textarea" />
          </label>

          <label class="field">
            <span class="label">Items</span>
            <textarea v-model="itemsText" class="textarea" />
          </label>

          <Button class="primary-button" type="submit">
            <PlusCircle :size="16" />
            Create event
          </Button>
        </form>
      </section>

      <section class="panel">
        <div class="panel-header">
          <h2 class="panel-title">Preview</h2>
          <span class="small muted">{{ parseItems().length }} items</span>
        </div>
        <div class="item-list">
          <article v-for="item in parseItems()" :key="`${item.name}-${item.sort_order}`" class="item-card">
            <div class="item-top">
              <div>
                <div class="item-name">{{ item.name }}</div>
                <div class="item-meta">{{ item.category }} &middot; {{ item.needed_qty }} {{ item.unit || 'needed' }}</div>
              </div>
              <span class="badge soft"><Plus :size="12" /> New</span>
            </div>
          </article>
        </div>
      </section>
    </main>

    <AppToast />
  </div>
</template>
