// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: ['@nuxt/ui'],
  app:{
    pageTransition: { name: 'page', mode: 'out-in' },
  },
  routeRules: {
    '*': { swr: true },
    '/api/**': { cors: true },
    '/incident/*': {ssr: true}
  }
})