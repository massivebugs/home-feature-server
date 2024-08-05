import 'modern-normalize/modern-normalize.css'
import './assets/main.scss'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'

import App from './App.vue'
import router from './router'
import { i18nOptions } from './i18n'

const app = createApp(App)

app.use(createPinia())
app.use(createI18n(i18nOptions))
app.use(router)

app.mount('#app')
