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
        autocomplete="off"
        v-model="inputText"
        :name="service"
        :placeholder="placeholder" 
        class="flex-1 text-xs text-black px-12 md:px-14 pr-4 py-2 border-2 placeholder:text-black/50 border-cyan-300/50 rounded-lg focus:outline-none focus:ring-2 focus:ring-cyan-400 focus:border-cyan-400 shadow-inner shadow-cyan-200/20"
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
        :disabled="textisEmpty || isPending"
        @click="handleButtonClick"
        :title="`Connect to database using ${service}`"
      >
        <ArrowRight v-if="!isPending" />
        <Loader2 v-else class="animate-spin" />
      </button>
      <div v-if="showError && error" class="absolute bg-red-200 rounded-md right-0 top-[120%] transform text-xs text-red-500 px-2">
        {{ error.message }}
      </div>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ArrowRight, Loader2 } from 'lucide-vue-next'
import { useMutation } from '@tanstack/vue-query';
import { useRouter } from 'vue-router';

interface Props {
  imgSrc: string
  placeholder: string
  service: string
  apiroute: string
}

const props = defineProps<Props>()

const inputText = ref("")
const textisEmpty = computed(() => (inputText.value.trim() === ""))
const showError = ref(false)
const router = useRouter()

const { mutate, error, isPending } = useMutation({
  mutationFn: async (uri: string) => {
    const response = await fetch(`/api/${props.apiroute}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ uri: uri })
    });
    if (!response.ok) {
      const errorData = await response.json().catch(() => ({ error: 'Network error' }));
      throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
    }
    return response.json();
  },
  onSuccess: (data: string) => {
    router.push(`/database/${data}`);
  },
  onError: (error: Error) => {
    console.error('Database connection failed:', error);
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
  }
})

const handleButtonClick = async () => {
  if (textisEmpty.value) return;
  
  try {
    mutate(inputText.value);
  } catch (err) {
    console.log(err)
    console.error('Error connecting to database:', err);
  }
}
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