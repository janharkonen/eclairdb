<!--
v-on: is the same as @
v-bind: is the same as :
-->
<template>
  <form class="
    input-div
    shadow-xl
    shadow-cyan-600/30
    border-2
    border-cyan-200/40
    rounded-xl
    p-2
    "
    >
    <div class="relative flex w-full h-12 md:h-16">
      <input 
        type="text" 
        :name="service"
        :placeholder="placeholder" 
        class="flex-1 text-black pl-12 md:pl-14 pr-4 py-2 border-2 placeholder:text-black/50 border-cyan-300/50 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-cyan-400 shadow-inner shadow-cyan-200/20"
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
        class="absolute right-2 md:right-3 top-1/2 transform -translate-y-1/2 bg-cyan-500 hover:bg-cyan-600 h-8 w-8 md:h-10 md:w-10 rounded flex items-center justify-center transition-colors duration-200 border border-blue-300/3 hover:cursor-pointer "
        type="button"
        aria-label="Submit"
        :disabled="textisEmpty"
        @click="handleButtonClick"
        :title="`Connect to database using ${service}`"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M10.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L12.586 11H5a1 1 0 110-2h7.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd" />
        </svg>
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { onMounted } from 'vue'
import { useQuery } from '@tanstack/vue-query';

interface Props {
  imgSrc: string
  placeholder: string
  service: string
  apiroute: string
}

const props = defineProps<Props>()

var inputValue = ""
const textisEmpty = ref(true)

onMounted(() => {
  const inputElement = document.querySelector(`input[name="${props.service}"]`)
  if (inputElement) {
    // Set initial value
    inputValue = (inputElement as HTMLInputElement).value
    textisEmpty.value = (inputValue === "")
    
    // Add event listener to update the value when input changes
    inputElement.addEventListener('input', (e) => {
      inputValue = (e.target as HTMLInputElement).value
      textisEmpty.value = (inputValue === "")
    })
  }
  
})
const handleButtonClick = () => {
  console.log('Input value:', inputValue);
}

/*

const { data, isLoading, error } = useQuery({
  queryKey: ['databaseConnect', props.apiroute],
  queryFn: () => fetch(`http://goapi:8081/${props.apiroute}`, {
    method: 'POST',
    body: JSON.stringify({
      uri: document.querySelector('input')?.value
    })
  }).then(res => res.json())
})
*/
</script>

<style scoped>
.input-div {
  display: flex;
  background-color: #d0f6ff;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 2rem;
  width: 100%;
  max-width: 920px;
  
}
button:disabled {
  background-color: #899396;
  cursor: not-allowed;
  opacity: 0.5;
}
</style> 