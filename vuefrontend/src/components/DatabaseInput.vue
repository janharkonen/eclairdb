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
          :alt="service + ' Logo'"
          class="h-8 w-8 px-1 py-1 md:h-10 md:w-10 aspect-square opacity-100" 
        />
      </div>
      <button 
        class="absolute right-2 md:right-3 top-1/2 transform -translate-y-1/2 bg-blue-500 hover:bg-blue-600 text-white h-8 w-8 md:h-10 md:w-10 rounded flex items-center justify-center transition-colors duration-200"
        type="button"
        aria-label="Submit"
        :title="`Connect to database using ${service}`"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M10.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L12.586 11H5a1 1 0 110-2h7.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd" />
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  imgSrc: string
  placeholder: string
  service: string
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