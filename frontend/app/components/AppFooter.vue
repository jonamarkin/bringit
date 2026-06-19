<script setup lang="ts">
// Footer component
</script>

<template>
  <footer class="mt-24 pb-8 px-4 md:px-8 lg:px-16 container mx-auto text-sm">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-start bg-card rounded-[40px] p-8 md:p-16">
      
      <!-- Calculator / Lead Form -->
      <div class="space-y-8">
        <h2 class="text-4xl md:text-5xl font-semibold tracking-tight text-card-foreground">Need numbers?</h2>
        
        <div class="flex flex-wrap gap-3">
          <button class="px-5 py-2 rounded-full bg-primary-foreground text-primary font-medium transition-transform hover:scale-105">Travel</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">Health</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">CASCO</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">Mortgage</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">Fire</button>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1.5">
            <label class="text-xs text-muted-foreground px-2">Enter the country</label>
            <input type="text" placeholder="Australia" class="w-full bg-background rounded-full px-5 py-4 outline-none border border-transparent focus:border-primary transition-colors text-card-foreground" />
          </div>
          <div class="space-y-1.5">
            <label class="text-xs text-muted-foreground px-2">Choose travel dates</label>
            <input type="text" placeholder="11 Apr - 20 Apr" class="w-full bg-background rounded-full px-5 py-4 outline-none border border-transparent focus:border-primary transition-colors text-card-foreground" />
          </div>
        </div>

        <button class="w-full py-4 rounded-full bg-primary text-primary-foreground font-semibold hover:bg-primary/90 transition-all hover:scale-[1.02] active:scale-[0.98]">
          Calculate
        </button>
      </div>

      <!-- Links -->
      <div class="grid grid-cols-2 sm:grid-cols-3 gap-8 lg:pl-12">
        <div class="space-y-6">
          <h4 class="font-semibold text-card-foreground">Insurance services</h4>
          <ul class="space-y-3 text-muted-foreground">
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Medical expenses</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Family health insurance</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Vehicle damage</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Property loss</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Fire</NuxtLink></li>
          </ul>
        </div>
        <div class="space-y-6">
          <h4 class="font-semibold text-card-foreground">Company</h4>
          <ul class="space-y-3 text-muted-foreground">
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">About us</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Career</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Help</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">FAQ</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Blog</NuxtLink></li>
          </ul>
        </div>
        <div class="space-y-6">
          <h4 class="font-semibold text-card-foreground">Bonus</h4>
          <ul class="space-y-3 text-muted-foreground">
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Gift certificate</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Affiliate Program</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Trainings</NuxtLink></li>
          </ul>
        </div>
      </div>
    </div>
    
    <div class="flex flex-col sm:flex-row items-center justify-between mt-8 pt-8 border-t border-border text-muted-foreground">
      <div class="flex items-center gap-2">
         <div class="w-5 h-5 rounded-full border-2 border-dashed border-foreground/40"></div>
         <span>Darcy's insurance products</span>
      </div>
      <NuxtLink to="#" class="hover:text-foreground transition-colors mt-4 sm:mt-0">Privacy Policy</NuxtLink>
    </div>
  </footer>
</template>
