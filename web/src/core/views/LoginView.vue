<template>
  <div class="hfs-login-view">
    <LoginDialogComponent
      v-if="showLoginDialog"
      class="hfs-login-view__login-dialog"
      pos="center"
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
import type { APIError, CreateAuthTokenDto } from '../models/dto'
import { useCoreStore } from '../stores'
import type { ValidationErrors } from '../utils/types'

const coreStore = useCoreStore()
const router = useRouter()
const showLoginDialog = ref<boolean>(false) // Toggles login dialog. Is true when user is not authenticated.
const isSubmitting = ref<boolean>(false) // Used to disable login form temporarily
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
    await coreStore.getAuthToken(payload)
    checkAuth()
  } catch (error) {
    if (error instanceof AxiosError) {
      if (error.code === AxiosError.ERR_BAD_REQUEST) {
        const res = error.response?.data as APIError
        console.log(res)
        errorMessage.value = res.error || ''
        validationErrors.value = { username: '', password: '', ...res.validation_messages }
      }
    }
  } finally {
    isSubmitting.value = false
  }
}

// Checks whether user has been authenticated (has JWT token)
// and redirects to desktop if so. Otherwise, displays login dialog.
const checkAuth = async () => {
  try {
    const res = await coreStore.getAuthUser()
    coreStore.authUser = new AuthUser(res.data)
    router.push({ name: 'desktop' })
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
