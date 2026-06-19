<script setup lang="ts">
defineProps<{
  title: string
  amount: string
  change: string
  changeClass: string
  helper: string
}>()
</script>

<template>
  <article class="metric-card">
    <div class="metric-card-head">
      <div class="metric-title">
        <span class="mini-badge">
          <svg viewBox="0 0 16 16" aria-hidden="true">
            <path d="M3.6 8H12.4"></path>
            <path d="M8 3.6V12.4"></path>
          </svg>
        </span>
        <span>{{ title }}</span>
      </div>
      <button class="link-button" type="button">View more</button>
    </div>
    <strong class="metric-value">{{ amount }} <span class="trend" :class="changeClass">{{ change }}</span></strong>
    <p>{{ helper }}</p>
  </article>
</template>
