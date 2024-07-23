import 'modern-normalize/modern-normalize.css'
import './assets/main.scss'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createI18n } from 'vue-i18n'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(
  createI18n({
    legacy: false,
    locale: 'en',
    fallbackLocale: 'ja',
    messages: {
      en: {
        app: {
          name: 'Massivebugs Systems',
        },
        login: {
          title: 'Log into @:app.name',
          login: 'Login',
        },
        home: 'Home',
        budgetPlanner: {
          name: 'Cashbunny',
          overview: 'Overview',
          accounts: 'Accounts',
          transactions: 'Transactions',
          balance: 'Balance',
        },
      },
      ja: {
        app: {
          name: 'Massivebugs Systems',
        },
        login: {
          title: '@:app.name にログインする',
          login: 'ログイン',
        },
        home: 'ホーム',
        budgetPlanner: {
          name: 'Cashbunny',
          overview: 'Overview',
          accounts: 'Accounts',
          transactions: 'Transactions',
          balance: '合計',
        },
      },
    },
  }),
)
app.use(router)

app.mount('#app')
