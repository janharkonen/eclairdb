<template>
  <div class="w-full h-full">
          <table class="w-full" :style="{ width: `${totalWidth}px` }">
            <thead id="header" class="sticky top-0 bg-white z-10">
              <tr>
                <th 
                v-for="(column, index) in columns"
                :key="`header-${column}`" 
                :style="{ width: `${columnWidths[index]}px`, height: `${rowHeight}px` }"
                class="text-left"
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
          <!--
          -->
        </table>
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

import { reactive } from 'vue'

const rows = computed(() => data.value || [])
const columns = computed(() => {
  if (!rows.value || rows.value.length === 0) return []
  return Object.keys(rows.value[0])
})
const rowCount = computed(() => data.value?.length || 0);
const rowHeight = 21;
const headerHeight = computed(() => document.getElementById('header')?.clientHeight || 0)
const totalHeight = computed(() => rowCount.value * rowHeight + headerHeight.value)
const columnWidths = computed(() => {
  return columns.value.map(() => 100)
})
const totalWidth = computed(() => columnWidths.value.reduce((acc, width) => acc + width, 0))


import { watch } from 'vue'

watch(data, (newData) => {
  console.log("totalWidth", totalWidth.value)
  console.log("rows", rows.value)
  console.log("columns", columns.value)
  console.log("newData", newData)
  console.log("columnWidths", columnWidths.value)
})
//console.log("data", data.value)

</script>
