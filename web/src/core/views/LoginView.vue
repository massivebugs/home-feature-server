<template>
  <div class="hfs-login-view">
    <LoginDialogComponent
      v-if="showLoginDialog"
      class="hfs-login-view__login-dialog"
      pos="center"
      :disabled="isSubmitting"
      :loading-spinner="isSubmitting && SpinnerTypes.dots"
      :error-message="loginErrorMessage"
      :validation-errors="loginValidationErrors"
      @submit="onSubmitLogin"
      @create-account="onClickShowCreateAccount"
    />
    <CreateAccountDialogComponent
      v-if="showCreateAccountDialog"
      class="hfs-login-view__login-dialog"
      :is-pending-approval="isUserPendingAdminApproval"
      pos="center"
      :disabled="isSubmitting"
      :loading-spinner="isSubmitting && SpinnerTypes.dots"
      :error-message="createUserErrorMessage"
      :validation-errors="createUserValidationErrors"
      @click-cancel="onCancelCreateAccount"
      @submit="onSubmitCreateAccount"
    />
  </div>
</template>

<script setup lang="ts">
import { AxiosError } from 'axios'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import CreateAccountDialogComponent, {
  type CreateAccountSubmitEvent,
} from '../components/CreateAccountDialogComponent.vue'
import LoginDialogComponent, { type LoginSubmitEvent } from '../components/LoginDialogComponent.vue'
import { SpinnerTypes } from '../components/SpinnerIconComponent.vue'
import { AuthUser } from '../models/auth_user'
import type { APIError, CreateAuthTokenDto, CreateUserDto } from '../models/dto'
import { useCoreStore } from '../stores'
import type { ValidationErrors } from '../utils/types'

const coreStore = useCoreStore()
const router = useRouter()
const showLoginDialog = ref<boolean>(false) // Toggles login dialog. Is true when user is not authenticated.
const showCreateAccountDialog = ref<boolean>(false)
const isUserPendingAdminApproval = ref<boolean>(false)
const isSubmitting = ref<boolean>(false) // Used to disable login form temporarily
const loginErrorMessage = ref<string>('')
const loginValidationErrors = ref<ValidationErrors<CreateAuthTokenDto>>({
  username: '',
  password: '',
})
const createUserErrorMessage = ref<string>('')
const createUserValidationErrors = ref<ValidationErrors<CreateUserDto>>({
  email: '',
  username: '',
  password: '',
})

onMounted(() => {
  checkAuth()
})

const onSubmitLogin = async (payload: LoginSubmitEvent) => {
  isSubmitting.value = true

  try {
    await coreStore.createAuthToken(payload)

    // TODO: Error handling when creating refresh token fails
    await coreStore.createRefreshToken()
    checkAuth()
  } catch (error) {
    if (error instanceof AxiosError) {
      if (error.code === AxiosError.ERR_BAD_REQUEST) {
        const res = error.response?.data as APIError
        loginErrorMessage.value = res.error || ''
        loginValidationErrors.value = { username: '', password: '', ...res.validation_messages }
      }
    }
  } finally {
    isSubmitting.value = false
  }
}

const onSubmitCreateAccount = async (payload: CreateAccountSubmitEvent) => {
  isSubmitting.value = true
  try {
    await coreStore.createUser(payload)
    isUserPendingAdminApproval.value = true
  } catch (error) {
    if (error instanceof AxiosError) {
      if (error.code === AxiosError.ERR_BAD_REQUEST) {
        const res = error.response?.data as APIError
        createUserErrorMessage.value = res.error || ''
        createUserValidationErrors.value = {
          email: '',
          username: '',
          password: '',
          ...res.validation_messages,
        }
      }
    }
  } finally {
    isSubmitting.value = false
  }
}

const onClickShowCreateAccount = () => {
  showLoginDialog.value = false
  showCreateAccountDialog.value = true
}

const onCancelCreateAccount = () => {
  showLoginDialog.value = true
  showCreateAccountDialog.value = false
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

.hfs-login-view__login-dialog,
.hfs-login-view__create-account-dialog {
  min-width: 350px;
}
</style>
