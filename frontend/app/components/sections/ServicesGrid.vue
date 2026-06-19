<script setup lang="ts">
import { onMounted } from 'vue'
import { gsap } from 'gsap'
import { ScrollTrigger } from 'gsap/ScrollTrigger'

onMounted(() => {
  if (import.meta.client) {
    gsap.registerPlugin(ScrollTrigger)
    
    gsap.fromTo('.service-card', 
      { y: 50, opacity: 0 },
      { 
        y: 0, 
        opacity: 1, 
        duration: 0.8, 
        stagger: 0.15,
        ease: 'power3.out',
        scrollTrigger: {
          trigger: '.services-grid-container',
          start: 'top 80%',
        }
      }
    )
  }
})
</script>

<template>
  <section class="services-grid-container mt-32">
    <!-- Top Row -->
    <div class="grid grid-cols-1 md:grid-cols-12 gap-6 mb-6">
      
      <!-- Home Insurance Card -->
      <div class="md:col-span-5 service-card relative rounded-[40px] overflow-hidden group h-[320px] shadow-2xl cursor-pointer">
        <img src="https://images.unsplash.com/photo-1583847268964-b28ce8f52859?q=80&w=1000&auto=format&fit=crop" alt="Home interior" class="absolute inset-0 w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-700" />
        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/20 to-transparent"></div>
        <div class="absolute bottom-6 left-6 right-6">
          <div class="bg-[#111726]/80 backdrop-blur-md px-6 py-3 rounded-full text-white font-medium inline-flex items-center gap-2 border border-white/10 group-hover:bg-white group-hover:text-black transition-colors">
            Home insurance
          </div>
        </div>
      </div>

      <!-- Lifebuoy Image Pill -->
      <div class="md:col-span-3 service-card flex justify-center">
        <div class="relative w-full max-w-[220px] h-[320px] overflow-hidden pill-mask-vertical shadow-[0_0_40px_rgba(255,100,118,0.1)] group">
          <img src="https://images.unsplash.com/photo-1601625902170-c509ce86b840?q=80&w=1000&auto=format&fit=crop" alt="Lifebuoy" class="w-full h-full object-cover transform group-hover:scale-105 transition-transform duration-700" />
          <div class="absolute inset-0 shadow-[inset_0_0_30px_rgba(0,0,0,0.3)]"></div>
        </div>
      </div>

      <!-- Best Choice Card -->
      <div class="md:col-span-4 service-card bg-[#1f283d] rounded-[40px] p-8 flex flex-col justify-between h-[320px] relative overflow-hidden group">
        <!-- Abstract gradient blob -->
        <div class="absolute -right-20 -top-20 w-48 h-48 bg-[#84d7e9] opacity-10 rounded-full blur-[50px] group-hover:opacity-20 transition-opacity"></div>
        <div class="relative z-10">
          <h3 class="text-4xl font-bold text-white mb-4">Best choice</h3>
          <p class="text-white/60 text-sm leading-relaxed">
            Not all insurance are the same. That's why agency knows that your insurance should be customized to fit your situation.
          </p>
        </div>
        <button class="relative z-10 bg-[#ff6476] text-white font-semibold py-4 px-6 rounded-full w-full hover:bg-[#e05666] transition-colors hover:shadow-[0_0_20px_rgba(255,100,118,0.4)]">
          View all services
        </button>
      </div>
      
    </div>

    <!-- Bottom Row -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
      
      <!-- Vehicle Damage -->
      <div class="service-card relative rounded-[40px] overflow-hidden group h-[200px] shadow-xl cursor-pointer">
        <img src="https://images.unsplash.com/photo-1549317661-bd32c8ce0db2?q=80&w=1000&auto=format&fit=crop" alt="Blue sports car" class="absolute inset-0 w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-700" />
        <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
        <div class="absolute bottom-4 left-4">
          <div class="bg-[#111726]/80 backdrop-blur-md px-5 py-2.5 rounded-full text-white text-sm font-medium border border-white/10 group-hover:bg-white group-hover:text-black transition-colors">
            Vehicle damage
          </div>
        </div>
      </div>

      <!-- Mortgage Insurance -->
      <div class="service-card relative rounded-[40px] overflow-hidden group h-[200px] shadow-xl cursor-pointer">
        <img src="https://images.unsplash.com/photo-1512917774080-9991f1c4c750?q=80&w=1000&auto=format&fit=crop" alt="Red house" class="absolute inset-0 w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-700" />
        <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
        <div class="absolute bottom-4 left-4">
          <div class="bg-[#111726]/80 backdrop-blur-md px-5 py-2.5 rounded-full text-white text-sm font-medium border border-white/10 group-hover:bg-white group-hover:text-black transition-colors">
            Mortgage insurance
          </div>
        </div>
      </div>

      <!-- Medical Expenses -->
      <div class="service-card relative rounded-[40px] overflow-hidden group h-[200px] shadow-xl cursor-pointer">
        <img src="https://images.unsplash.com/photo-1584308666744-24d5c474f2ae?q=80&w=1000&auto=format&fit=crop" alt="Pills" class="absolute inset-0 w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-700" />
        <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
        <div class="absolute bottom-4 left-4">
          <div class="bg-[#111726]/80 backdrop-blur-md px-5 py-2.5 rounded-full text-white text-sm font-medium border border-white/10 group-hover:bg-white group-hover:text-black transition-colors">
            Medical expenses
          </div>
        </div>
      </div>

    </div>
  </section>
</template>
