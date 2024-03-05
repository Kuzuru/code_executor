<template>
  <section class="main-page">
    <div class="main-page__input">
      <InputArea/>
      <InputInfo/>
    </div>
    <OutputConsole/>
  </section>
</template>
<script setup lang="ts">
import InputInfo from "~/components/InputInfo.vue";
</script>

<style>
@import "../assets/MainPage.css";
</style>