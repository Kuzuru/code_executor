// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    devtools: {enabled: true},
    modules: [
        ['@nuxtjs/google-fonts', {
            families: {
                Roboto: true,
                Inter: [400, 700],
                'Josefin+Sans': true,
                Lato: [100, 300],
                Raleway: {
                    wght: [100, 400],
                    ital: [100]
                },
                'Crimson Pro': {
                    wght: '200..900',
                    ital: '200..700',
                },
                Manrope: [400, 500, 600, 700]
            }
        }],
    ],
})
