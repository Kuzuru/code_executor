<template>
  <section class="header">
    <nuxt-link to="/" class="header__button_home"><img class="header__logo" src="../assets/images/Logo.svg" alt="logo"></nuxt-link>
    <div class="header__buttons">
      <nuxt-link class="header__buttons_history" to="/history">История</nuxt-link>
      <nuxt-link class="header__buttons_profile" to="/profile">
        <img class="header__buttons_profile_img" src="../assets/images/Profile.svg" alt="profile">
      </nuxt-link>
    </div>
  </section>
</template>

<script lang="ts">

</script>

<style>
@import '../assets/Header.css';
</style>