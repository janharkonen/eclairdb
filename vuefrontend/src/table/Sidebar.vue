<template>
    <div class="h-full w-full sidebar bg-cyan-800 p-4 overflow-y-auto border-r-2 border-cyan-300">
        <p v-for="schema in Object.keys(props.data)">{{ schema }}</p>
    </div>
</template>

<script setup lang="ts">
import { defineProps, type PropType } from 'vue'

const props = defineProps({
    data: {
        type: Object as PropType<any>,
        required: true
    }
})
</script>

<style scoped>
p {
  margin: 0;
  color: white;
}
</style>