<template>
  <div v-if="isLoading && !data">
    Loading... 
  </div>
  <div v-else-if="isFetching && !data">
    Fetching...
  </div>
  <div v-else-if="error && !data">
    Error:
  </div>
  <div v-else-if="data" class="flex flex-col h-full p-4">
    <ScrollAreaRoot class="w-full h-full overflow-auto rounded-t-xl border-x-4 border-t-4 border-cyan-400 bg-gray-100">
      <ScrollAreaViewport class="w-full h-full">
        <table :style="{ tableLayout: 'fixed', width: `${totalWidth}px` }">
          <thead id="header" class="sticky top-0 bg-gray-200 z-10">
            <tr class="w-full">
              <th 
              v-for="(column, index) in data.columnList"
              :key="`header-${column}`" 
              :style="{ 
                height: `${headerHeight}px`, 
              }"
              class="border-b border-gray-300 text-left relative"
              >
              <!--Filter input component-->
              <div class="p-1 pr-1.5 flex flex-col h-full w-full">
                <div class="flex-none top-0 left-0 text-sm font-sans font-bold"
                :style="{fontSize: '16px', fontWeight: 'bold', fontFamily: 'sans-serif'} ">
                  {{ column }}
                </div>
                <input :list="`${column}-options`" 
                :id="`${column}-input`" 
                class="w-full h-full p-2 border-2  border-cyan-200 rounded-md bg-white" 
                :value="filterValues.get(column) || ''"
                @input="(e: Event) => setFilter(e, column)"
                autocomplete="off"
                />
                <datalist :id="`${column}-options`">
                  <option v-for="option in Object.keys(data.columnOptions[column])" :value="option"/>
                </datalist>
              </div>
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
              </th>
            </tr>
          </thead>
          <tbody>
            <tr 
            v-for="row in data.rowList"
            :style="{ 
              height: `${rowHeight}px`,
              maxHeight: `${rowHeight}px`,
              minHeight: `${rowHeight}px`,
              overflow: 'hidden',
              textOverflow: 'ellipsis',
              lineHeight: `${rowHeight}px`
            }"
            class="bg-white hover:bg-cyan-200"
            >
              <td 
                v-for="(column, index) in data.columnList"
                class="truncate cursor-pointer border-b border-cyan-400 font-sans"
                :style="{ 
                  padding: '0 4px',
                  boxSizing: 'border-box',
                  fontSize: '12px',
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
    <div className="flex-none h-12 w-full rounded-b-xl bg-gray-200 border-b-4 border-x-4 border-cyan-400 flex items-center justify-between px-6 shadow-sm">
      <div className="flex items-center gap-2">
        <span className="text-sm font-medium text-gray-700">
          Showing 
          <span className="font-semibold">{{data.pageInfo.indexStart}}</span> 
          to 
          <span className="font-semibold">{{data.pageInfo.indexEnd}}</span> 
          of 
          <span className="font-semibold">{{data.pageInfo.numOfRows}}</span> results
        </span>
      </div>
      <div className="flex items-center gap-3 text-cyan-800 ">
        <button :class="buttonClasses" 
        :disabled="pageNumber === 1"
        @click="pageNumber = 1">First</button>
        <button :class="buttonClasses" 
        :disabled="pageNumber === 1"
        @click="pageNumber--">Previous</button>
        <button :class="buttonClasses" 
        :disabled="pageNumber === Math.ceil(data.pageInfo.numOfRows / defaultPageSize)"
        @click="pageNumber++">Next</button>
        <button :class="buttonClasses" 
        :disabled="pageNumber === Math.ceil(data.pageInfo.numOfRows / defaultPageSize)"
        @click="pageNumber = Math.ceil(data.pageInfo.numOfRows / defaultPageSize)">Last</button>
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
            bg-white
            hover:bg-gray-100
            disabled:opacity-40 
            disabled:cursor-not-allowed 
            disabled:hover:bg-white 
            disabled:hover:border-gray-300 
            cursor-pointer
            `

import { ScrollAreaRoot, ScrollAreaScrollbar, ScrollAreaThumb, ScrollAreaViewport, ScrollAreaCorner } from 'radix-vue'
import { useQuery } from '@tanstack/vue-query'
import { computed, reactive, watch, ref } from 'vue'

const { shownTable, shownSchema, hash } = defineProps<{
  shownTable: string
  shownSchema: string
  hash: string
}>()

const filterValues = reactive<Map<string, string>>(new Map())
const pageNumber = ref(1)
const defaultPageSize = 50
const params = computed(() => {
  const params_base = `--hash=${hash}&--schema=${shownSchema}&--table=${shownTable}`
  const params_indexes = `&--indexes=${(pageNumber.value - 1) * defaultPageSize + 1}-${pageNumber.value * defaultPageSize}`
  var params_filters = ""
  filterValues.forEach((value, key) => {
    if (value !== "") {
      params_filters += `&${key}=${value}`
    }
  })
  return params_base + params_indexes + params_filters
})

const { data, isLoading, error, isFetching } = useQuery({
  queryKey: ['table', params],
  queryFn: async () => {
    const response = await fetch(`/api/filtered_paginated_rows?${params.value}`)
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`)
    }
    return response.json()
  },
  placeholderData: (previousData) => previousData,
})

const rowCount = computed(() => data.value?.rowList?.length || 0);
const rowHeight = 20;
const headerHeight = 65;
const totalHeight = computed(() => rowCount.value * (rowHeight + 1) + headerHeight)

const columnWidths = reactive<number[]>([150])
const totalWidth = reactive({ value: 150 })


watch(data, (newData) => {
  if (!newData || newData.rowList.length === 0) return
  const columns = newData.columnList
  const columnCount = columns.length
  if (columnWidths.length > columnCount) {
    columnWidths.splice(0, columnWidths.length, ...columnWidths.slice(0, columnCount))
  } else {
    columnWidths.push(...Array(columnCount - columnWidths.length).fill(150))
  }
  totalWidth.value = columnWidths.reduce((acc, width) => acc + width, 0)
})

watch([() => shownTable, () => shownSchema], ([newTable, newSchema]) => {
  filterValues.clear()
  data.value?.columnList.forEach((column: string) => {
    filterValues.set(column, "")
  })
  pageNumber.value = 1
})
//console.log("data", data.value)
const startResize = (index: number, e: MouseEvent) => {
  const startX = e.clientX
  const startWidth = columnWidths[index]

  const handleMouseMove = (e: MouseEvent) => {
    const deltaX = e.clientX - startX
    const newWidth = Math.max(20, startWidth + deltaX)

    // Update the reactive array
    columnWidths[index] = newWidth
    totalWidth.value = totalWidth.value + deltaX
  }

  const handleMouseUp = () => {
    document.removeEventListener("mousemove", handleMouseMove)
    document.removeEventListener("mouseup", handleMouseUp)
  }

  document.addEventListener("mousemove", handleMouseMove)
  document.addEventListener("mouseup", handleMouseUp)
  console.log("filterValues", filterValues)
}

const setFilter = (e: Event, column: string) => {
  filterValues.set(column, (e.target as HTMLInputElement).value)
  pageNumber.value = 1
}

</script>