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
      <div class="title">{{ t('cashbunny.createTransaction') }}</div>
      <TextInputComponent
        name="transactionDescription"
        :label="t('cashbunny.transactionDescription')"
        :placeholder="t('cashbunny.transactionDescriptionPlaceholder')"
        :error-message="validationErrors.description"
        v-model="description"
      />
      <NumberInputComponent
        name="transactionAmount"
        :label="t('cashbunny.transactionAmount')"
        placeholder="0"
        :min="0"
        :units="['CAD', 'JPY']"
        :error-message="validationErrors.amount || validationErrors.currency"
        v-model:value="amount"
        v-model:unit="currency"
      />
      <SelectInputComponent
        name="sourceAccount"
        :label="t('cashbunny.transactionSourceAccount')"
        :options="accounts.map((a) => a.id)"
        :error-message="validationErrors.source_account_id"
        v-model:value="sourceAccountId"
      />
      <SelectInputComponent
        name="destinationAccount"
        :label="t('cashbunny.transactionDestinationAccount')"
        :options="accounts.map((a) => a.id)"
        :error-message="validationErrors.destination_account_id"
        v-model:value="destinationAccountId"
      />
      <TextInputComponent
        name="transactedAt"
        :label="t('cashbunny.transactionTransactedAt')"
        :placeholder="t('cashbunny.transactionTransactedAtPlaceholder')"
        :error-message="validationErrors.transacted_at"
        v-model="transactedAt"
      />
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { AxiosError } from 'axios'
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { APIResponse } from '@/core/models/dto'
import type { RelativePosition } from '@/core/models/relative_position'
import type { RelativeSize } from '@/core/models/relative_size'
import type { AccountDto, CreateTransactionDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'

const emit = defineEmits<{
  (e: 'success'): void
}>()

defineProps<{
  pos: RelativePosition
  size: RelativeSize
  title: string
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const description = ref<string>('')
const amount = ref<number>(0)
const currency = ref<string>('CAD')
const sourceAccountId = ref<number>(0)
const destinationAccountId = ref<number>(0)
const transactedAt = ref<string>('')
const errorMessage = ref<string>('')
const validationErrors = ref<{ [k in keyof CreateTransactionDto]: string }>({
  description: '',
  amount: '',
  currency: '',
  source_account_id: '',
  destination_account_id: '',
  transacted_at: '',
})
const accounts = ref<AccountDto[]>([])

const onClickSubmit = async () => {
  await store
    .createTransaction({
      description: description.value,
      amount: amount.value,
      currency: currency.value,
      source_account_id: sourceAccountId.value,
      destination_account_id: destinationAccountId.value,
      transacted_at: transactedAt.value.toString(),
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

onMounted(async () => {
  const res = await store.getAccounts()
  if (res.data.error === null) {
    accounts.value = res.data.data
  }
})
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
