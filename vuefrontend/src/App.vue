<script setup>
import Homepage from './Homepage.vue';
</script>

<template>
  <main class="w-full h-full flex flex-col items-center justify-center">
    <Homepage />
  </main>
  <footer className="flex-none w-full bg-gray-100 text-gray-800">
    <div className="mx-auto py-4 px-4 sm:px-6 lg:px-8">
      <p className="text-center text-sm text-gray-500">
        © Jan Härkönen 2025
      </p>
    </div>
  </footer>
</template>

<style scoped>
/*
header {
  img {
    max-width: 400px;
    max-height: 400px;
  }
}
*/
</style>