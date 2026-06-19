<script setup lang="ts">
import { onMounted } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'
import HeroSection from '~/components/sections/HeroSection.vue'
import ServicesGrid from '~/components/sections/ServicesGrid.vue'
import TestimonialSection from '~/components/sections/TestimonialSection.vue'
import SupportSection from '~/components/sections/SupportSection.vue'

if (import.meta.client) {
  gsap.registerPlugin(ScrollTrigger)
}

onMounted(() => {
  // GSAP animations can be initialized here
})
</script>

<template>
  <div class="flex flex-col gap-32">
    <HeroSection />
    <ServicesGrid />
    <TestimonialSection />
    <SupportSection />
  </div>
</template>
