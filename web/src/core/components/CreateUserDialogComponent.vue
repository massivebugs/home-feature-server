<template>
  <DialogComponent
    class="hfs-create-user-dialog"
    pos="center"
    :fit-content="true"
    :buttons="{
      success: t('createUser.submit'),
      cancel: true,
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
      {{ t('createUser.title') }}
    </template>
    <div class="hfs-create-user-dialog__container">
      <p class="hfs-create-user-dialog__error-message" v-if="errorMessage">{{ errorMessage }}</p>
      <p class="hfs-create-user-dialog__create-user-disclaimer">
        {{ t('createUser.disclaimer') }}
      </p>
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
        name="email"
        label="Email"
        placeholder="Enter email"
        v-model="email"
        :error-message="validationErrors?.email"
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
import type { CreateUserDto } from '../models/dto'
import type { ValidationErrors } from '../utils/types'
import LockIconComponent from './LockIconComponent.vue'
import TextInputComponent from './TextInputComponent.vue'

export type CreateUserSubmitEvent = {
  username: string
  email: string
  password: string
}

const emit = defineEmits<{
  (e: 'submit', payload: CreateUserSubmitEvent): void
}>()

defineProps<{
  pos?: RelativePosition | 'center'
  disabled?: boolean
  errorMessage?: string
  validationErrors?: ValidationErrors<CreateUserDto>
}>()

const { t } = useI18n()
const username = ref('')
const email = ref('')
const password = ref('')

const onClickSuccess = () => {
  emit('submit', { email: email.value, username: username.value, password: password.value })
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';
.hfs-create-user-dialog {
  transition: box-shadow 0.2s;
}

.hfs-create-user-dialog:hover {
  box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
  -webkit-box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
  -moz-box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
}

.hfs-create-user-dialog__container {
  width: 100%;
  margin-bottom: 1em;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.hfs-create-user-dialog__error-message {
  margin: 0;
  font-size: 0.8em;
  color: colors.$red-cmyk;
}

.hfs-create-user-dialog__create-user-disclaimer {
  white-space: pre-line;
  margin: 0;
  font-size: 0.9em;
  text-align: center;
  line-height: 1.6em;
}
</style>
