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
    @click-success="onClickSubmit"
  >
    <div class="container">
      <div class="title">{{ t('cashbunny.addAccount') }}</div>
      <TextInputComponent
        name="accountName"
        :label="t('cashbunny.accountName')"
        :placeholder="t('cashbunny.accountNamePlaceholder')"
        v-model="accountName"
      />
      <TextInputComponent
        name="categoryName"
        :label="t('cashbunny.categoryName')"
        :placeholder="t('cashbunny.categoryNamePlaceholder')"
        v-model="categoryName"
      />
      <TextInputComponent
        name="accountDescription"
        :label="t('cashbunny.accountDescription')"
        :placeholder="t('cashbunny.accountDescription')"
        v-model="accountDescription"
      />
      <NumberInputComponent
        name="accountBalance"
        :label="t('cashbunny.accountBalance')"
        placeholder="0"
        :min="0"
        :units="['CAD', 'JPY']"
        v-model:value="accountBalance"
        v-model:unit="accountBalanceUnit"
      />
      <SelectInputComponent
        name="accountType"
        :label="t('cashbunny.accountType')"
        :options="['debit', 'credit']"
        v-model:value="accountType"
      />
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { RelativePosition } from '@/core/models/relative_position'
import type { RelativeSize } from '@/core/models/relative_size'

defineProps<{
  pos: RelativePosition
  size: RelativeSize
  title: string
}>()

const { t } = useI18n()
const accountName = ref<string>()
const categoryName = ref<string>()
const accountDescription = ref<string>()
const accountBalance = ref<number>(0)
const accountBalanceUnit = ref<string>('CAD')
const accountType = ref<string>('debit')

const onClickSubmit = () => {
  //
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
