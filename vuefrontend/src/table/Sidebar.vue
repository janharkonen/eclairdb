<template>
    <div class="text-white h-full w-full sidebar bg-cyan-800 p-4  overflow-y-auto border-r-2 border-cyan-300">
        
        <div v-for="schemaName in Object.keys(props.data)" :key="schemaName" class="mb-2">
            <div 
                class="flex items-center cursor-pointer hover:bg-cyan-700 p-1 rounded"
                @click="toggleSchema(schemaName)"
            >
                <span class="mr-2 select-none overflow-hidden text-ellipsis whitespace-nowrap">{{ isSchemaOpen[schemaName] ? `▼ ${schemaName}` : `► ${schemaName}` }}</span>
            </div>
            <div v-if="isSchemaOpen[schemaName]" class="ml-4 mt-1">
                <div 
                    v-for="tableName in Object.keys(props.data[schemaName])" 
                    :key="tableName"
                    class="px-2 text-cyan-300 hover:bg-cyan-700 p-1 rounded cursor-pointer select-none overflow-hidden text-ellipsis whitespace-nowrap"
                >
                    {{ tableName }}
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, type PropType, ref } from 'vue'

const props = defineProps({
    data: {
        type: Object as PropType<any>,
        required: true
    }
})

const isSchemaOpen = ref<Record<string, boolean>>({})

const toggleSchema = (schemaName: string) => {
    isSchemaOpen.value[schemaName] = !isSchemaOpen.value[schemaName]
}
</script>