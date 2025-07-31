<template>
  <div class="w-full h-full">
          <p class="whitespace-nowrap">
            <span class="inline-block">
              Your table content will go here
              Your table content will go here
              Your table content will go here
              Your table content will go here
              Your table content will go here
            </span>
          </p>
    <!--
      <table class="w-full">
        <thead id="header" class="sticky top-0 bg-white z-10">
          <tr>
            <th 
              v-for="(column, index) in columns"
              :key="`header-${column}`" 
              class="border-b border-gray-200 text-left relative"
              :style="{ width: `${columnWidths[index]}px` }"
              >
              {{ column }}
            </th>
          </tr>
      </thead>
      <tbody>
        <tr 
            v-for="row in rows"
            :style="{ width: `${totalWidth}px`, height: `${rowHeight}px` }"
          >
          <td 
                v-for="(column, index) in columns"
                :style="{ 
                  width: `${columnWidths[index]}px`, 
                  height: `${rowHeight}px`,
                  maxWidth: `${columnWidths[index]}px`,
                  minWidth: `${columnWidths[index]}px`,
                  padding: '0 4px',
                  boxSizing: 'border-box',
                  fontSize: '14px'
                }"
                class="truncate overflow-hidden whitespace-nowrap"
                :title="row[column]"
                >
                {{ row[column] }}
              </td>
            </tr>
          </tbody>
        </table>
        -->
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
  const params_indexes = `&--indexes=1-50`
  return params_base + params_indexes
})

const { data, isLoading, error } = useQuery({
  queryKey: ['table', params],
  queryFn: async () => {
    const response = await fetch(`/api/filtered_paginated_products?${params.value}`)
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }
    return response.json()
  },
})

import { ref } from 'vue'

const rows = computed(() => data.value)
const columns = computed(() => Object.keys(rows.value[0]))
const rowCount = computed(() => data.value.length);
const rowHeight = 21;
const headerHeight = computed(() => document.getElementById('header')?.clientHeight || 0)
const totalHeight = computed(() => rowCount.value * rowHeight + headerHeight.value)
const columnWidths = ref<number[]>([])
const totalWidth = computed(() => columnWidths.value.reduce((acc, width) => acc + width, 0))




console.log(params)
</script>
