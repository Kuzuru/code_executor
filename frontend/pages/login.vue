<template>
  <section class="login">
    <h2 class="login__title">Вход</h2>
    <form @submit.prevent="login" class="login__form">
      <input
          v-model="name"
          class="login__form_data"
          id="login"
          placeholder="Логин"
          type="text"
          required
      />
      <input
          v-model="password"
          class="login__form_data"
          id="password"
          placeholder="Пароль"
          type="password"
          required
      />
      <button type="submit" class="login__form_button">Войти</button>
    </form>
    <div class="login__issue">
      <p class="login__issue_text">Еще нет аккаунта?</p>
      <nuxt-link class="login__issue_link" to="/register">Регистрация</nuxt-link>
    </div>
  </section>
</template>

<script setup lang="ts">
definePageMeta({
  layout: 'empty',
})
const name = ref();
const password = ref();
const userStore = useUserStore()
const router = useRouter()

const login = async () => {
  await userStore.login({
    name: name.value,
    password: password.value
  })

    await router.push('/')
}


</script>

<style scoped>
@import "../assets/login.css";
</style>
