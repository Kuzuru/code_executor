<template>
  <section class="index">
    <button class="index__button" @click="openPopup">Создать новый файл</button>
  </section>
  <CreatePopup/>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: 'protected'
})
const openPopup = () => {
  const popupElement = document.querySelector('.popup');
  if (popupElement)
    popupElement.classList.remove('hidden');
  const overlayElement = document.querySelector('.overlay');
  if (overlayElement)
    overlayElement.classList.remove('hidden');
}
</script>

<style>
@import "../assets/index.css";
</style>