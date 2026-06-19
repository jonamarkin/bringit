<script setup lang="ts">
import type { BringItNavId } from '~/composables/useBringItNavigation'

defineProps<{
  activeNav?: BringItNavId
  mobileOpen?: boolean
}>()

defineEmits<{
  close: []
  navigate: [id: BringItNavId]
}>()

const { sidebarLinks } = useBringItNavigation()
</script>

<template>
  <aside class="sidebar" :class="{ 'sidebar-open': mobileOpen }">
    <div class="brand-row">
      <div class="brand">
        <span class="brand-mark" aria-hidden="true"></span>
        <span class="brand-name">BringIt</span>
      </div>
      <button class="sidebar-mini-button" type="button" aria-label="Close sidebar" @click="$emit('close')">
        <svg viewBox="0 0 16 16" aria-hidden="true">
          <rect x="3.2" y="3.2" width="9.6" height="9.6" rx="2.3"></rect>
          <path d="M6 5.5V10.5"></path>
          <path d="M10 5.5V10.5"></path>
        </svg>
      </button>
    </div>

    <div class="sidebar-group">
      <span class="sidebar-label">General</span>
      <nav class="sidebar-nav">
        <button
          v-for="link in sidebarLinks.general"
          :key="link.id"
          class="sidebar-link"
          :class="{ active: activeNav === link.id }"
          type="button"
          @click="$emit('navigate', link.id)"
        >
          <span class="sidebar-icon">
            <FlowSidebarIcon :name="link.icon" />
          </span>
          <span>{{ link.label }}</span>
        </button>
      </nav>
    </div>

    <div class="sidebar-group">
      <span class="sidebar-label">Tools</span>
      <nav class="sidebar-nav">
        <button
          v-for="link in sidebarLinks.tools"
          :key="link.id"
          class="sidebar-link"
          :class="{ active: activeNav === link.id }"
          type="button"
          @click="$emit('navigate', link.id)"
        >
          <span class="sidebar-icon">
            <FlowSidebarIcon :name="link.icon" />
          </span>
          <span>{{ link.label }}</span>
        </button>
      </nav>
    </div>

    <div class="sidebar-footer">
      <button
        v-for="link in sidebarLinks.footer"
        :key="link.id"
        class="sidebar-link"
        :class="{ active: activeNav === link.id }"
        type="button"
        @click="$emit('navigate', link.id)"
      >
        <span class="sidebar-icon">
          <FlowSidebarIcon :name="link.icon" />
        </span>
        <span>{{ link.label }}</span>
      </button>
    </div>
  </aside>
</template>
