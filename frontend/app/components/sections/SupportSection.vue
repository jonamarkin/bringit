<script setup lang="ts">
import { onMounted } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'

onMounted(() => {
  if (import.meta.client) {
    gsap.registerPlugin(ScrollTrigger)
    
    gsap.fromTo('.support-text', 
      { y: 50, opacity: 0 },
      { 
        y: 0, 
        opacity: 1, 
        duration: 1, 
        stagger: 0.2,
        ease: 'power3.out',
        scrollTrigger: {
          trigger: '.support-section',
          start: 'top 70%',
        }
      }
    )

    gsap.fromTo('.support-image', 
      { scale: 0.9, opacity: 0 },
      { 
        scale: 1, 
        opacity: 1, 
        duration: 1.5, 
        ease: 'power3.out',
        scrollTrigger: {
          trigger: '.support-section',
          start: 'top 70%',
        }
      }
    )
  }
})
</script>

<template>
  <section class="support-section grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-8 items-center mt-40">
    
    <!-- Left Title -->
    <div class="lg:col-span-4 flex flex-col justify-center">
      <h2 class="text-6xl md:text-7xl xl:text-8xl font-bold tracking-tighter text-[#f2f4f7] mix-blend-lighten leading-[0.9] support-text">
        Support<br/>
        at every<br/>
        Step
      </h2>
    </div>

    <!-- Center Content -->
    <div class="lg:col-span-4 flex flex-col gap-12 justify-center pl-0 lg:pl-8">
      <div class="space-y-4 support-text">
        <div class="w-10 h-10 rounded-full bg-[#1f283d] flex items-center justify-center border border-white/10">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 2L12 22M2 12L22 12M5 5L19 19M19 5L5 19" stroke="white" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>
        <p class="text-white/70 text-sm leading-relaxed pr-8">
          Participate in the <span class="text-[#84d7e9] underline underline-offset-4 decoration-[#84d7e9]/50">loyalty program</span>. Pay for the service online to get bonuses for our other services. We have a lot of bonuses!
        </p>
      </div>

      <div class="space-y-4 support-text">
        <div class="w-10 h-10 rounded-full bg-[#1f283d] flex items-center justify-center border border-white/10">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 2L12 22M2 12L22 12M5 5L19 19M19 5L5 19" stroke="white" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>
        <p class="text-white/70 text-sm leading-relaxed pr-8">
          You can help other people: choose a donation fund and share the link with your friends!
        </p>
        <button class="mt-4 border border-[#ff6476] text-[#ff6476] font-medium py-3 px-8 rounded-full hover:bg-[#ff6476] hover:text-white transition-colors duration-300">
          Choose a fund
        </button>
      </div>
    </div>

    <!-- Right Image -->
    <div class="lg:col-span-4 flex justify-center lg:justify-end support-image">
      <div class="relative w-[320px] h-[480px] overflow-hidden pill-mask-vertical shadow-[0_0_60px_rgba(255,255,255,0.05)]">
        <!-- Woman with child and red balloon placeholder -->
        <img src="https://images.unsplash.com/photo-1511895426328-dc8714191300?q=80&w=1000&auto=format&fit=crop" alt="Support program" class="w-full h-full object-cover transform hover:scale-105 transition-transform duration-1000" />
      </div>
    </div>

  </section>
</template>
