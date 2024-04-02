<template>
  <section class="register">
    <h2 class="register__title">Регистрация</h2>
    <form @submit.prevent="register" class="register__form">
      <input
          v-model="name"
          class="register__form_data"
          id="login"
          placeholder="Логин"
          type="text"
          required
      />
      <input
          v-model="password"
          class="register__form_data"
          id="password"
          placeholder="Пароль"
          type="password"
          required
      />
      <button type="submit" class="register__form_button">Зарегистрироваться</button>
    </form>
    <div class="register__issue">
      <p class="register__issue_text">Уже есть аккаунт?</p>
      <nuxt-link class="register__issue_link" to="/login">Вход</nuxt-link>
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

    const register = async () => {
      try {
        const response = await axios.post('http://localhost:8080/user/register', {
          name: name.value,
          password: password.value,
        });

        if (response.status === 200) {
          console.log('Регистрация прошла успешно:', response.data);
          await router.push({path: "/"})
        } else {
          error.value = 'Регистрация не удалась. Пожалуйста, попробуйте еще раз.';
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
      register,
    };
  },
});
</script>

<style scoped>
@import '../assets/register.css';
</style>
