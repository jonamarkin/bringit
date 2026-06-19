<script setup lang="ts">
import { Moon, Sun } from 'lucide-vue-next'

const colorMode = useColorMode()

const toggleTheme = () => {
  colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark'
}
</script>

<template>
  <button 
    @click="toggleTheme" 
    class="flex items-center justify-center w-10 h-10 rounded-full bg-card hover:bg-secondary/20 transition-colors border border-border"
    aria-label="Toggle theme"
  >
    <Sun v-if="colorMode.value === 'dark'" class="w-5 h-5 text-foreground" />
    <Moon v-else class="w-5 h-5 text-foreground" />
  </button>
</template>
