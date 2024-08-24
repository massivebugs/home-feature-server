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
        :units="store.userPreferences?.user_currencies"
        :error-message="validationErrors.amount || validationErrors.currency"
        v-model:value="formValues.amount"
        v-model:unit="formValues.currency"
      />
      <SelectInputComponent
        name="sourceAccount"
        :label="t('cashbunny.transactionSourceAccount')"
        :options="accounts.map((a) => a.id)"
        :error-message="validationErrors.source_account_id"
        v-model:value="formValues.sourceAccountId"
      />
      <SelectInputComponent
        name="destinationAccount"
        :label="t('cashbunny.transactionDestinationAccount')"
        :options="accounts.map((a) => a.id)"
        :error-message="validationErrors.destination_account_id"
        v-model:value="formValues.destinationAccountId"
      />
      <TextInputComponent
        name="transactedAt"
        :label="t('cashbunny.transactionTransactedAt')"
        :placeholder="t('cashbunny.transactionTransactedAtPlaceholder')"
        :error-message="validationErrors.transacted_at"
        v-model="formValues.transactedAt"
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
  transactedAt: string
}

const emit = defineEmits<{
  (e: 'success'): void
}>()

const props = defineProps<{
  pos: RelativePosition
  size: RelativeSize
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
        transactedAt: props.transaction.transacted_at,
      }
    : {
        description: '',
        amount: 0,
        currency: 'CAD',
        sourceAccountId: 0,
        destinationAccountId: 0,
        transactedAt: '',
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
        transacted_at: formValues.value.transactedAt,
      })
    : store.createTransaction({
        description: formValues.value.description,
        amount: formValues.value.amount,
        currency: formValues.value.currency,
        source_account_id: formValues.value.sourceAccountId,
        destination_account_id: formValues.value.destinationAccountId,
        transacted_at: formValues.value.transactedAt.toString(),
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
