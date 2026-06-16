export interface User {
  id: string
  email: string
  display_name: string
}

export interface EventStats {
  attending: number
  maybe: number
  not_attending: number
  party_size_total: number
  items_needed: number
  items_claimed: number
  items_remaining: number
  claims_count: number
}

export interface BringItEvent {
  id: string
  title: string
  description: string
  starts_at: string
  timezone: string
  location_name: string
  location_note: string
  share_code: string
  status: string
}

export interface EventItem {
  id: string
  event_id: string
  name: string
  category: string
  needed_qty: number
  claimed_qty: number
  unit: string
  priority: string
  notes: string
  sort_order: number
  is_purchased: boolean
}

export interface Guest {
  id: string
  event_id: string
  name: string
  email: string
  phone: string
  rsvp_status: string
  party_size: number
  note: string
  session_token?: string
}

export interface ItemClaim {
  id: string
  event_id: string
  item_id: string
  guest_id: string
  item_name: string
  guest_name: string
  quantity: number
  note: string
  status: string
  created_at: string
}

export interface Activity {
  id: string
  actor_name: string
  type: string
  message: string
  created_at: string
}

export interface Notification {
  id: string
  title: string
  message: string
  type: string
  is_read: boolean
  created_at: string
}

export interface EventSnapshot {
  event: BringItEvent
  stats: EventStats
  items: EventItem[]
  guests: Guest[]
  claims: ItemClaim[]
  activities: Activity[]
}

export interface EventSummary extends BringItEvent {
  stats: EventStats
}
