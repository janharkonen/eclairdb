<template>
  <div v-if="isLoading">Loading...</div>
  <div v-else-if="error">Error: {{ error.message }}</div>
  <div v-else>
    <p v-for="schema in Object.keys(data)">{{ schema }}</p>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useQuery } from '@tanstack/vue-query'
import { computed } from 'vue'
const route = useRoute()
const hash = route.params.hash

const { data, isLoading, error } = useQuery({
  queryKey: ['schemasAndTables', hash],
  queryFn: async () => {
    const response = await fetch(`/api/get-schemas-and-tables?hash=${hash}`);
    
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    
    return response.json();
  },
})

const schemas = computed(() => {
  return Object.keys(data.value)
})
</script>

<style scoped>
p {
  margin: 0;
  color: white;
}
</style>