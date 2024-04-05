import {useUserStore} from "~/stores/auth";

export default defineNuxtRouteMiddleware((to, from) => {
    const userStore = useUserStore()
    const token = userStore.token
    const route = useRoute()

    if (route.path !== '/login' && token === undefined) {
        return navigateTo('/login')
    }
})