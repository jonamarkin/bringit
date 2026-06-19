<script setup lang="ts">
import { LogIn } from 'lucide-vue-next'
</script>

<template>
  <header class="fixed top-0 left-0 right-0 z-50 py-6 px-4 md:px-8 lg:px-16 container mx-auto">
    <div class="flex items-center justify-between rounded-full bg-[#111726]/80 backdrop-blur-md px-6 py-4 shadow-xl border border-white/5">
      <!-- Logo -->
      <NuxtLink to="/" class="flex items-center gap-2">
        <div class="w-8 h-8 rounded-full border-[3px] border-dashed border-[#ff6476] animate-[spin_10s_linear_infinite]"></div>
      </NuxtLink>

      <!-- Navigation -->
      <nav class="hidden lg:flex items-center gap-8 text-sm font-medium">
        <NuxtLink to="/" class="text-white hover:text-[#84d7e9] transition-colors">Home</NuxtLink>
        <NuxtLink to="#" class="text-white/70 hover:text-[#84d7e9] transition-colors flex items-center gap-1">
          Services
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>
        </NuxtLink>
        <NuxtLink to="#" class="text-white/70 hover:text-[#84d7e9] transition-colors">About us</NuxtLink>
        <NuxtLink to="#" class="text-white/70 hover:text-[#84d7e9] transition-colors">Blog</NuxtLink>
        <NuxtLink to="#" class="text-white/70 hover:text-[#84d7e9] transition-colors">Career</NuxtLink>
      </nav>

      <!-- Actions -->
      <div class="flex items-center gap-6">
        <button class="bg-[#84d7e9] text-[#0b0f17] font-semibold px-6 py-2.5 rounded-full hover:bg-white transition-colors duration-300 text-sm">
          Contact us
        </button>
        <button class="hidden sm:flex items-center gap-2 text-white/80 hover:text-white transition-colors text-sm font-medium">
          <LogIn class="w-4 h-4" />
          Login
        </button>
      </div>
    </div>
  </header>
</template>
