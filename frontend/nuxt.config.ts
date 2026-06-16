import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
  compatibilityDate: '2025-01-01',
  ssr: true,
  routeRules: {
    '/host/**': { ssr: false },
    '/create/**': { ssr: false },
    '/event/**': { ssr: false },
  },
  css: ['~/assets/css/main.css'],
  app: {
    head: {
      htmlAttrs: {
        lang: 'en',
      },
      title: 'BringIt',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1, viewport-fit=cover' },
        { name: 'description', content: 'Coordinate RSVPs and shared contributions for events.' },
        { property: 'og:type', content: 'website' },
        { property: 'og:title', content: 'BringIt' },
        { property: 'og:description', content: 'RSVP, claim what you can bring, and keep event logistics clear.' },
        { name: 'theme-color', content: '#efefef' },
      ],
      link: [
        { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' },
      ],
    },
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || (process.env.NODE_ENV === 'production' ? 'https://api.bringit.example.com' : 'http://localhost:8080'),
    },
  },
  typescript: {
    strict: true,
  },
  components: {
    dirs: [
      {
        path: '~/components',
        pathPrefix: false,
        extensions: ['.vue'],
      },
    ],
  },
  modules: ['shadcn-nuxt', '@vite-pwa/nuxt'],
  shadcn: {
    prefix: '',
    componentDir: './components/ui',
  },
  vite: {
    plugins: [tailwindcss()],
  },
  pwa: {
    registerType: 'autoUpdate',
    manifest: {
      name: 'BringIt',
      short_name: 'BringIt',
      description: 'Coordinate RSVPs and shared event contributions.',
      theme_color: '#efefef',
      background_color: '#efefef',
      display: 'standalone',
      orientation: 'portrait',
      start_url: '/host',
      icons: [
        {
          src: '/favicon.svg',
          sizes: '192x192',
          type: 'image/svg+xml',
          purpose: 'any maskable',
        },
      ],
    },
    workbox: {
      navigateFallback: '/',
      globPatterns: ['**/*.{js,css,html,png,svg,ico}'],
    },
  },
})
