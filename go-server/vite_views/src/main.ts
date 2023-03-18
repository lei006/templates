import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './style.css'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'


const pinia = createPinia()
let app = createApp(App)

app.use(pinia)
app.use(ElementPlus)

app.mount('#app')
