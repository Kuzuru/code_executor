<template>
  <section class="login">
    <h2 class="login__title">Вход</h2>
    <input class="login__data" id="login" placeholder="Login">
    <input class="login__data" id="password" placeholder="Password">
    <div class="login__issue">
      <p class="login__issue_text">Еще нет акаунта?</p>
      <nuxt-link class="login__issue_link" to="/register">Регистрация</nuxt-link>
    </div>
  </section>
</template>

<script lang="ts">
definePageMeta({
  layout: 'empty'
})
</script>

<style>
@import "../assets/login.css";
</style>