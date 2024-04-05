<template>
  <section class="history">
    <HistoryItem/>
    <HistoryItem/>
    <HistoryItem/>
  </section>
</template>

<script lang="ts">
definePageMeta({
  middleware: 'protected'
})
</script>

<style>
@import "../assets/History.css";
</style>