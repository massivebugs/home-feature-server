<template>
  <div class="hfs-login-view">
    <LoginDialogComponent
      v-if="currDialog === Dialogs.login"
      class="hfs-login-view__dialog"
      :disabled="isSubmitting"
      :loading-spinner="isSubmitting && SpinnerTypes.dots"
      :error-message="loginErrorMessage"
      :validation-errors="loginValidationErrors"
      @submit="onSubmitLogin"
      @create-user="onClickShowCreateUser"
    />
    <CreateUserDialogComponent
      v-else-if="currDialog === Dialogs.createUser"
      class="hfs-login-view__dialog"
      :disabled="isSubmitting"
      :loading-spinner="isSubmitting && SpinnerTypes.dots"
      :error-message="createUserErrorMessage"
      :validation-errors="createUserValidationErrors"
      @click-cancel="onCancelDialog"
      @submit="onSubmitCreateUser"
    />
    <UserPendingAdminApprovalDialogComponent
      v-else-if="currDialog === Dialogs.pendingApproval"
      class="hfs-login-view__dialog"
      @click-cancel="onCancelDialog"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import type { CreateUserSubmitEvent } from '../components/CreateUserDialogComponent.vue'
import CreateUserDialogComponent from '../components/CreateUserDialogComponent.vue'
import LoginDialogComponent, { type LoginSubmitEvent } from '../components/LoginDialogComponent.vue'
import { SpinnerTypes } from '../components/SpinnerIconComponent.vue'
import UserPendingAdminApprovalDialogComponent from '../components/UserPendingAdminApprovalDialogComponent.vue'
import { useAPI } from '../composables/useAPI'
import { API_URL } from '../constants'
import type { CreateAuthTokenDto, CreateUserDto } from '../models/dto'
import { User } from '../models/user'
import { useCoreStore } from '../stores'
import type { ValidationErrors } from '../utils/types'

const Dialogs = {
  nothing: 'nothing',
  login: 'login',
  createUser: 'createUser',
  pendingApproval: 'pendingApproval',
} as const
type Dialog = (typeof Dialogs)[keyof typeof Dialogs]

const coreStore = useCoreStore()
const router = useRouter()
const api = useAPI(API_URL)
const currDialog = ref<Dialog>(Dialogs.nothing)
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
    await api.createJWTToken(payload, {
      400: (error) => {
        loginErrorMessage.value = error.message
        loginValidationErrors.value = { username: '', password: '', ...error.validation_messages }
      },
      403: (error) => {
        onClickShowCreateUser()
        loginErrorMessage.value = error.message
        currDialog.value = Dialogs.pendingApproval
      },
    })

    // TODO: Error handling when creating refresh token fails
    // Maybe invalidate token
    await api.createJWTRefreshToken({
      403: (error) => {
        onClickShowCreateUser()
        loginErrorMessage.value = error.message
        currDialog.value = Dialogs.pendingApproval
      },
    })

    checkAuth()
  } catch {
    // TODO
  } finally {
    isSubmitting.value = false
  }
}

const onSubmitCreateUser = async (payload: CreateUserSubmitEvent) => {
  isSubmitting.value = true

  try {
    await api.createUser(payload, {
      400: (error) => {
        createUserErrorMessage.value = error.message
        createUserValidationErrors.value = {
          email: '',
          username: '',
          password: '',
          ...error.validation_messages,
        }
      },
    })
    currDialog.value = Dialogs.pendingApproval
  } catch {
    // TODO
  } finally {
    isSubmitting.value = false
  }
}

const onClickShowCreateUser = () => {
  currDialog.value = Dialogs.createUser
}

const onCancelDialog = () => {
  currDialog.value = Dialogs.login
  loginErrorMessage.value = ''
}

// Checks whether user has been authenticated (has JWT token)
// and redirects to desktop if so. Otherwise, displays login dialog.
const checkAuth = async () => {
  try {
    const res = await api.getUser()
    coreStore.user = new User(res.data.user)
    router.push({ name: 'desktop' })
  } catch {
    currDialog.value = Dialogs.login
  }
}
</script>

<style scoped lang="scss">
.hfs-login-view {
  width: 100vw;
  height: 100vh;
}

.hfs-login-view__dialog {
  min-width: 350px;
}
</style>
