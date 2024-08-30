<template>
  <div class="hfs-login-view">
    <LoginDialogComponent
      pos="center"
      class="hfs-login-view__login-dialog"
      v-if="showLoginDialog"
      :disabled="isSubmitting"
      :loading-spinner="isSubmitting && SpinnerTypes.dots"
      :error-message="errorMessage"
      :validation-errors="validationErrors"
      @submit="onSubmit"
    />
  </div>
</template>

<script setup lang="ts">
import { AxiosError } from 'axios'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import LoginDialogComponent, { type LoginSubmitEvent } from '../components/LoginDialogComponent.vue'
import { SpinnerTypes } from '../components/SpinnerIconComponent.vue'
import { AuthUser } from '../models/auth_user'
import type { APIResponse, CreateAuthTokenDto } from '../models/dto'
import { useCoreStore } from '../stores'
import type { ValidationErrors } from '../utils/types'

const store = useCoreStore()
const router = useRouter()
const showLoginDialog = ref<boolean>(false)
const isSubmitting = ref<boolean>(false)
const errorMessage = ref<string>('')
const validationErrors = ref<ValidationErrors<CreateAuthTokenDto>>({
  username: '',
  password: '',
})

onMounted(() => {
  checkAuth()
})

const onSubmit = async (payload: LoginSubmitEvent) => {
  isSubmitting.value = true
  try {
    const res = await store.getAuthToken(payload)
    if (res.data.error === null) {
      localStorage.setItem('token', res.data.data)
      checkAuth()
    }
  } catch (error) {
    if (error instanceof AxiosError) {
      if (error.code === AxiosError.ERR_BAD_REQUEST) {
        const res = error.response?.data as APIResponse<null>
        errorMessage.value = res.error?.message || ''
        validationErrors.value = { ...validationErrors.value, ...res.error?.validation_errors }
      }
    }
  } finally {
    isSubmitting.value = false
  }
}

const checkAuth = async () => {
  try {
    const res = await store.getAuthUser()
    if (res.data.error === null) {
      store.authUser = new AuthUser(res.data.data)
      router.push({ name: 'desktop' })
    }
  } catch (_) {
    showLoginDialog.value = true
  }
}
</script>

<style scoped lang="scss">
.hfs-login-view {
  width: 100vw;
  height: 100vh;
}

.hfs-login-view__login-dialog {
  min-width: 350px;
}
</style>
