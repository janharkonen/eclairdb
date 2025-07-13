<template>
  <div class="
    input-div
    ">
    <div class="relative flex w-full h-12 md:h-16">
      <input 
        type="text" 
        :placeholder="placeholder" 
        class="flex-1 pl-12 md:pl-14 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-white focus:border-transparent"
      >
      </input>
      <div class="absolute left-2 top-1/2 transform -translate-y-1/2">
        <img 
          :src="imgSrc"
          :alt="imgAlt" 
          class="h-8 w-8 px-1 py-1 md:h-10 md:w-10 aspect-square opacity-100" 
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  imgSrc: string
  imgAlt: string
  placeholder: string
}

defineProps<Props>()
</script>

<style scoped>
.input-div {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 2rem;
  width: 100%;
  max-width: 920px;
}
</style> 