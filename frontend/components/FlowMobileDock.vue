<script setup lang="ts">
import type { BringItNavId } from '~/composables/useBringItNavigation'

defineProps<{
  activeNav?: BringItNavId
  moreOpen?: boolean
}>()

defineEmits<{
  navigate: [id: BringItNavId]
  more: []
}>()

const { dockLinks } = useBringItNavigation()
</script>

<template>
  <nav class="mobile-dock" aria-label="Mobile navigation">
    <button
      v-for="item in dockLinks"
      :key="item.id"
      class="dock-item"
      :class="{ active: activeNav === item.id }"
      type="button"
      @click="$emit('navigate', item.id)"
    >
      <span class="dock-icon">
        <FlowSidebarIcon :name="item.icon" />
      </span>
      <span>{{ item.label }}</span>
    </button>

    <button class="dock-item dock-more" :class="{ active: moreOpen }" type="button" @click="$emit('more')">
      <span class="dock-icon">
        <svg viewBox="0 0 16 16" aria-hidden="true">
          <circle cx="4" cy="8" r="1.1"></circle>
          <circle cx="8" cy="8" r="1.1"></circle>
          <circle cx="12" cy="8" r="1.1"></circle>
        </svg>
      </span>
      <span>More</span>
    </button>
  </nav>
</template>
