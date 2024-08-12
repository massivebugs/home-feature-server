<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :title="title"
    :buttons="{
      success: t('cashbunny.save'),
      cancel: true,
    }"
    :blocking="true"
    :is-resizable="true"
    @click-success="onClickSubmit"
  >
    <div class="container">
      <div class="title">
        {{ account ? t('cashbunny.editAccount') : t('cashbunny.addAccount') }}
      </div>
      <TextInputComponent
        name="accountName"
        :label="t('cashbunny.accountName')"
        :placeholder="t('cashbunny.accountNamePlaceholder')"
        :error-message="validationErrors.name"
        v-model="formValues.name"
      />
      <SelectInputComponent
        name="accountCategory"
        :label="t('cashbunny.accountCategory')"
        :options="['assets', 'liabilities', 'revenues', 'expenses']"
        :error-message="validationErrors.category"
        v-model:value="formValues.category"
      />
      <TextInputComponent
        name="accountDescription"
        :label="t('cashbunny.accountDescription')"
        :placeholder="t('cashbunny.accountDescriptionPlaceholder')"
        :error-message="validationErrors.description"
        v-model="formValues.description"
      />
      <SelectInputComponent
        name="accountCurrency"
        :label="t('cashbunny.transactionCurrency')"
        :options="store.userPreferences?.user_currencies"
        :error-message="validationErrors.currency"
        v-model:value="formValues.currency"
      />
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { AxiosError } from 'axios'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { APIResponse } from '@/core/models/dto'
import type { RelativePosition } from '@/core/models/relative_position'
import type { RelativeSize } from '@/core/models/relative_size'
import type { AccountDto, CreateAccountDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'

export type AccountFormValues = {
  name: string
  category: string
  description: string
  currency: string
}

const emit = defineEmits<{
  (e: 'success'): void
}>()

const props = defineProps<{
  pos: RelativePosition
  size: RelativeSize
  title: string
  nextAccountIndex: number
  account?: AccountDto
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const formValues = ref<AccountFormValues>(
  props.account
    ? {
        name: props.account.name,
        category: props.account.category,
        description: props.account.description,
        currency: props.account.currency,
      }
    : {
        name: '',
        category: 'assets',
        description: '',
        currency: 'CAD',
      },
)
const errorMessage = ref<string>('')
const validationErrors = ref<{ [k in keyof CreateAccountDto]: string }>({
  name: '',
  category: '',
  description: '',
  currency: '',
  order_index: '',
})

const onClickSubmit = async () => {
  const request = props.account
    ? store.updateAccount(props.account.id, {
        name: formValues.value.name,
        description: formValues.value.description,
      })
    : store.createAccount({
        name: formValues.value.name,
        category: formValues.value.category,
        description: formValues.value.description,
        currency: formValues.value.currency,
        order_index: props.nextAccountIndex,
      })

  await request
    .then(() => {
      emit('success')
    })
    .catch((error: AxiosError) => {
      if (error.code === AxiosError.ERR_BAD_REQUEST) {
        const res = error.response?.data as APIResponse<null>
        errorMessage.value = res.error?.message || ''
        validationErrors.value = { ...validationErrors.value, ...res.error?.validation_errors }
      }
    })
}
</script>

<style scoped lang="scss">
.container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.title {
  text-align: center;
  font-weight: 500;
}
</style>
