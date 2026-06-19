<script setup lang="ts">
import { ArrowDown } from 'lucide-vue-next'
</script>

<template>
  <section class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center min-h-[80vh] relative pt-12">
    <!-- Left Typography -->
    <div class="lg:col-span-7 flex flex-col justify-center">
      <h1 class="text-[5rem] sm:text-[6rem] md:text-[8rem] lg:text-[9rem] leading-[0.95] font-bold tracking-tighter text-foreground mix-blend-normal z-10 hero-title">
        <span class="block overflow-hidden pb-8 -mb-8"><span class="block transform translate-y-full title-line">BringIt</span></span>
        <span class="block overflow-hidden pb-8 -mb-8"><span class="block transform translate-y-full title-line">to the</span></span>
        <span class="block overflow-hidden pb-8 -mb-8"><span class="block transform translate-y-full title-line">party</span></span>
      </h1>
    </div>

    <!-- Center text blurbs -->
    <div class="lg:col-span-2 flex flex-col gap-12 justify-center lg:ml-[-2rem] z-20 mt-8 lg:mt-0 hero-blurbs opacity-0">
      <div class="space-y-4 max-w-[200px]">
        <div class="w-10 h-10 rounded-full bg-card flex items-center justify-center border border-border">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-card-foreground">
            <path d="M12 2L12 22M2 12L22 12M5 5L19 19M19 5L5 19" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>
        <p class="text-muted-foreground text-sm leading-relaxed">
          Coordinate your <span class="text-secondary underline underline-offset-4 decoration-secondary/50">potlucks</span>, BBQs, and dinners without the chaotic group chats.
        </p>
      </div>
      <div class="space-y-4 max-w-[200px]">
        <div class="w-10 h-10 rounded-full bg-card flex items-center justify-center border border-border">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="text-card-foreground">
            <path d="M12 2L12 22M2 12L22 12M5 5L19 19M19 5L5 19" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
        </div>
        <p class="text-muted-foreground text-sm leading-relaxed">
          Let guests RSVP and claim what they're bringing in real-time.
        </p>
      </div>
    </div>

    <!-- Right Image Mask -->
    <div class="lg:col-span-3 flex justify-end h-full mt-12 lg:mt-0">
      <div class="relative w-[300px] h-[550px] overflow-hidden pill-mask-vertical shadow-[0_0_80px_rgba(132,215,233,0.15)] hero-image transform scale-95 opacity-0">
        <!-- Using a placeholder for the balloons image -->
        <img src="https://images.unsplash.com/photo-1555939594-58d7cb561ad1?q=80&w=1000&auto=format&fit=crop" alt="BBQ Party" class="object-cover w-full h-full transform hover:scale-110 transition-transform duration-1000" />
        <!-- Soft glow overlay -->
        <div class="absolute inset-0 shadow-[inset_0_0_50px_rgba(0,0,0,0.5)]"></div>
      </div>
    </div>

    <!-- Scroll Down Indicator -->
    <div class="absolute bottom-[-40px] left-8 lg:left-32 animate-bounce cursor-pointer opacity-50 hover:opacity-100 transition-opacity">
       <div class="w-12 h-16 rounded-full bg-background border border-border flex items-center justify-center">
         <ArrowDown class="w-5 h-5 text-foreground" />
       </div>
    </div>
  </section>
</template>

<style scoped>
.title-line {
  animation: reveal 1s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}
.title-line:nth-child(1) { animation-delay: 0.1s; }
.title-line:nth-child(2) { animation-delay: 0.2s; }
.title-line:nth-child(3) { animation-delay: 0.3s; }

.hero-image {
  animation: fade-in-up 1.5s cubic-bezier(0.16, 1, 0.3, 1) 0.4s forwards;
}

.hero-blurbs {
  animation: fade-in-up 1.2s cubic-bezier(0.16, 1, 0.3, 1) 0.6s forwards;
}

@keyframes reveal {
  0% { transform: translateY(100%); }
  100% { transform: translateY(0); }
}

@keyframes fade-in-up {
  0% { opacity: 0; transform: translateY(40px) scale(0.95); }
  100% { opacity: 1; transform: translateY(0) scale(1); }
}
</style>
