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
        v-model="accountName"
      />
      <TextInputComponent
        name="categoryName"
        :label="t('cashbunny.categoryName')"
        :placeholder="t('cashbunny.categoryNamePlaceholder')"
        :list="existingCategoryNames"
        :error-message="validationErrors.category_name"
        v-model="categoryName"
      />
      <TextInputComponent
        name="accountDescription"
        :label="t('cashbunny.accountDescription')"
        :placeholder="t('cashbunny.accountDescriptionPlaceholder')"
        :error-message="validationErrors.description"
        v-model="accountDescription"
      />
      <NumberInputComponent
        name="accountBalance"
        :label="t('cashbunny.accountBalance')"
        placeholder="0"
        :min="0"
        :units="['CAD', 'JPY']"
        :error-message="validationErrors.balance || validationErrors.currency"
        v-model:value="accountBalance"
        v-model:unit="accountCurrency"
      />
      <SelectInputComponent
        name="accountType"
        :label="t('cashbunny.accountType')"
        :options="['debit', 'credit']"
        :error-message="validationErrors.type"
        v-model:value="accountType"
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
const accountName = ref<string>('')
const categoryName = ref<string>('')
const accountDescription = ref<string>('')
const accountBalance = ref<number>(0)
const accountCurrency = ref<string>('CAD')
const accountType = ref<string>('debit')
const existingCategoryNames = ref<string[]>([])
const errorMessage = ref<string>('')
const validationErrors = ref({
  name: '',
  category_name: '',
  description: '',
  balance: '',
  currency: '',
  type: '',
  order_index: '',
})

const onClickSubmit = async () => {
  await store
    .createAccount({
      name: accountName.value,
      category_name: categoryName.value,
      description: accountDescription.value,
      balance: accountBalance.value,
      currency: accountCurrency.value,
      type: accountType.value,
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

onMounted(async () => {
  const res = await store.getAccountCategories()
  if (res.data.error === null) {
    existingCategoryNames.value = res.data.data.map((category) => category.name)
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
