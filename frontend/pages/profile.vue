<template>
  <section class="profile">
    <div class="profile__personal">
      <h2 class="profile__title">Личные данные</h2>
      <h3 class="profile__subtitle">Имя</h3>

      <p class="profile__data">Иван</p>
      <h3 class="profile__subtitle">Логин</h3>
      <p class="profile__data">Ivan228</p>
    </div>

    <div class="profile__safety">
      <h2 class="profile__title">Безопасность</h2>
      <h3 class="profile__subtitle">Пароль</h3>
      <p class="profile__data" id="password">******</p>
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

<style>
@import '../assets/Profile.css';
</style>