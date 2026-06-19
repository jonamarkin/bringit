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
      <h2 class="text-6xl md:text-7xl xl:text-8xl font-bold tracking-tighter text-foreground mix-blend-normal leading-[0.9] support-text">
        Split the<br/>
        bill<br/>
        effortlessly
      </h2>
    </div>

    <!-- Center Content -->
    <div class="lg:col-span-4 flex flex-col gap-12 justify-center pl-0 lg:pl-8">
      <div class="space-y-4 support-text">
        <div class="w-10 h-10 rounded-full bg-card flex items-center justify-center border border-border">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-card-foreground">
            <path d="M12 2L12 22M2 12L22 12M5 5L19 19M19 5L5 19" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>
        <p class="text-muted-foreground text-sm leading-relaxed pr-8">
          Organizing an event shouldn't mean you pay for everything. With <span class="text-secondary underline underline-offset-4 decoration-secondary/50">BringIt</span>, you can easily track who bought what and split the costs.
        </p>
      </div>

      <div class="space-y-4 support-text">
        <div class="w-10 h-10 rounded-full bg-card flex items-center justify-center border border-border">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-card-foreground">
            <path d="M12 2L12 22M2 12L22 12M5 5L19 19M19 5L5 19" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>
        <p class="text-muted-foreground text-sm leading-relaxed pr-8">
          Let your friends contribute exactly their share, directly from the event page.
        </p>
        <button class="mt-4 border border-primary text-primary font-medium py-3 px-8 rounded-full hover:bg-primary hover:text-primary-foreground transition-colors duration-300">
          Learn about payments
        </button>
      </div>
    </div>

    <!-- Right Image -->
    <div class="lg:col-span-4 flex justify-center lg:justify-end support-image">
      <div class="relative w-[320px] h-[480px] overflow-hidden pill-mask-vertical shadow-[0_0_60px_rgba(255,255,255,0.05)]">
        <!-- Woman with card placeholder -->
        <img src="https://images.unsplash.com/photo-1554224155-8d04cb21cd6c?q=80&w=1000&auto=format&fit=crop" alt="Splitting bill" class="w-full h-full object-cover transform hover:scale-105 transition-transform duration-1000" />
      </div>
    </div>

  </section>
</template>
