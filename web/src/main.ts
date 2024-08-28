import * as dayjs from 'dayjs'
import * as objectSupport from 'dayjs/plugin/objectSupport'
import 'modern-normalize/modern-normalize.css'
import { createPinia } from 'pinia'
import { createApp } from 'vue'
import { createI18n } from 'vue-i18n'
import App from './App.vue'
import './assets/main.scss'
import { i18nOptions } from './i18n'
import router from './router'

dayjs.extend(objectSupport.default)

const app = createApp(App)

app.use(createPinia())
app.use(createI18n(i18nOptions))
app.use(router)
// app.use(mediaQuery, {
//   breakpoints: {
//     xl: 1200,
//     l: 992,
//     m: 768,
//     s: 576,
//   },
// })

app.mount('#app')
