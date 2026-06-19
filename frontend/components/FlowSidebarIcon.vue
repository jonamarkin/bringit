<script setup lang="ts">
defineProps<{
  name: string
}>()
</script>

<template>
  <svg v-if="name === 'grid'" viewBox="0 0 16 16" aria-hidden="true">
    <rect x="2.5" y="2.5" width="4.2" height="4.2" rx="1.2"></rect>
    <rect x="9.3" y="2.5" width="4.2" height="4.2" rx="1.2"></rect>
    <rect x="2.5" y="9.3" width="4.2" height="4.2" rx="1.2"></rect>
    <rect x="9.3" y="9.3" width="4.2" height="4.2" rx="1.2"></rect>
  </svg>

  <svg v-else-if="name === 'transaction'" viewBox="0 0 16 16" aria-hidden="true">
    <rect x="2.5" y="3.5" width="11" height="9" rx="2"></rect>
    <path d="M5 6.5H11"></path>
    <path d="M5 9.5H9"></path>
  </svg>

  <svg v-else-if="name === 'payment'" viewBox="0 0 16 16" aria-hidden="true">
    <rect x="2.7" y="3.2" width="10.6" height="9.6" rx="2"></rect>
    <path d="M2.7 6.4H13.3"></path>
  </svg>

  <svg v-else-if="name === 'report'" viewBox="0 0 16 16" aria-hidden="true">
    <path d="M3 12.5V3.5"></path>
    <path d="M6.5 12.5V7"></path>
    <path d="M10 12.5V5"></path>
    <path d="M13 12.5V8.4"></path>
  </svg>

  <svg v-else-if="name === 'budget'" viewBox="0 0 16 16" aria-hidden="true">
    <path d="M4 3.5H12"></path>
    <path d="M4 6.8H12"></path>
    <path d="M4 10.1H10.5"></path>
    <path d="M4 13.4H9"></path>
  </svg>

  <svg v-else-if="name === 'star'" viewBox="0 0 16 16" aria-hidden="true">
    <path d="M8 2.7L9.6 6.1L13.3 6.5L10.5 8.9L11.3 12.5L8 10.7L4.7 12.5L5.5 8.9L2.7 6.5L6.4 6.1Z"></path>
  </svg>

  <svg v-else-if="name === 'savings'" viewBox="0 0 16 16" aria-hidden="true">
    <path d="M3.5 10.8C3.5 7.7 5.9 5.2 9 5.2H12.7"></path>
    <path d="M12.7 5.2L10.8 3.5"></path>
    <path d="M12.7 5.2L11 7.1"></path>
    <path d="M12.5 12.5H3.5"></path>
  </svg>

  <svg v-else-if="name === 'tracker'" viewBox="0 0 16 16" aria-hidden="true">
    <path d="M3.3 11.8L6.2 8.9L8.3 11L12.7 6.6"></path>
    <path d="M10.1 6.6H12.7V9.2"></path>
  </svg>

  <svg v-else-if="name === 'globe'" viewBox="0 0 16 16" aria-hidden="true">
    <circle cx="8" cy="8" r="4.8"></circle>
    <path d="M8 2.5V13.5"></path>
    <path d="M2.5 8H13.5"></path>
  </svg>

  <svg v-else-if="name === 'help'" viewBox="0 0 16 16" aria-hidden="true">
    <circle cx="8" cy="8" r="5.2"></circle>
    <path d="M6.4 6.4C6.4 5.5 7 4.9 8 4.9C8.9 4.9 9.6 5.5 9.6 6.4C9.6 7.1 9.2 7.5 8.5 7.9C8.1 8.1 8 8.4 8 8.8"></path>
    <circle cx="8" cy="11.2" r="0.55" fill="currentColor" stroke="none"></circle>
  </svg>

  <svg v-else-if="name === 'setting'" viewBox="0 0 16 16" aria-hidden="true">
    <circle cx="8" cy="8" r="2.4"></circle>
    <path d="M8 2.5V4"></path>
    <path d="M8 12V13.5"></path>
    <path d="M13.5 8H12"></path>
    <path d="M4 8H2.5"></path>
    <path d="M11.9 4.1L10.9 5.1"></path>
    <path d="M5.1 10.9L4.1 11.9"></path>
    <path d="M11.9 11.9L10.9 10.9"></path>
    <path d="M5.1 5.1L4.1 4.1"></path>
  </svg>

  <svg v-else viewBox="0 0 16 16" aria-hidden="true">
    <circle cx="8" cy="5.4" r="2.2"></circle>
    <path d="M3.8 12.8C4.6 10.9 6.1 10 8 10C9.9 10 11.4 10.9 12.2 12.8"></path>
  </svg>
</template>
