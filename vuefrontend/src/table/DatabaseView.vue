<template>
  <div v-if="isLoading">Loading...</div>
  <div v-else-if="error">Error: {{ error.message }}</div>
  <div v-else class="h-full flex flex-row w-full">
    <div class="h-full flex"
      ref="sidebarContainer"
      :style="{ width: sidebarWidth + 'px' }"
      >
      <Sidebar 
        class="flex-grow" 
        :schemasAndTables="schemasAndTablesLoading" 
        v-model:selectedTable="table" 
        v-model:selectedSchema="schema" />
    </div>
    <div class="w-1 h-full cursor-ew-resize bg-cyan-300 hover:bg-cyan-400 active:bg-cyan-500" 
      @mousedown="startResize"
    ></div>
    <div class="flex flex-col h-full w-full">
      <div v-if="!table || !schema" class="h-full w-full flex items-center justify-center"> 
        <p>No table selected</p>
      </div>
      <div v-else class="h-full flex flex-col">
        <div 
          class="
          text-white 
          text-xl 
          p-4 
          h-10
          gap-2 
          flex 
          items-center 
          justify-left 
          border-b 
          border-cyan-600
          bg-cyan-900
        " 
        >
          <span class="mr-2 text-gray-500">Schema: </span>
          <span class="text-gray-300">{{ schema }}</span>
          <span class="mr-2 text-gray-500">, Table: </span>
          <span class="text-cyan-300">{{ table }}</span>
        </div>
        <ScrollAreaRoot class="w-full h-full overflow-auto">
          <ScrollAreaViewport class="w-full h-full">
              <Table 
              :shownTable="table" 
              :shownSchema="schema" 
              :hash="hash"
              />
          </ScrollAreaViewport>
          <ScrollAreaScrollbar orientation="horizontal" class="flex touch-none select-none p-0.5 bg-gray-100 transition-colors duration-150 ease-out hover:bg-gray-200 data-[orientation=horizontal]:h-2.5 data-[orientation=vertical]:w-2.5 data-[orientation=horizontal]:flex-col">
            <ScrollAreaThumb class="flex-1 bg-gray-400 rounded-[10px] relative before:content-[''] before:absolute before:top-1/2 before:left-1/2 before:-translate-x-1/2 before:-translate-y-1/2 before:w-full before:h-full before:min-w-[44px] before:min-h-[44px]" />
          </ScrollAreaScrollbar>
          <ScrollAreaScrollbar orientation="vertical" class="flex touch-none select-none p-0.5 bg-gray-100 transition-colors duration-150 ease-out hover:bg-gray-200 data-[orientation=horizontal]:h-2.5 data-[orientation=vertical]:w-2.5 data-[orientation=vertical]:flex-col">
            <ScrollAreaThumb class="flex-1 bg-gray-400 rounded-[10px] relative before:content-[''] before:absolute before:top-1/2 before:left-1/2 before:-translate-x-1/2 before:-translate-y-1/2 before:w-full before:h-full before:min-w-[44px] before:min-h-[44px]" />
          </ScrollAreaScrollbar>
          <ScrollAreaCorner class="bg-gray-200" />
        </ScrollAreaRoot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ScrollAreaRoot, ScrollAreaScrollbar, ScrollAreaThumb, ScrollAreaViewport } from 'radix-vue'
import { useRoute } from 'vue-router'
import { useQuery } from '@tanstack/vue-query'
import Sidebar from './Sidebar.vue'
import Table from './Table.vue'
import { ref } from 'vue'


const sidebarWidth = ref(240)
const schemasAndTablesLoading = ref<Record<string, Record<string, boolean>>>({})
const table = ref<string>('')
const schema = ref<string>('')
const route = useRoute()
const hash = route.params.hash as string

const { data, isLoading, error } = useQuery({
  queryKey: ['schemasAndTables', hash],
  queryFn: async () => {
    const response = await fetch(`/api/get-schemas-and-tables?hash=${hash}`);
    
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    
    const data = await response.json();
    schemasAndTablesLoading.value = data;
    return data;
  },
})

schemasAndTablesLoading.value = data.value

const sidebarContainer = ref<HTMLElement | null>(null)

const startResize = (event: MouseEvent) => {
  const startX = event.clientX
  const startWidth = sidebarContainer.value?.clientWidth ?? 0;

  const handleMouseMove = (moveEvent: MouseEvent) => {
    const currentX = moveEvent.clientX
    const currentWidth = (startWidth ?? 0) + (currentX - startX)
    sidebarWidth.value = Math.max(currentWidth, 120)
  }

  const handleMouseUp = () => {
    document.removeEventListener('mousemove', handleMouseMove)
    document.removeEventListener('mouseup', handleMouseUp)
  }

  document.addEventListener('mousemove', handleMouseMove)
  document.addEventListener('mouseup', handleMouseUp)
}
const eventSource = new EventSource(`/api/get-schemas-and-tables-stream?hash=${hash}`)

import { onMounted } from 'vue'
onMounted(() => {
  
  eventSource.addEventListener('table_ready', (event: MessageEvent) => {
    const [schema, table] = event.data.split(':')
    schemasAndTablesLoading.value[schema][table] = true
  })

  eventSource.addEventListener('complete', () => { 
    console.log('complete')
    for (const schema in schemasAndTablesLoading.value) {
      for (const table in schemasAndTablesLoading.value[schema]) {
        schemasAndTablesLoading.value[schema][table] = true
      }
    }
    eventSource.close() 
  })

  eventSource.onopen = () => {
    console.log('Connected to SSE')
}

// Add error handler to close connection on error
eventSource.onerror = (error) => {
    console.error('SSE error:', error)
    eventSource.close()
  }

})

window.addEventListener('beforeunload', () => {
  if (eventSource) {
    console.log('Closing EventSource before page refresh')
    eventSource.close()
  }
})

</script>

<style scoped>
p {
  margin: 0;
  color: white;
}
</style>