<template>
  <section class="register">
    <h2 class="register__title">Регистрация</h2>
    <input class="register__data" id="login" placeholder="Login">
    <input class="register__data" id="password" placeholder="Password">
    <div class="register__issue">
    <p class="register__issue_text">Уже есть акаунт?</p>
      <nuxt-link class="register__issue_link" to="/login">Вход</nuxt-link>
    </div>
  </section>
</template>

<script lang="ts">
definePageMeta({
  layout: 'empty'
})
</script>

<style>
@import "../assets/register.css";
</style>