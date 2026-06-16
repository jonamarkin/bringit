<script setup lang="ts">
import type { EventStats } from '~/types/bringit'

defineProps<{
  stats: EventStats
}>()
</script>

<template>
  <div class="metric-row">
    <div class="metric">
      <div class="metric-value">{{ stats.attending }}</div>
      <div class="metric-label">Going</div>
    </div>
    <div class="metric">
      <div class="metric-value">{{ stats.items_claimed }}</div>
      <div class="metric-label">Claimed</div>
    </div>
    <div class="metric">
      <div class="metric-value">{{ stats.items_remaining }}</div>
      <div class="metric-label">Left</div>
    </div>
    <div class="metric">
      <div class="metric-value">{{ Math.min(100, Math.round((stats.items_claimed / Math.max(stats.items_needed, 1)) * 100)) }}%</div>
      <div class="metric-label">Covered</div>
    </div>
  </div>
</template>
