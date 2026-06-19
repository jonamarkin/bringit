<script setup lang="ts">
// Footer component
</script>

<template>
  <footer class="mt-24 pb-8 px-4 md:px-8 lg:px-16 container mx-auto text-sm">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-start bg-[#1f283d] rounded-[40px] p-8 md:p-16">
      
      <!-- Calculator / Lead Form -->
      <div class="space-y-8">
        <h2 class="text-4xl md:text-5xl font-semibold tracking-tight text-white">Need numbers?</h2>
        
        <div class="flex flex-wrap gap-3">
          <button class="px-5 py-2 rounded-full bg-white text-[#0b0f17] font-medium transition-transform hover:scale-105">Travel</button>
          <button class="px-5 py-2 rounded-full bg-[#111726] text-white/70 hover:text-white transition-all hover:bg-white/10">Health</button>
          <button class="px-5 py-2 rounded-full bg-[#111726] text-white/70 hover:text-white transition-all hover:bg-white/10">CASCO</button>
          <button class="px-5 py-2 rounded-full bg-[#111726] text-white/70 hover:text-white transition-all hover:bg-white/10">Mortgage</button>
          <button class="px-5 py-2 rounded-full bg-[#111726] text-white/70 hover:text-white transition-all hover:bg-white/10">Fire</button>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1.5">
            <label class="text-xs text-white/50 px-2">Enter the country</label>
            <input type="text" placeholder="Australia" class="w-full bg-[#111726] rounded-full px-5 py-4 outline-none border border-transparent focus:border-[#ff6476] transition-colors text-white" />
          </div>
          <div class="space-y-1.5">
            <label class="text-xs text-white/50 px-2">Choose travel dates</label>
            <input type="text" placeholder="11 Apr - 20 Apr" class="w-full bg-[#111726] rounded-full px-5 py-4 outline-none border border-transparent focus:border-[#ff6476] transition-colors text-white" />
          </div>
        </div>

        <button class="w-full py-4 rounded-full bg-[#ff6476] text-white font-semibold hover:bg-[#e05666] transition-all hover:scale-[1.02] active:scale-[0.98]">
          Calculate
        </button>
      </div>

      <!-- Links -->
      <div class="grid grid-cols-2 sm:grid-cols-3 gap-8 lg:pl-12">
        <div class="space-y-6">
          <h4 class="font-semibold text-white">Insurance services</h4>
          <ul class="space-y-3 text-white/60">
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Medical expenses</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Family health insurance</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Vehicle damage</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Property loss</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Fire</NuxtLink></li>
          </ul>
        </div>
        <div class="space-y-6">
          <h4 class="font-semibold text-white">Company</h4>
          <ul class="space-y-3 text-white/60">
            <li><NuxtLink to="#" class="hover:text-white transition-colors">About us</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Career</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Help</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">FAQ</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Blog</NuxtLink></li>
          </ul>
        </div>
        <div class="space-y-6">
          <h4 class="font-semibold text-white">Bonus</h4>
          <ul class="space-y-3 text-white/60">
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Gift certificate</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Affiliate Program</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-white transition-colors">Trainings</NuxtLink></li>
          </ul>
        </div>
      </div>
    </div>
    
    <div class="flex flex-col sm:flex-row items-center justify-between mt-8 pt-8 border-t border-white/10 text-white/40">
      <div class="flex items-center gap-2">
         <div class="w-5 h-5 rounded-full border-2 border-dashed border-white/40"></div>
         <span>Darcy's insurance products</span>
      </div>
      <NuxtLink to="#" class="hover:text-white transition-colors mt-4 sm:mt-0">Privacy Policy</NuxtLink>
    </div>
  </footer>
</template>
