export type BringItSectionId = 'overview' | 'payment' | 'budgets' | 'report' | 'activity'

export type BringItNavId =
  | 'overview'
  | 'transactions'
  | 'payment'
  | 'report'
  | 'budgets'
  | 'ai-insights'
  | 'savings-planner'
  | 'expense-tracker'
  | 'investment-tracker'
  | 'help-center'
  | 'setting'
  | 'profile'

export type BringItNavItem = {
  id: BringItNavId
  icon: string
  label: string
  section: BringItSectionId
}

const generalLinks: BringItNavItem[] = [
  { id: 'overview', label: 'overview', icon: 'grid', section: 'overview' },
  { id: 'transactions', label: 'Transactions', icon: 'transaction', section: 'activity' },
  { id: 'payment', label: 'Payment', icon: 'payment', section: 'payment' },
  { id: 'report', label: 'Report', icon: 'report', section: 'report' },
  { id: 'budgets', label: 'Budgets', icon: 'budget', section: 'budgets' },
]

const toolLinks: BringItNavItem[] = [
  { id: 'ai-insights', label: 'AI Insights', icon: 'star', section: 'report' },
  { id: 'savings-planner', label: 'Savings Planner', icon: 'savings', section: 'budgets' },
  { id: 'expense-tracker', label: 'Expense Tracker', icon: 'tracker', section: 'activity' },
  { id: 'investment-tracker', label: 'Investment Tracker', icon: 'globe', section: 'report' },
]

const footerLinks: BringItNavItem[] = [
  { id: 'help-center', label: 'Help Center', icon: 'help', section: 'activity' },
  { id: 'setting', label: 'Setting', icon: 'setting', section: 'overview' },
  { id: 'profile', label: 'Profile', icon: 'profile', section: 'overview' },
]

const dockLinks: BringItNavItem[] = [
  { id: 'overview', label: 'Overview', icon: 'grid', section: 'overview' },
  { id: 'transactions', label: 'Activity', icon: 'transaction', section: 'activity' },
  { id: 'payment', label: 'Payment', icon: 'payment', section: 'payment' },
]

export function useBringItNavigation() {
  const sidebarLinks = {
    general: generalLinks,
    tools: toolLinks,
    footer: footerLinks,
  }

  const allLinks = [...generalLinks, ...toolLinks, ...footerLinks]

  return {
    allLinks,
    dockLinks,
    sidebarLinks,
  }
}
