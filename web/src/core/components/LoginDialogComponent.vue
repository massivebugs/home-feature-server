<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :buttons="{
      success: t('login.login'),
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
    @click-success="onClickSuccess"
  >
    <template #title>
      <LockIconComponent />
      {{ t('login.title') }}
    </template>
    <div class="hfs-login-dialog__container">
      <p class="hfs-login-dialog__error-message" v-if="errorMessage">{{ errorMessage }}</p>
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
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'
import type { CreateAuthTokenDto } from '../models/dto'
import type { ValidationErrors } from '../utils/types'
import LockIconComponent from './LockIconComponent.vue'
import TextInputComponent from './TextInputComponent.vue'

export type LoginSubmitEvent = {
  username: string
  password: string
}

const emit = defineEmits<{
  (e: 'submit', payload: LoginSubmitEvent): void
}>()

defineProps<{
  pos: RelativePosition
  size: RelativeSize
  disabled?: boolean
  errorMessage?: string
  validationErrors?: ValidationErrors<CreateAuthTokenDto>
}>()

const { t } = useI18n()
const username = ref('testuser')
const password = ref('test_password_123')

const onClickSuccess = () => {
  emit('submit', { username: username.value, password: password.value })
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-login-dialog__container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.hfs-login-dialog__error-message {
  margin: 0;
  font-size: 0.8em;
  color: colors.$red-cmyk;
}
</style>
