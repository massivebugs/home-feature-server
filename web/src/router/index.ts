import { createRouter, createWebHistory } from 'vue-router'
import DesktopView from '@/core/views/DesktopView.vue'
import SplashView from '@/core/views/SplashView.vue'
import LoginView from '@/core/views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'splash',
      component: SplashView,
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/desktop',
      name: 'desktop',
      component: DesktopView,
    },
  ],
})

export default router
