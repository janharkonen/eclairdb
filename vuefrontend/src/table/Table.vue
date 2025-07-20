<template>
  <div v-if="!shownTable || !shownSchema" class="flex justify-center items-center h-full">
    <p>No table or schema selected</p>
  </div>
  <div v-else class="flex justify-center items-center h-full">
    <p>Table: {{ shownTable }}</p>
    <p>Schema: {{ shownSchema }}</p>
  </div>
</template>

<script setup lang="ts">

const shownTable = defineModel<string>('shownTable')
const shownSchema = defineModel<string>('shownSchema')
</script>