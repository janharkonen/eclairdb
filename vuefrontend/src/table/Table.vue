<template>
  <div class="flex flex-col justify-center items-center h-full">
    <p>Params: {{ params }}</p>
    <p>Data: {{ data }}</p>
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
  const params_indexes = `&--indexes=0-10`
  return params_base + params_indexes
})

const { data, isLoading, error } = useQuery({
  queryKey: ['table', params],
  queryFn: async () => {
    const response = await fetch(`http://localhost:8081/filtered_paginated_products?${params.value}`)
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }
    return response.json()
  },
})

console.log(params)
</script>
