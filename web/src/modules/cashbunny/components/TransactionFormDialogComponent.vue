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
    :resizable="true"
    @click-success="onClickSubmit"
  >
    <div class="container">
      <div class="title">
        {{ transaction ? t('cashbunny.editTransaction') : t('cashbunny.createTransaction') }}
      </div>
      <TextInputComponent
        name="transactionDescription"
        :label="t('cashbunny.transactionDescription')"
        :placeholder="t('cashbunny.transactionDescriptionPlaceholder')"
        :error-message="validationErrors.description"
        v-model="formValues.description"
      />
      <NumberInputComponent
        name="transactionAmount"
        :label="t('cashbunny.transactionAmount')"
        placeholder="0"
        :min="0"
        :units="store.userPreference?.userCurrencies"
        :error-message="validationErrors.amount || validationErrors.currency"
        v-model:value="formValues.amount"
        v-model:unit="formValues.currency"
      />
      <SelectInputComponent
        name="sourceAccount"
        :label="t('cashbunny.transactionSourceAccount')"
        :options="accounts.map((a) => ({ label: a.name, value: a.id }))"
        :error-message="validationErrors.source_account_id"
        v-model="formValues.sourceAccountId"
      />
      <SelectInputComponent
        name="destinationAccount"
        :label="t('cashbunny.transactionDestinationAccount')"
        :options="accounts.map((a) => ({ label: a.name, value: a.id }))"
        :error-message="validationErrors.destination_account_id"
        v-model="formValues.destinationAccountId"
      />
      <DateTimeInputComponent
        name="transactedAt"
        :label="t('cashbunny.transactionTransactedAt')"
        :error-message="validationErrors.transacted_at"
        v-model="formValues.transactedAt"
      />
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { AxiosError } from 'axios'
import dayjs from 'dayjs'
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DateTimeInputComponent from '@/core/components/DateTimeInputComponent.vue'
import DialogComponent from '@/core/components/DialogComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { APIResponse } from '@/core/models/dto'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'
import type { AccountDto, CreateTransactionDto, TransactionDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'

export type TransactionFormValues = {
  description: string
  amount: number
  currency: string
  sourceAccountId: number
  destinationAccountId: number
  transactedAt: Date
}

const emit = defineEmits<{
  (e: 'success'): void
}>()

const props = defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
  title: string
  transaction?: TransactionDto
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const formValues = ref<TransactionFormValues>(
  props.transaction
    ? {
        description: props.transaction.description,
        amount: props.transaction.amount,
        currency: props.transaction.currency,
        sourceAccountId: props.transaction.source_account_id,
        destinationAccountId: props.transaction.destination_account_id,
        transactedAt: new Date(props.transaction.transacted_at),
      }
    : {
        description: '',
        amount: 0,
        currency: 'CAD',
        sourceAccountId: 0,
        destinationAccountId: 0,
        transactedAt: new Date(),
      },
)
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
  const request = props.transaction
    ? store.updateTransaction(props.transaction.id, {
        description: formValues.value.description,
        amount: formValues.value.amount,
        transacted_at: dayjs(formValues.value.transactedAt).toISOString(),
      })
    : store.createTransaction({
        description: formValues.value.description,
        amount: formValues.value.amount,
        currency: formValues.value.currency,
        source_account_id: formValues.value.sourceAccountId,
        destination_account_id: formValues.value.destinationAccountId,
        transacted_at: dayjs(formValues.value.transactedAt).toISOString(),
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

onMounted(async () => {
  const res = await store.getAccounts()
  if (res.data.error === null) {
    accounts.value = res.data.data
  }
})
</script>

<style scoped lang="scss">
.container {
  min-width: 300px;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
  margin-bottom: 1em;
}

.title {
  text-align: center;
  font-weight: 500;
}
</style>
