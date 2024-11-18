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
        :label="t('cashbunny.transaction.description')"
        :placeholder="t('cashbunny.transaction.form.descriptionPlaceholder')"
        :error-message="validationErrors.description"
        v-model="formValues.description"
      />
      <NumberInputComponent
        name="transactionAmount"
        :label="t('cashbunny.transaction.amount')"
        placeholder="0"
        :min="0"
        :units="store.userPreference?.userCurrencies"
        :error-message="validationErrors.amount || validationErrors.currency"
        v-model:value="formValues.amount"
        v-model:unit="formValues.currency"
      />
      <SelectInputComponent
        name="sourceAccount"
        :label="t('cashbunny.transaction.sourceAccount')"
        :options="accounts.map((a) => ({ label: a.name, value: a.id }))"
        :error-message="validationErrors.sourceAccountId"
        v-model="formValues.sourceAccountId"
      />
      <SelectInputComponent
        name="destinationAccount"
        :label="t('cashbunny.transaction.destinationAccount')"
        :options="accounts.map((a) => ({ label: a.name, value: a.id }))"
        :error-message="validationErrors.destinationAccountId"
        v-model="formValues.destinationAccountId"
      />
      <DateTimeInputComponent
        name="transactedAt"
        :label="t('cashbunny.transaction.transactedAt')"
        :error-message="validationErrors.transactedAt"
        v-model="formValues.transactedAt"
      />
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DateTimeInputComponent from '@/core/components/DateTimeInputComponent.vue'
import DialogComponent from '@/core/components/DialogComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type {
  API,
  CashbunnyAccountResponse,
  CashbunnyTransactionResponse,
  CreateCashbunnyTransactionRequest,
} from '@/core/composables/useAPI'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'
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
  api: API
  pos?: RelativePosition | 'center'
  size?: RelativeSize
  title: string
  transaction?: CashbunnyTransactionResponse
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const accounts = ref<CashbunnyAccountResponse[]>([])
const formValues = ref<TransactionFormValues>(
  props.transaction
    ? {
        description: props.transaction.description,
        amount: props.transaction.amount,
        currency: props.transaction.currency,
        sourceAccountId: props.transaction.sourceAccountId,
        destinationAccountId: props.transaction.destinationAccountId,
        transactedAt: new Date(props.transaction.transactedAt),
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
const validationErrors = ref<{ [k in keyof CreateCashbunnyTransactionRequest]: string }>({
  description: '',
  amount: '',
  currency: '',
  sourceAccountId: '',
  destinationAccountId: '',
  transactedAt: '',
})

const onClickSubmit = async () => {
  const request = props.transaction
    ? props.api.updateCashbunnyTransaction(
        props.transaction.id,
        {
          description: formValues.value.description,
          amount: formValues.value.amount,
          transactedAt: formValues.value.transactedAt.toISOString(),
        },
        {
          400: () => {
            // Handle in catch below
          },
        },
      )
    : props.api.createCashbunnyTransaction(
        {
          description: formValues.value.description,
          amount: formValues.value.amount,
          currency: formValues.value.currency,
          sourceAccountId: formValues.value.sourceAccountId,
          destinationAccountId: formValues.value.destinationAccountId,
          transactedAt: formValues.value.transactedAt.toISOString(),
        },
        {
          400: () => {
            // Handle in catch below
          },
        },
      )

  try {
    await request
    emit('success')
  } catch (error) {
    if (props.api.isError(error) && error.status === 400) {
      errorMessage.value = error.message
      validationErrors.value = error.validationMessages as any
    } else {
      throw error
    }
  }
}

onMounted(async () => {
  const res = await props.api.getCashbunnyAccounts()
  accounts.value = res.accounts

  if (!props.transaction && accounts.value.length > 0) {
    formValues.value.sourceAccountId = accounts.value[0].id
    formValues.value.destinationAccountId = accounts.value[0].id
    formValues.value.currency = accounts.value[0].currency
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
