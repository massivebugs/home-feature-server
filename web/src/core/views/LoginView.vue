<template>
  <div class="login">
    <p>
      {{ t('login.title') }}
    </p>
    <p>
      <input v-model="username" />
    </p>
    <p>
      <input v-model="password" type="password" />
    </p>
    <button @click="onSubmit">
      {{ t('login.login') }}
    </button>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import { AuthUser } from '../models/auth_user'
import { useStore } from '../stores'

const { t } = useI18n()
const store = useStore()
const router = useRouter()
const username = ref('massivebugs')
const password = ref('this_is_meant_to_be_public123')

onMounted(() => {
  checkAuth()
})

const onSubmit = async () => {
  const res = await store.getAuthToken(username.value, password.value)
  if (res.data.error === null) {
    localStorage.setItem('token', res.data.data)
    checkAuth()
  }
}

const checkAuth = async () => {
  const res = await store.getAuthUser()
  if (res.data.error === null) {
    store.authUser = new AuthUser(res.data.data)
    router.push({ name: 'desktop' })
  }
}
</script>

<style scoped lang="scss">
.login {
  width: 100vw;
  height: 100vh;
}
</style>
