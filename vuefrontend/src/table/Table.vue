<template>
  <div v-if="isLoading">
    Loading...
  </div>
  <div v-else-if="isFetching">
    Fetching...
  </div>
  <div v-else-if="error">
    Error:
  </div>
  <div v-else-if="data && data.length > 0" class="flex flex-col h-full">
    <ScrollAreaRoot class="w-full h-full overflow-auto">
      <ScrollAreaViewport class="w-full h-full">
        <table :style="{ tableLayout: 'fixed' }">
          <thead id="header" class="sticky top-0 bg-white z-10">
            <tr>
              <th 
              v-for="(column, index) in Object.keys(data[0])"
              :key="`header-${column}`" 
              :style="{ 
                height: `${headerHeight}px`, 
              }"
              class="border-b border-green-200 text-left relative"
              >
              <!--Resize handle-->
              <div
                  className="absolute 
                  top-0 
                  right-0 w-0.5 
                  hover:w-1 
                  cursor-col-resize 
                  bg-cyan-300 
                  hover:bg-cyan-400
                  h-full
                  "
                  :style="{ height: `${totalHeight}px`}"
                  @mousedown="(e) => startResize(index, e)"

                /> 
                <div class="p-2 truncate">
                  {{ column }}
                </div>
              </th>
            </tr>
          </thead>
          <tbody>
            <tr 
            v-for="row in data"
            :style="{ 
              height: `${rowHeight}px`,
              maxHeight: `${rowHeight}px`,
              minHeight: `${rowHeight}px`,
              overflow: 'hidden',
              textOverflow: 'ellipsis',
              lineHeight: `${rowHeight}px`
            }"
            class="bg-cyan-200 hover:bg-cyan-300"
            >
              <td 
                v-for="(column, index) in Object.keys(data[0])"
                class="overflow-hidden truncate whitespace-nowrap cursor-pointer border-b border-cyan-400"
                :style="{ 
                  padding: '0 4px',
                  boxSizing: 'border-box',
                  fontSize: '14px',
                  width: `${columnWidths[index]}px`,
                  maxWidth: `${columnWidths[index]}px`,
                  minWidth: `${columnWidths[index]}px`,
                }"
                :title="row[column]"
              >
                {{ row[column] }}
              </td>
            </tr>
          </tbody>
          <!--
          -->
        </table>
      </ScrollAreaViewport>
      <ScrollAreaScrollbar orientation="horizontal" class="z-50 flex touch-none select-none p-0.5 bg-gray-100 transition-colors duration-150 ease-out hover:bg-gray-200 data-[orientation=horizontal]:h-2.5 data-[orientation=vertical]:w-2.5 data-[orientation=horizontal]:flex-col">
        <ScrollAreaThumb class="flex-1 bg-gray-400 rounded-[10px] relative before:content-[''] before:absolute before:top-1/2 before:left-1/2 before:-translate-x-1/2 before:-translate-y-1/2 before:w-full before:h-full before:min-w-[44px] before:min-h-[44px]" />
      </ScrollAreaScrollbar>
      <ScrollAreaScrollbar orientation="vertical" class="z-50 flex touch-none select-none p-0.5 bg-gray-100 transition-colors duration-150 ease-out hover:bg-gray-200 data-[orientation=horizontal]:h-2.5 data-[orientation=vertical]:w-2.5 data-[orientation=horizontal]:flex-col">
        <ScrollAreaThumb class="flex-1 bg-gray-400 rounded-[10px] relative before:content-[''] before:absolute before:top-1/2 before:left-1/2 before:-translate-x-1/2 before:-translate-y-1/2 before:w-full before:h-full before:min-w-[44px] before:min-h-[44px]" />
      </ScrollAreaScrollbar>
      <ScrollAreaCorner class="z-50 bg-gray-200 touch-none select-none" />
    </ScrollAreaRoot>
    <!--Pagination section-->
    <div className="flex-none h-12 w-full bg-cyan-400 border-t-2 border-cyan-500 flex items-center justify-between px-6 shadow-sm">
      <div className="flex items-center gap-2">
        <span className="text-sm font-medium text-gray-700">
          Showing 
          <span className="font-semibold">{{1}}</span> 
          to 
          <span className="font-semibold">{{50}}</span> 
          of 
          <span className="font-semibold">{{42069}}</span> results
        </span>
      </div>
      <div className="flex items-center gap-3 text-cyan-800 ">
        <button 
          :class="buttonClasses"
        >
          First
        </button>
        <button 
          :class="buttonClasses"
        >
          Previous
        </button>
        <button
          :class="buttonClasses"
        >
          Next
        </button>
        <button
          :class="buttonClasses"
        >
          Last
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">

const buttonClasses : string = `
            px-4 
            py-2 
            text-sm 
            font-medium 
            border 
            border-cyan-200 
            rounded-md 
            bg-cyan-100
            hover:bg-cyan-200
            disabled:opacity-40 
            disabled:cursor-not-allowed 
            disabled:hover:bg-white 
            disabled:hover:border-gray-300 
            cursor-pointer
            `

import { ScrollAreaRoot, ScrollAreaScrollbar, ScrollAreaThumb, ScrollAreaViewport, ScrollAreaCorner } from 'radix-vue'
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

const { data, isLoading, error, isFetching } = useQuery({
  queryKey: ['table', params],
  queryFn: async () => {
    const response = await fetch(`/api/filtered_paginated_products?${params.value}`)
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }
    return response.json()
  },
})

import { reactive, watch } from 'vue'

const columns = computed(() => {
  if (!data.value || data.value.length === 0) return []
  return Object.keys(data.value[0])
})
const rowCount = computed(() => data.value?.length || 0);
const rowHeight = 20;
const headerHeight = 60;
const totalHeight = computed(() => rowCount.value * (rowHeight + 1) + headerHeight)

const columnWidths = reactive<number[]>([150])
const totalWidth = reactive({ value: 150 })




watch(data, (newData) => {
  console.log("columnWidths: ", columnWidths)
  console.log("totalWidth", totalWidth.value)
  console.log("columns", columns.value)
  console.log("newData", newData)
})

watch(data, (newColumns) => {
  if (!data.value || data.value.length === 0) return
  const columns = Object.keys(newColumns[0])
  const columnCount = columns.length
  console.log("old columnWidths", columnWidths)
  console.log("columnCount", columnCount)
  if (columnWidths.length > columnCount) {
    columnWidths.splice(0, columnWidths.length, ...columnWidths.slice(0, columnCount))
  } else {
    columnWidths.push(...Array(columnCount - columnWidths.length).fill(150))
    console.log("columnWidths", columnWidths)
  }
  console.log("columnCount", columnCount)
  totalWidth.value = columnWidths.reduce((acc, width) => acc + width, 0)
  console.log("new columnWidths", columnWidths)
})
//console.log("data", data.value)
const startResize = (index: number, e: MouseEvent) => {
  e.preventDefault()

  const startX = e.clientX
  const startWidth = columnWidths[index]

  const handleMouseMove = (e: MouseEvent) => {
    const deltaX = e.clientX - startX
    const newWidth = Math.max(100, startWidth + deltaX)

    // Update the reactive array
    columnWidths[index] = newWidth
    totalWidth.value = totalWidth.value + deltaX
  }

  const handleMouseUp = () => {
    document.removeEventListener("mousemove", handleMouseMove)
    document.removeEventListener("mouseup", handleMouseUp)
    document.body.style.cursor = ""
    document.body.style.userSelect = ""
  }

  document.addEventListener("mousemove", handleMouseMove)
  document.addEventListener("mouseup", handleMouseUp)
  document.body.style.cursor = "col-resize"
  document.body.style.userSelect = "none"
}

</script>