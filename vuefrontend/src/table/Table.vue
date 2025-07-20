<template>
  <div class="flex justify-center items-center h-full">
    <p>Hash: {{ hash }}</p>
    <p>Table: {{ shownTable }}</p>
    <p>Schema: {{ shownSchema }}</p>
  </div>
</template>

<script setup lang="ts">
import { useQuery } from '@tanstack/vue-query'
import { computed } from 'vue'

const { shownTable, shownSchema, hash } = defineProps<{
  shownTable: string
  shownSchema: string
  hash: string
}>()

const params = computed(() => {
  const params_base = `hash=${hash}&schema=${shownSchema}&table=${shownTable}`
  return params_base
})

const { data, isLoading, error } = useQuery({
  queryKey: ['table', hash, shownSchema, shownTable],
  queryFn: async () => {
    const response = await fetch(`/api/filtered_paginated_products?${params.value}`)
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }
    return response.json()
  },
})

console.log(data)
</script>