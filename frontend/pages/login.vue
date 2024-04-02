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

<script lang="ts">
import { defineComponent } from 'vue';
import axios from 'axios';

export default defineComponent({
  setup() {
    const name = ref('');
    const password = ref('');
    const error = ref('');
    const router = useRouter()
    const login = async () => {
      try {
        const response = await axios.post('http://localhost:8080/user/auth', {
          name: name.value,
          password: password.value,
        });

        if (response.status === 200) {
          console.log('Авторизация прошла успешно:', response.data);
          await router.push({path: "/"})
        } else {
          error.value = 'Неверный логин или пароль.';
        }
      } catch (err) {
        if (err instanceof Error) {
          error.value = err.message || 'Произошла ошибка при авторизации.';
        } else {
          console.error('Неожиданный тип ошибки:', err);
        }
      }
    };

    return {
      name,
      password,
      error,
      login,
    };
  },
});
</script>

<style scoped>
@import "../assets/login.css";
</style>
