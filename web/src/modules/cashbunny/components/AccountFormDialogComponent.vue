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
      <div class="title">{{ t('cashbunny.addAccount') }}</div>
      <TextInputComponent
        name="accountName"
        :label="t('cashbunny.accountName')"
        :placeholder="t('cashbunny.accountNamePlaceholder')"
        :error-message="validationErrors.name"
        v-model="name"
      />
      <SelectInputComponent
        name="accountCategory"
        :label="t('cashbunny.accountCategory')"
        :options="['assets', 'liabilities', 'revenue', 'expenses']"
        :error-message="validationErrors.category"
        v-model:value="category"
      />
      <TextInputComponent
        name="accountDescription"
        :label="t('cashbunny.accountDescription')"
        :placeholder="t('cashbunny.accountDescriptionPlaceholder')"
        :error-message="validationErrors.description"
        v-model="description"
      />
      <NumberInputComponent
        name="accountBalance"
        :label="t('cashbunny.accountBalance')"
        placeholder="0"
        :min="0"
        :units="store.userPreferences?.user_currencies"
        :error-message="validationErrors.balance || validationErrors.currency"
        v-model:value="balance"
        v-model:unit="currency"
      />
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { AxiosError } from 'axios'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { APIResponse } from '@/core/models/dto'
import type { RelativePosition } from '@/core/models/relative_position'
import type { RelativeSize } from '@/core/models/relative_size'
import type { CreateAccountDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'

const emit = defineEmits<{
  (e: 'success'): void
}>()

const props = defineProps<{
  pos: RelativePosition
  size: RelativeSize
  title: string
  nextAccountIndex: number
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const name = ref<string>('')
const category = ref<string>('assets')
const description = ref<string>('')
const balance = ref<number>(0)
const currency = ref<string>('CAD')
const errorMessage = ref<string>('')
const validationErrors = ref<{ [k in keyof CreateAccountDto]: string }>({
  name: '',
  category: '',
  description: '',
  balance: '',
  currency: '',
  order_index: '',
})

const onClickSubmit = async () => {
  await store
    .createAccount({
      name: name.value,
      category: category.value,
      description: description.value,
      balance: balance.value,
      currency: currency.value,
      order_index: props.nextAccountIndex,
    })
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
