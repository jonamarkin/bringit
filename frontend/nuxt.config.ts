export default defineNuxtConfig({
  ssr: true,
  devtools: { enabled: true },
  css: ['~/assets/styles/main.css'],
  compatibilityDate: '2025-01-01',
  app: {
    head: {
      title: 'BringIt',
      htmlAttrs: {
        lang: 'en',
      },
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1, maximum-scale=1, viewport-fit=cover' },
        { name: 'theme-color', content: '#1d6f39' },
        { name: 'mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-status-bar-style', content: 'default' },
      ],
      link: [
        { rel: 'manifest', href: '/manifest.webmanifest' },
        { rel: 'icon', type: 'image/svg+xml', href: '/bringit-icon.svg' },
      ],
    },
  },
})
