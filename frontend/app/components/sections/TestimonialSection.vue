<script setup lang="ts">
import { onMounted } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'

onMounted(() => {
  if (import.meta.client) {
    gsap.registerPlugin(ScrollTrigger)
    
    gsap.fromTo('.testimonial-card', 
      { x: -50, opacity: 0 },
      { 
        x: 0, 
        opacity: 1, 
        duration: 1, 
        ease: 'power3.out',
        scrollTrigger: {
          trigger: '.testimonial-section',
          start: 'top 75%',
        }
      }
    )

    gsap.fromTo('.testimonial-title', 
      { x: 50, opacity: 0 },
      { 
        x: 0, 
        opacity: 1, 
        duration: 1, 
        ease: 'power3.out',
        scrollTrigger: {
          trigger: '.testimonial-section',
          start: 'top 75%',
        }
      }
    )
  }
})
</script>

<template>
  <section class="testimonial-section grid grid-cols-1 lg:grid-cols-2 gap-16 items-center mt-32">
    
    <!-- Left Card -->
    <div class="testimonial-card relative w-full lg:w-[110%] bg-card rounded-[40px] p-8 md:p-12 pl-32 md:pl-48">
      <!-- Image Pill overlapping on the left -->
      <div class="absolute left-[-20px] md:left-[-40px] top-1/2 transform -translate-y-1/2 w-40 md:w-56 h-64 md:h-80 overflow-hidden pill-mask-vertical shadow-2xl">
        <img src="https://images.unsplash.com/photo-1500648767791-00dcc994a43e?q=80&w=600&auto=format&fit=crop" alt="Alex T." class="w-full h-full object-cover" />
      </div>

      <div class="space-y-6">
        <div>
          <h3 class="text-3xl font-bold text-card-foreground mb-1">Alex T.</h3>
          <p class="text-muted-foreground/60 text-sm">Frequent Event Host</p>
        </div>
        <p class="text-muted-foreground text-sm leading-relaxed max-w-[350px]">
          Organizing our Luleå BBQ used to be a mess of WhatsApp messages and spreadsheets. With BringIt, people just clicked the link, claimed the sausages, and showed up. It saved me hours of headaches!
        </p>
        <button class="bg-secondary text-secondary-foreground font-semibold py-3 px-8 rounded-full hover:bg-secondary/80 transition-colors">
          Read case study
        </button>
      </div>

      <!-- Pagination dots -->
      <div class="absolute -bottom-16 left-1/2 transform -translate-x-1/2 flex gap-3">
        <div class="w-2.5 h-2.5 rounded-full bg-foreground cursor-pointer"></div>
        <div class="w-2.5 h-2.5 rounded-full bg-foreground/20 cursor-pointer hover:bg-foreground/50 transition-colors"></div>
        <div class="w-2.5 h-2.5 rounded-full bg-foreground/20 cursor-pointer hover:bg-foreground/50 transition-colors"></div>
        <div class="w-2.5 h-2.5 rounded-full bg-foreground/20 cursor-pointer hover:bg-foreground/50 transition-colors"></div>
        <div class="w-2.5 h-2.5 rounded-full bg-foreground/20 cursor-pointer hover:bg-foreground/50 transition-colors"></div>
      </div>
    </div>

    <!-- Right Title -->
    <div class="testimonial-title flex flex-col justify-center lg:ml-20 mt-20 lg:mt-0">
      <div class="flex items-start gap-4">
        <div class="w-12 h-12 rounded-full bg-card flex items-center justify-center shrink-0 mt-2 border border-border">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-card-foreground">
            <path d="M12 2L12 22M2 12L22 12M5 5L19 19M19 5L5 19" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>
        <h2 class="text-6xl md:text-8xl font-bold tracking-tighter text-foreground mix-blend-normal leading-[0.9]">
          Our<br/>
          users<br/>
          say
        </h2>
      </div>
    </div>

  </section>
</template>
