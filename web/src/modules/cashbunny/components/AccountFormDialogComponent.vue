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
        {{ account ? t('cashbunny.editAccount') : t('cashbunny.addAccount') }}
      </div>
      <TextInputComponent
        name="accountName"
        :label="t('cashbunny.account.name')"
        :placeholder="t('cashbunny.account.form.namePlaceholder')"
        :error-message="validationErrors.name"
        v-model="formValues.name"
      />
      <SelectInputComponent
        name="accountCategory"
        :label="t('cashbunny.account.category')"
        :options="[
          { label: 'Assets', value: 'assets' },
          { label: 'Liabilities', value: 'liabilities' },
          { label: 'Revenues', value: 'revenues' },
          { label: 'Expenses', value: 'expenses' },
        ]"
        :error-message="validationErrors.category"
        v-model="formValues.category"
      />
      <TextInputComponent
        name="accountDescription"
        :label="t('cashbunny.account.description')"
        :placeholder="t('cashbunny.account.form.descriptionPlaceholder')"
        :error-message="validationErrors.description"
        v-model="formValues.description"
      />
      <SelectInputComponent
        name="accountCurrency"
        :label="t('cashbunny.transaction.currency')"
        :options="
          cashbunnyStore.userPreference?.userCurrencies.map((v) => ({
            label: `${v} (${cashbunnyStore.currencies[v]})`,
            value: v,
          }))
        "
        :error-message="validationErrors.currency"
        v-model="formValues.currency"
      />
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type {
  API,
  CashbunnyAccountResponse,
  CreateCashbunnyAccountRequest,
} from '@/core/composables/useAPI'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'
import { useCashbunnyStore } from '../stores'

const emit = defineEmits<{
  (e: 'success'): void
}>()

const props = defineProps<{
  api: API
  pos?: RelativePosition | 'center'
  size?: RelativeSize
  title: string
  nextAccountIndex: number
  account?: CashbunnyAccountResponse
}>()

const { t } = useI18n()
const cashbunnyStore = useCashbunnyStore()
const formValues = ref<CreateCashbunnyAccountRequest>(
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
        currency: cashbunnyStore.userPreference?.userCurrencies[0] ?? 'CAD',
      },
)
const errorMessage = ref<string>('')
const validationErrors = ref<{ [k in keyof CreateCashbunnyAccountRequest]: string }>({
  name: '',
  category: '',
  description: '',
  currency: '',
  orderIndex: '',
})

const onClickSubmit = async () => {
  const request = props.account
    ? props.api.updateCashbunnyAccount(
        props.account.id,
        {
          name: formValues.value.name,
          description: formValues.value.description,
          orderIndex: props.account.orderIndex,
        },
        {
          400: () => {
            // Handle in catch below
          },
        },
      )
    : props.api.createCashbunnyAccount(
        {
          name: formValues.value.name,
          category: formValues.value.category,
          description: formValues.value.description,
          currency: formValues.value.currency,
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
