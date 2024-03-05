// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  router: {
    routes: [
      {
        path: '/',
        component: '@/components/MainPage.vue',
      },
      {
        path: '/history',
        component: '@/components/CodeHistory.vue',
      },
    ],
  }
})
