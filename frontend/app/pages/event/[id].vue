<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { MapPin, Calendar, Clock, Check, Plus, Minus, User } from 'lucide-vue-next'
import { gsap } from 'gsap'

const route = useRoute()
const eventId = route.params.id

// Mock Event Data
const event = ref({
  id: eventId,
  title: 'Luleå Summer BBQ',
  host: 'Alex T.',
  date: 'Saturday, August 12th',
  time: '17:00',
  location: 'Gültzauudden Beach, Luleå',
  description: 'Let us make the most of the short Swedish summer! We are organizing a BBQ hangout. Please RSVP and claim what you can bring. We will split the remaining costs.',
})

// Mock RSVP State
const isAttending = ref<boolean | null>(null)
const guestName = ref('')
const hasRsvpd = ref(false)

const submitRsvp = () => {
  if (isAttending.value && !guestName.value) return
  hasRsvpd.value = true
}

// Mock Items Data
const items = ref([
  { id: 1, name: 'Sausages (Packs)', totalNeeded: 5, claimedByOthers: 2, myClaim: 0, category: 'Food' },
  { id: 2, name: 'Jollof Rice (Trays)', totalNeeded: 3, claimedByOthers: 1, myClaim: 0, category: 'Food' },
  { id: 3, name: 'Charcoal (Bags)', totalNeeded: 2, claimedByOthers: 0, myClaim: 0, category: 'Supplies' },
  { id: 4, name: 'Soda & Juice (Bottles)', totalNeeded: 10, claimedByOthers: 4, myClaim: 0, category: 'Drinks' },
  { id: 5, name: 'Paper Plates & Cups', totalNeeded: 1, claimedByOthers: 1, myClaim: 0, category: 'Supplies' }
])

const groupedItems = computed(() => {
  const groups: Record<string, typeof items.value> = {}
  items.value.forEach(item => {
    if (!groups[item.category]) groups[item.category] = []
    groups[item.category]!.push(item)
  })
  return groups
})

const incrementClaim = (item: any) => {
  if (item.claimedByOthers + item.myClaim < item.totalNeeded) {
    item.myClaim++
  }
}

const decrementClaim = (item: any) => {
  if (item.myClaim > 0) {
    item.myClaim--
  }
}

const totalClaimed = (item: any) => item.claimedByOthers + item.myClaim
const isFullyClaimed = (item: any) => totalClaimed(item) >= item.totalNeeded

// Animations
onMounted(() => {
  if (import.meta.client) {
    const tl = gsap.timeline()
    
    tl.fromTo('.event-header-anim', 
      { y: 30, opacity: 0 },
      { y: 0, opacity: 1, duration: 0.8, stagger: 0.15, ease: 'power3.out' }
    )
    
    tl.fromTo('.rsvp-card-anim', 
      { y: 40, opacity: 0 },
      { y: 0, opacity: 1, duration: 0.8, ease: 'power3.out' },
      '-=0.4'
    )
  }
})
</script>

<template>
  <div class="min-h-screen pt-32 pb-20 px-4 md:px-8 lg:px-16 container mx-auto relative overflow-hidden">
    
    <!-- Background abstract blobs to match the premium theme -->
    <div class="absolute top-0 right-0 w-96 h-96 bg-primary opacity-5 rounded-full blur-[100px] pointer-events-none"></div>
    <div class="absolute bottom-0 left-0 w-[500px] h-[500px] bg-secondary opacity-5 rounded-full blur-[120px] pointer-events-none"></div>

    <div class="max-w-5xl mx-auto grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-16 items-start relative z-10">
      
      <!-- Left Column: Event Details (Replacing the simple header) -->
      <div class="lg:col-span-5 flex flex-col justify-center">
        <!-- Event Image in a Pill Mask -->
        <div class="w-full h-48 md:h-64 overflow-hidden pill-mask-horizontal mb-8 shadow-xl event-header-anim">
          <img src="https://images.unsplash.com/photo-1555939594-58d7cb561ad1?q=80&w=1000&auto=format&fit=crop" alt="BBQ Event" class="w-full h-full object-cover" />
        </div>

        <div class="inline-block bg-secondary/10 text-secondary border border-secondary/20 font-medium px-5 py-2 rounded-full mb-6 event-header-anim text-sm backdrop-blur-md w-max">
          You're Invited
        </div>
        <h1 class="text-5xl md:text-6xl lg:text-7xl font-bold tracking-tight text-foreground mb-6 event-header-anim leading-[1.1]">{{ event.title }}</h1>
        <p class="text-muted-foreground text-lg mb-8 leading-relaxed event-header-anim">{{ event.description }}</p>
        
        <div class="flex flex-col gap-4 text-foreground/80 event-header-anim bg-card border border-border p-6 rounded-3xl shadow-sm">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-background flex items-center justify-center shrink-0">
              <Calendar class="w-5 h-5 text-primary" />
            </div>
            <span class="font-medium">{{ event.date }}</span>
          </div>
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-background flex items-center justify-center shrink-0">
              <Clock class="w-5 h-5 text-primary" />
            </div>
            <span class="font-medium">{{ event.time }}</span>
          </div>
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-background flex items-center justify-center shrink-0">
              <MapPin class="w-5 h-5 text-primary" />
            </div>
            <span class="font-medium">{{ event.location }}</span>
          </div>
        </div>
      </div>

      <!-- Right Column: Main Interaction Card -->
      <div class="lg:col-span-7 bg-card rounded-[40px] p-8 md:p-12 shadow-2xl border border-border rsvp-card-anim mt-4 lg:mt-0 relative overflow-hidden">
        <!-- Abstract gradient blob inside card -->
        <div class="absolute -right-20 -top-20 w-64 h-64 bg-primary opacity-[0.03] rounded-full blur-[60px] pointer-events-none"></div>

        
        <!-- Step 1: RSVP -->
        <div v-if="!hasRsvpd" class="space-y-8">
          <div>
            <h2 class="text-2xl font-semibold text-card-foreground mb-2">Are you attending?</h2>
            <p class="text-muted-foreground">Let {{ event.host }} know if you can make it.</p>
          </div>
          
          <div class="flex flex-col sm:flex-row gap-4 relative z-10">
            <button 
              @click="isAttending = true"
              :class="[
                'flex-1 py-4 px-6 rounded-full border-2 transition-all font-semibold flex items-center justify-center gap-2',
                isAttending === true ? 'border-primary bg-primary text-primary-foreground shadow-lg shadow-primary/20 transform scale-[1.02]' : 'border-border text-foreground hover:border-primary/50 bg-background hover:bg-background/80'
              ]"
            >
              Yes, I'm in
            </button>
            <button 
              @click="isAttending = false"
              :class="[
                'flex-1 py-4 px-6 rounded-full border-2 transition-all font-semibold flex items-center justify-center gap-2',
                isAttending === false ? 'border-foreground text-background bg-foreground shadow-lg transform scale-[1.02]' : 'border-border text-foreground hover:border-foreground/50 bg-background hover:bg-background/80'
              ]"
            >
              Can't make it
            </button>
          </div>

          <div v-if="isAttending === true" class="space-y-5 animate-in fade-in slide-in-from-top-4 duration-500 relative z-10">
            <div>
              <label class="text-sm font-medium text-muted-foreground ml-4 mb-2 block">What's your name?</label>
              <input 
                v-model="guestName"
                type="text" 
                placeholder="John Doe" 
                class="w-full bg-background rounded-full px-6 py-4 outline-none border border-border focus:border-primary transition-colors text-card-foreground shadow-sm" 
              />
            </div>
            <button 
              @click="submitRsvp"
              :disabled="!guestName"
              class="w-full py-4 rounded-full bg-secondary text-secondary-foreground font-semibold hover:bg-secondary/90 transition-all active:scale-[0.98] disabled:opacity-50 disabled:pointer-events-none shadow-lg shadow-secondary/20"
            >
              Confirm RSVP
            </button>
          </div>
          
          <div v-if="isAttending === false" class="animate-in fade-in slide-in-from-top-4 duration-500 relative z-10">
             <button 
              @click="submitRsvp"
              class="w-full py-4 rounded-full bg-muted text-muted-foreground font-semibold hover:bg-muted/80 transition-all active:scale-[0.98]"
            >
              Send Apologies
            </button>
          </div>
        </div>

        <!-- Step 2: Item Claiming -->
        <div v-else-if="isAttending" class="space-y-8 animate-in fade-in duration-500 relative z-10">
          <div class="flex items-center justify-between border-b border-border pb-6">
            <div>
              <h2 class="text-2xl font-semibold text-card-foreground mb-1">Awesome, {{ guestName }}! 🎉</h2>
              <p class="text-muted-foreground">What would you like to bring?</p>
            </div>
            <div class="w-12 h-12 bg-primary/20 rounded-full flex items-center justify-center text-primary">
              <Check class="w-6 h-6" />
            </div>
          </div>

          <div class="space-y-8">
            <div v-for="(catItems, category) in groupedItems" :key="category">
              <h3 class="text-lg font-medium text-foreground mb-4 pl-2">{{ category }}</h3>
              <div class="space-y-3">
                
                <div v-for="item in catItems" :key="item.id" 
                  :class="[
                    'p-4 rounded-3xl border transition-all flex flex-col sm:flex-row sm:items-center justify-between gap-4',
                    item.myClaim > 0 ? 'border-primary bg-primary/5 shadow-inner' : 'border-border bg-background shadow-sm hover:border-foreground/20',
                    isFullyClaimed(item) && item.myClaim === 0 ? 'opacity-50 grayscale' : ''
                  ]"
                >
                  <div class="flex-1 pl-2">
                    <h4 class="font-semibold text-card-foreground mb-1 text-lg" :class="{ 'line-through text-muted-foreground': isFullyClaimed(item) && item.myClaim === 0 }">{{ item.name }}</h4>
                    <div class="flex items-center gap-3 text-sm text-muted-foreground">
                      <div class="w-full max-w-[120px] h-2.5 rounded-full bg-border overflow-hidden">
                        <div class="h-full bg-primary transition-all duration-500 rounded-full" :style="{ width: `${(totalClaimed(item) / item.totalNeeded) * 100}%` }"></div>
                      </div>
                      <span class="font-medium">{{ totalClaimed(item) }} / {{ item.totalNeeded }} claimed</span>
                    </div>
                  </div>

                  <div class="flex items-center gap-4 shrink-0">
                    <div v-if="isFullyClaimed(item) && item.myClaim === 0" class="text-sm font-medium px-3 py-1 rounded-full bg-border text-muted-foreground">
                      Fully Claimed
                    </div>
                    <div v-else class="flex items-center gap-3 bg-card border border-border rounded-full p-1 shadow-sm">
                      <button 
                        @click="decrementClaim(item)"
                        :disabled="item.myClaim === 0"
                        class="w-8 h-8 rounded-full flex items-center justify-center hover:bg-muted text-foreground transition-colors disabled:opacity-30 disabled:hover:bg-transparent"
                      >
                        <Minus class="w-4 h-4" />
                      </button>
                      <span class="font-semibold w-4 text-center">{{ item.myClaim }}</span>
                      <button 
                        @click="incrementClaim(item)"
                        :disabled="isFullyClaimed(item)"
                        class="w-8 h-8 rounded-full flex items-center justify-center hover:bg-muted text-foreground transition-colors disabled:opacity-30 disabled:hover:bg-transparent"
                      >
                        <Plus class="w-4 h-4" />
                      </button>
                    </div>
                  </div>
                </div>

              </div>
            </div>
          </div>
          
          <div class="pt-8 border-t border-border mt-4">
            <button class="w-full py-4 rounded-full bg-primary text-primary-foreground font-semibold hover:bg-primary/90 transition-all active:scale-[0.98] shadow-lg shadow-primary/20 text-lg">
              Save Contributions
            </button>
          </div>

        </div>

        <!-- Declined State -->
        <div v-else class="text-center py-16 animate-in fade-in zoom-in duration-500 relative z-10">
           <div class="w-20 h-20 bg-muted rounded-full flex items-center justify-center mx-auto mb-6 text-muted-foreground">
              <User class="w-10 h-10" />
           </div>
           <h2 class="text-3xl font-semibold text-card-foreground mb-3">We'll miss you!</h2>
           <p class="text-muted-foreground text-lg">Thanks for letting us know. Your RSVP has been recorded.</p>
        </div>

      </div>

    </div>
  </div>
</template>
