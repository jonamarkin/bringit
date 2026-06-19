<script setup lang="ts">
import { LogIn } from 'lucide-vue-next'
import ThemeToggle from './ThemeToggle.vue'
</script>

<template>
  <header class="fixed top-0 left-0 right-0 z-50 py-6 px-4 md:px-8 lg:px-16 container mx-auto">
    <div class="flex items-center justify-between rounded-full bg-card/80 backdrop-blur-md px-6 py-4 shadow-xl border border-border/50">
      <!-- Logo -->
      <NuxtLink to="/" class="flex items-center gap-3">
        <div class="w-8 h-8 rounded-full border-[3px] border-dashed border-primary animate-[spin_10s_linear_infinite]"></div>
        <span class="font-bold text-lg text-foreground tracking-tight">BringIt</span>
      </NuxtLink>

      <!-- Navigation -->
      <nav class="hidden lg:flex items-center gap-8 text-sm font-medium">
        <NuxtLink to="/" class="text-foreground hover:text-secondary transition-colors">Home</NuxtLink>
        <NuxtLink to="#" class="text-muted-foreground hover:text-secondary transition-colors flex items-center gap-1">
          Features
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>
        </NuxtLink>
        <NuxtLink to="#" class="text-muted-foreground hover:text-secondary transition-colors">Pricing</NuxtLink>
        <NuxtLink to="#" class="text-muted-foreground hover:text-secondary transition-colors">Use Cases</NuxtLink>
        <NuxtLink to="#" class="text-muted-foreground hover:text-secondary transition-colors">Blog</NuxtLink>
      </nav>

      <!-- Actions -->
      <div class="flex items-center gap-4 sm:gap-6">
        <ThemeToggle />
        <button class="bg-secondary text-secondary-foreground font-semibold px-6 py-2.5 rounded-full hover:bg-secondary/80 transition-colors duration-300 text-sm">
          Create Event
        </button>
        <button class="hidden sm:flex items-center gap-2 text-muted-foreground hover:text-foreground transition-colors text-sm font-medium">
          <LogIn class="w-4 h-4" />
          Login
        </button>
      </div>
    </div>
  </header>
</template>
