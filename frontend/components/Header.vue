<template>
  <header class="header">
    <img class="header__image" src="../assets/images/Logo.svg" alt="logo">
    <div class="header__buttons">
      <button class="header__buttons_history">История</button>
      <button class="header__buttons_profile">
        <img class="header__buttons_profile_img" src="../assets/images/Profile.svg" alt="profile">
      </button>
    </div>
  </header>
</template>

<script lang="ts">
export default {
  name: 'HeaderComponent',
  props: {},
  components: {}
}
</script>

<style>
@import '../assets/Header.css';
</style>