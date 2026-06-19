<script setup lang="ts">
import AppHeader from '~/components/AppHeader.vue'
import AppFooter from '~/components/AppFooter.vue'
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <AppHeader />
    <main class="flex-grow pt-32 pb-16 px-4 md:px-8 lg:px-16 container mx-auto">
      <slot />
    </main>
    <AppFooter />
  </div>
</template>
