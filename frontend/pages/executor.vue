<template>
  <section class="main-page">
    <div class="main-page__input">
      <InputArea />
      <InputInfo />
    </div>
    <OutputConsole />
  </section>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: 'protected'
})
</script>


<style>
@import "../assets/MainPage.css";
</style>