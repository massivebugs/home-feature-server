<template>
  <form ref="form" class="cashbunny-planner-wizard-form" @submit.prevent="onSubmit">
    <div class="cashbunny-planner-wizard-form__header">
      <ButtonComponent
        :type="ButtonTypes.danger"
        v-if="showFormInputs"
        @click="onClickCancelCustom"
      >
        <CloseIconComponent stroke="white" />
      </ButtonComponent>
    </div>
    <div
      class="cashbunny-planner-wizard-form__preset-buttons"
      v-if="!showFormInputs && props.display === PlannerWizardAssetFormDisplays.preset"
    >
      <ButtonComponent
        v-for="option in presetOptions"
        :key="option.asset.name"
        :type="option.buttonType"
        @click.prevent="onClickPresetButton(option.asset)"
      >
        {{ option.asset.name }}
      </ButtonComponent>
      <ButtonComponent @click="onClickCustomButton"> Custom </ButtonComponent>
    </div>
    <div class="cashbunny-planner-wizard-form__inputs" v-else>
      <TextInputComponent
        name="assetName"
        :label="t('cashbunny.planner.asset.name')"
        :placeholder="t('cashbunny.planner.asset.namePlaceholder')"
        :required="true"
        :max="255"
        v-model="values.name"
      />
      <TextInputComponent
        name="assetDescription"
        :label="t('cashbunny.planner.asset.description')"
        :placeholder="t('cashbunny.planner.asset.descriptionPlaceholder')"
        :max="255"
        v-model="values.description"
      />
      <NumberInputComponent
        name="assetAddAmount"
        :label="t('cashbunny.planner.asset.amount')"
        placeholder="0"
        :min="0"
        :units="store.userPreferences?.user_currencies"
        v-model:value="values.amount"
        v-model:unit="values.currency"
      />
      <ButtonComponent type="success">
        {{ t('common.save') }}
      </ButtonComponent>
    </div>
  </form>
</template>

<script lang="ts">
export const PlannerWizardAssetFormDisplays = {
  preset: 'preset',
  form: 'form',
} as const
export type PlannerWizardAssetFormDisplay =
  (typeof PlannerWizardAssetFormDisplays)[keyof typeof PlannerWizardAssetFormDisplays]
</script>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import ButtonComponent, {
  ButtonTypes,
  type ButtonType,
} from '@/core/components/ButtonComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { PlannerAssetDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import { ref } from 'vue'
import CloseIconComponent from '@/core/components/CloseIconComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'

const props = defineProps<{
  display?: PlannerWizardAssetFormDisplay
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const values = defineModel<PlannerAssetDto>({ required: true })
const form = ref<HTMLFormElement>()
const showFormInputs = ref<boolean>(false)
const presetOptions: { buttonType?: ButtonType; asset: PlannerAssetDto }[] = [
  {
    buttonType: ButtonTypes.success,
    asset: {
      asset_account_id: '',
      name: t('cashbunny.planner.asset.presets.checkingAccount.name'),
      description: t('cashbunny.planner.asset.presets.checkingAccount.description'),
      amount: 0,
      currency: store.userPreferences?.user_currencies[0] ?? '',
    },
  },
  {
    buttonType: ButtonTypes.primary,
    asset: {
      asset_account_id: '',
      name: t('cashbunny.planner.asset.presets.savingsAccount.name'),
      description: t('cashbunny.planner.asset.presets.checkingAccount.description'),
      amount: 0,
      currency: store.userPreferences?.user_currencies[0] ?? '',
    },
  },
  {
    buttonType: ButtonTypes.warning,
    asset: {
      asset_account_id: '',
      name: t('cashbunny.planner.asset.presets.sinkingFunds.name'),
      description: t('cashbunny.planner.asset.presets.sinkingFunds.description'),
      amount: 0,
      currency: store.userPreferences?.user_currencies[0] ?? '',
    },
  },
]

const onClickPresetButton = (option: PlannerAssetDto) => {
  values.value.asset_account_id = option.asset_account_id
  values.value.name = option.name
  values.value.description = option.description
  values.value.currency = option.currency
  values.value.amount = option.amount
  showFormInputs.value = true
}

const onClickCustomButton = () => {
  showFormInputs.value = true
}

const onSubmit = (e: Event) => {
  if (!form.value?.reportValidity()) {
    e.stopImmediatePropagation()
  }
  showFormInputs.value = false
}

const onClickCancelCustom = () => {
  showFormInputs.value = false
}
</script>

<style scoped lang="scss">
.cashbunny-planner-wizard-form__header {
  display: flex;
  justify-content: flex-end;
}

.cashbunny-planner-wizard-form {
  width: 300px;
}

.cashbunny-planner-wizard-form__inputs {
  display: flex;
  flex-direction: column;
  gap: 1em;
}

.cashbunny-planner-wizard-form__preset-buttons {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5em;

  > * {
    width: calc(50% - 0.5em);
  }
}
</style>
