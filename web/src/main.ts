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
    locale: 'en',
    fallbackLocale: 'ja',
    messages: {
      en: {
        home: 'Home',
        budgetPlanner: {
          name: 'Budget Planner',
          overview: 'Overview',
          accounts: 'Accounts',
          transactions: 'Transactions',
          incomes: 'Incomes',
          expenses: 'Expenses',
          balance: 'Balance',
        },
      },
      ja: {
        home: 'ホーム',
        budgetPlanner: {
          name: '家計簿',
          overview: 'Overview',
          accounts: 'Accounts',
          transactions: 'Transactions',
          incomes: 'Incomes',
          expenses: 'Expenses',
          balance: '合計',
        },
      },
    },
  }),
)
app.use(router)

app.mount('#app')
