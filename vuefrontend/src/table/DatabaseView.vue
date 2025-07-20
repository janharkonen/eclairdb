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
        :data="data" 
        v-model:selectedTable="table" 
        v-model:selectedSchema="schema" />
    </div>
    <div class="w-1 h-full cursor-ew-resize bg-cyan-300 hover:bg-cyan-400 active:bg-cyan-500" 
      @mousedown="startResize"
    ></div>
    <div class="flex-1 h-full">
      <div v-if="!table || !schema" class="flex justify-center items-center h-full">
        <p>No table selected</p>
      </div>
      <Table 
        v-else
        :shownTable="table" 
        :shownSchema="schema" 
        :hash="hash"
      />
      <!-- Content area will go here -->
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useQuery } from '@tanstack/vue-query'
import Sidebar from './Sidebar.vue'
import Table from './Table.vue'
import { ref } from 'vue'


const sidebarWidth = ref(180)
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
    
    return response.json();
  },
})

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

</script>

<style scoped>
p {
  margin: 0;
  color: white;
}
</style>