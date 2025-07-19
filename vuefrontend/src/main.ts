import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import { VueQueryPlugin } from '@tanstack/vue-query'
import { createMemoryHistory, createRouter } from 'vue-router'

import Homepage from './Homepage.vue'
import Table from './Table.vue'

const routes = [
  { path: '/', component: Homepage },
  { path: '/table', component: Table },
]
const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

const app = createApp(App)
app.use(VueQueryPlugin)
app.use(router)
app.mount('#app')
