<template>
  <DialogComponent
    class="hfs-login-dialog"
    pos="center"
    :fit-content="true"
    :buttons="{
      success: t('login.submit'),
      cancel: false,
    }"
    :controls="{
      close: false,
      minimize: false,
      maximize: false,
    }"
    :blocking="false"
    :resizable="false"
    :disabled="disabled"
    form="hfs-login-form"
    @click-success="onClickSuccess"
  >
    <template #title>
      <LockIconComponent />
      {{ t('login.title') }}
    </template>
    <div class="hfs-login-dialog__container">
      <p class="hfs-login-dialog__error-message" v-if="errorMessage">{{ errorMessage }}</p>
      <p class="hfs-login-dialog__create-user-message">
        {{ t('createUser.message') }}
        <a href="#" @click="emit('createUser')">{{ t('createUser.linkTitle') }}</a>
      </p>
      <form id="hfs-login-form" @submit="onFormSubmit">
        <TextInputComponent
          :disabled="disabled"
          name="username"
          label="Username"
          placeholder="Enter username"
          v-model="username"
          :error-message="validationErrors?.username"
          autocomplete="off"
        />
        <TextInputComponent
          :disabled="disabled"
          type="password"
          name="password"
          label="Password"
          placeholder="Enter password"
          v-model="password"
          :error-message="validationErrors?.password"
          autocomplete="off"
        />
      </form>
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { CreateJWTTokenRequest } from '../composables/useAPI'
import type { ValidationErrors } from '../utils/types'
import LockIconComponent from './LockIconComponent.vue'
import TextInputComponent from './TextInputComponent.vue'

export type LoginSubmitEvent = {
  username: string
  password: string
}

const emit = defineEmits<{
  (e: 'submit', payload: LoginSubmitEvent): void
  (e: 'createUser'): void
}>()

defineProps<{
  pos?: RelativePosition | 'center'
  disabled?: boolean
  errorMessage?: string
  validationErrors?: ValidationErrors<CreateJWTTokenRequest>
}>()

const { t } = useI18n()
const username = ref('public')
const password = ref('')

const onClickSuccess = () => {
  emit('submit', { username: username.value, password: password.value })
}

const onFormSubmit = () => {
  console.log('foo')
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';
.hfs-login-dialog {
  transition: box-shadow 0.2s;
}

.hfs-login-dialog:hover {
  box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
  -webkit-box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
  -moz-box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
}

.hfs-login-dialog__container {
  width: 100%;
  margin-bottom: 1em;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.hfs-login-dialog__error-message {
  margin: 0;
  font-size: 0.8em;
  color: colors.$red-cmyk;
}

.hfs-login-dialog__create-user-message {
  margin: 0;
  font-size: 0.8em;
}
</style>
