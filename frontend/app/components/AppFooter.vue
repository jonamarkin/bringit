<script setup lang="ts">
// Footer component
</script>

<template>
  <footer class="mt-32 pb-8 px-4 md:px-8 lg:px-16 container mx-auto text-sm">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 lg:gap-24 items-start bg-card rounded-[40px] p-8 md:p-16">
      
      <!-- Calculator / Lead Form -->
      <div class="space-y-8">
        <h2 class="text-4xl md:text-5xl font-semibold tracking-tight text-card-foreground">Ready to host?</h2>
        
        <div class="flex flex-wrap gap-3">
          <button class="px-5 py-2 rounded-full bg-primary-foreground text-primary font-medium transition-transform hover:scale-105">BBQ</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">Birthday</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">Potluck</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">Wedding</button>
          <button class="px-5 py-2 rounded-full bg-background text-muted-foreground hover:text-foreground transition-all hover:bg-foreground/10">Dinner</button>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1.5">
            <label class="text-xs text-muted-foreground px-2">Event Name</label>
            <input type="text" placeholder="Summer BBQ" class="w-full bg-background rounded-full px-5 py-4 outline-none border border-transparent focus:border-primary transition-colors text-card-foreground" />
          </div>
          <div class="space-y-1.5">
            <label class="text-xs text-muted-foreground px-2">Event Date</label>
            <input type="text" placeholder="12 Aug - 5 PM" class="w-full bg-background rounded-full px-5 py-4 outline-none border border-transparent focus:border-primary transition-colors text-card-foreground" />
          </div>
        </div>

        <button class="w-full py-4 rounded-full bg-primary text-primary-foreground font-semibold hover:bg-primary/90 transition-all hover:scale-[1.02] active:scale-[0.98]">
          Create Event
        </button>
      </div>

      <!-- Links -->
      <div class="grid grid-cols-2 sm:grid-cols-3 gap-8 lg:pl-12">
        <div class="space-y-6">
          <h4 class="font-semibold text-card-foreground">Features</h4>
          <ul class="space-y-3 text-muted-foreground">
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Instant RSVP</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Item Claiming</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Host Dashboard</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Expense Splitting</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Smart Reminders</NuxtLink></li>
          </ul>
        </div>
        <div class="space-y-6">
          <h4 class="font-semibold text-card-foreground">BringIt</h4>
          <ul class="space-y-3 text-muted-foreground">
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">About us</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Career</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Help center</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">FAQ</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Blog</NuxtLink></li>
          </ul>
        </div>
        <div class="space-y-6">
          <h4 class="font-semibold text-card-foreground">Resources</h4>
          <ul class="space-y-3 text-muted-foreground">
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Event Templates</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Affiliate Program</NuxtLink></li>
            <li><NuxtLink to="#" class="hover:text-foreground transition-colors">Community</NuxtLink></li>
          </ul>
        </div>
      </div>
    </div>
    
    <div class="flex flex-col sm:flex-row items-center justify-between mt-8 pt-8 border-t border-border text-muted-foreground">
      <div class="flex items-center gap-2">
         <div class="w-5 h-5 rounded-full border-2 border-dashed border-foreground/40"></div>
         <span>BringIt Platform</span>
      </div>
      <NuxtLink to="#" class="hover:text-foreground transition-colors mt-4 sm:mt-0">Privacy Policy</NuxtLink>
    </div>
  </footer>
</template>
