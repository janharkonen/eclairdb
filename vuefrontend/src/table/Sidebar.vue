<template>
    <div class="text-white h-full w-full sidebar bg-cyan-800 p-4  overflow-y-auto border-r-2 border-cyan-300">
        <div v-for="schemaName in Object.keys(props.data)" :key="schemaName" class="mb-2">
            <div 
                class="flex items-center cursor-pointer hover:bg-cyan-700 p-1 rounded border border-cyan-600"
                @click="toggleSchema(schemaName)"
            >
                <span class="mr-2 select-none overflow-hidden text-ellipsis whitespace-nowrap">{{ isSchemaOpen[schemaName] ? `▼ ${schemaName}` : `► ${schemaName}` }}</span>
            </div>
            <div v-if="isSchemaOpen[schemaName]" class="ml-4 mt-1">
                <button 
                    v-for="tableName in Object.keys(props.data[schemaName])" 
                    :key="tableName"
                    :class="
                        selectedTable !== tableName || selectedSchema !== schemaName ? 
                        'text-cyan-300 bg-cyan-800 hover:bg-cyan-900' 
                        : 
                        'text-cyan-900 bg-cyan-400 hover:bg-cyan-100'
                    "
                    class="
                        px-2 
                        w-full 
                        text-left 
                        p-1 
                        rounded 
                        cursor-pointer 
                        select-none 
                        overflow-hidden 
                        text-ellipsis 
                        whitespace-nowrap
                        "
                    @click="selectTable(schemaName, tableName)"
                >
                    {{ tableName }}
                </button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { defineProps, ref } from 'vue'

const props = defineProps<{
    data: Record<string, Record<string, string[]>>,
}>()

const isSchemaOpen = ref<Record<string, boolean>>({})

const toggleSchema = (schemaName: string) => {
    isSchemaOpen.value[schemaName] = !isSchemaOpen.value[schemaName]
}

const selectedTable = defineModel<string>('selectedTable')
const selectedSchema = defineModel<string>('selectedSchema')

const selectTable = (schemaName: string, tableName: string) => {
    selectedTable.value = tableName
    selectedSchema.value = schemaName
}

</script>