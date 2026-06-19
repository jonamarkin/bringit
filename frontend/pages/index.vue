<script setup lang="ts">
import type { BringItNavId, BringItSectionId } from '~/composables/useBringItNavigation'

const mobileDrawerOpen = ref(false)
const activeNav = ref<BringItNavId>('overview')

const overviewSection = ref<HTMLElement | null>(null)
const paymentSection = ref<HTMLElement | null>(null)
const budgetsSection = ref<HTMLElement | null>(null)
const reportSection = ref<HTMLElement | null>(null)
const activitySection = ref<HTMLElement | null>(null)

const { allLinks } = useBringItNavigation()

if (import.meta.client) {
  const toggleScrollLock = (locked: boolean) => {
    const value = locked ? 'hidden' : ''
    document.documentElement.style.overflow = value
    document.body.style.overflow = value
  }

  watch(mobileDrawerOpen, (open) => {
    toggleScrollLock(open)
  })

  onBeforeUnmount(() => {
    toggleScrollLock(false)
  })
}

const sectionRefs: Record<BringItSectionId, typeof overviewSection> = {
  overview: overviewSection,
  payment: paymentSection,
  budgets: budgetsSection,
  report: reportSection,
  activity: activitySection,
}

const navigateTo = (id: BringItNavId) => {
  const item = allLinks.find((link) => link.id === id)

  if (!item) {
    return
  }

  activeNav.value = id
  mobileDrawerOpen.value = false

  nextTick(() => {
    sectionRefs[item.section].value?.scrollIntoView({
      behavior: 'smooth',
      block: 'start',
    })
  })
}

const metricCards = [
  {
    title: 'Total Income',
    amount: '$28,982.00',
    change: '20.0% ↗',
    changeClass: 'positive',
    helper: 'vs 2,293.00 last month',
  },
  {
    title: 'Total Expense',
    amount: '$29,249.00',
    change: '20.0% ↘',
    changeClass: 'negative',
    helper: 'vs 2,293.00 last month',
  },
  {
    title: 'Total Saving',
    amount: '$15,340.00',
    change: '20.0% ↗',
    changeClass: 'positive',
    helper: 'vs 2,293.00 last month',
  },
]

const activityRows = [
  {
    type: 'Received',
    typeClass: 'received',
    amount: '+$1,982',
    methodTitle: 'Credit Card',
    methodSubtitle: '•••• 9129',
    personName: 'Mikolin Slavana',
    personClass: 'light',
    date: 'July 24, 2024 - 12:00 PM',
  },
  {
    type: 'Sent',
    typeClass: 'sent',
    amount: '-$1,982',
    methodTitle: 'Stripe',
    methodSubtitle: '@alexander',
    personName: 'Ann Baker',
    personClass: 'dark',
    date: 'July 24, 2024 - 12:00 PM',
  },
]
</script>

<template>
  <div class="page-shell">
    <button
      v-if="mobileDrawerOpen"
      class="mobile-drawer-scrim"
      type="button"
      aria-label="Close navigation"
      @click="mobileDrawerOpen = false"
    />

    <FlowSidebar
      :active-nav="activeNav"
      :mobile-open="mobileDrawerOpen"
      @close="mobileDrawerOpen = false"
      @navigate="navigateTo"
    />

    <main class="dashboard">
      <section ref="overviewSection" class="dashboard-section">
        <FlowTopbar />
        <FlowIntroToolbar />
      </section>

      <section ref="paymentSection" class="dashboard-section">
        <FlowBalanceBanner />
      </section>

      <section ref="budgetsSection" class="metric-grid dashboard-section">
        <FlowMetricCard
          v-for="card in metricCards"
          :key="card.title"
          :title="card.title"
          :amount="card.amount"
          :change="card.change"
          :change-class="card.changeClass"
          :helper="card.helper"
        />
      </section>

      <section ref="reportSection" class="content-grid dashboard-section">
        <FlowChartCard />
        <FlowCashflowCard />
      </section>

      <section ref="activitySection" class="activity-section dashboard-section">
        <div class="card-header table-header">
          <div>
            <h2>Recent Activity</h2>
          </div>
        </div>

        <div class="table-card">
          <table class="activity-table">
            <thead>
              <tr>
                <th>Type</th>
                <th>Amount</th>
                <th>Payment Method</th>
                <th>Status</th>
                <th>People</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="row in activityRows" :key="`${row.type}-${row.personName}`">
                <td>
                  <div class="type-cell">
                    <span class="type-icon" :class="row.typeClass">{{ row.typeClass === 'received' ? '↘' : '↗' }}</span>
                    <span>{{ row.type }}</span>
                  </div>
                </td>
                <td>{{ row.amount }}</td>
                <td>
                  <div class="method-cell">
                    <strong>{{ row.methodTitle }}</strong>
                    <span>{{ row.methodSubtitle }}</span>
                  </div>
                </td>
                <td><span class="status-chip">✓ Success</span></td>
                <td>
                  <div class="person-cell">
                    <span class="person-avatar" :class="row.personClass"></span>
                    <span>{{ row.personName }}</span>
                  </div>
                </td>
                <td>{{ row.date }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </main>

    <FlowMobileDock
      :active-nav="activeNav"
      :more-open="mobileDrawerOpen"
      @more="mobileDrawerOpen = true"
      @navigate="navigateTo"
    />
  </div>
</template>
