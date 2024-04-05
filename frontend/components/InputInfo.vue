<template>
  <section id="info">
    <div class="info">
      <div class="info__item">
        <h3 class="info__item_title">Язык</h3>
        <p class="info__item_text">Java</p>
      </div>
      <div class="info__item">
        <h3 class="info__item_title">Последний запуск</h3>
        <p class="info__item_text">22.02.2024 13:10:52T+03:00</p>
      </div>
      <div class="info__item">
        <h3 class="info__item_title">Последнее сохранение</h3>
        <p class="info__item_text">21.02.2024 19:23:42T+03:00</p>
      </div>
      <div class="info__item">
        <h3 class="info__item_title">Создано</h3>
        <p class="info__item_text">Kuzuru (ID: 2914)</p>
      </div>
      <div class="info__item">
        <h3 class="info__item_title">Доступ</h3>
        <p class="info__item_text">https://codexecutor.com/iS45ujX2</p>
      </div>
    </div>
    <button class="info__button" @click="runCode">Запустить</button>
  </section>
</template>
<script setup lang="ts">
import runCode from '../pages/executor.vue'
</script>

<style>
@import "../assets/InputInfo.css";
</style>