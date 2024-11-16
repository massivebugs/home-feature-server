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
      v-if="!showFormInputs && props.display === PlannerWizardRevenueFormDisplays.preset"
    >
      <ButtonComponent
        v-for="option in presetOptions"
        :key="option.revenue.scheduled_transaction_id"
        :type="option.buttonType"
        @click.prevent="onClickPresetButton(option.revenue)"
      >
        {{ option.revenue.description }}
      </ButtonComponent>
      <ButtonComponent @click="onClickCustomButton"> Custom </ButtonComponent>
    </div>
    <div class="cashbunny-planner-wizard-form__inputs" v-else>
      <TextInputComponent
        name="revenueDescription"
        :label="t('cashbunny.planner.revenue.description')"
        :placeholder="t('cashbunny.planner.revenue.descriptionPlaceholder')"
        :max="255"
        :required="true"
        v-model="values.description"
      />
      <NumberInputComponent
        name="revenueAddAmount"
        :label="t('cashbunny.planner.revenue.amount')"
        placeholder="1"
        :min="1"
        :units="store.userPreference?.userCurrencies"
        :required="true"
        v-model:value="values.amount"
        v-model:unit="values.currency"
      />
      <SelectInputComponent
        name="sourceAccount"
        :label="t('cashbunny.planner.revenue.from')"
        :options="
          props.revenueAccounts
            .filter((v) => v.currency === values.currency)
            .map(({ id, name }) => ({ label: name, value: id.toString() }))
        "
        :required="true"
        v-model="values.source_revenue_account_id"
      />
      <SelectInputComponent
        name="destinationAccount"
        :label="t('cashbunny.planner.revenue.to')"
        :options="
          props.assets
            .filter((v) => v.currency === values.currency)
            .map(({ asset_account_id, name }) => ({
              label: name,
              value: asset_account_id,
            }))
        "
        :required="true"
        v-model="values.destination_asset_account_id"
      />
      <RecurrenceRuleInputComponent v-model="values.recurrence_rule" />
      <ButtonComponent type="success">
        {{ t('common.save') }}
      </ButtonComponent>
    </div>
  </form>
</template>

<script lang="ts">
export const PlannerWizardRevenueFormDisplays = {
  preset: 'preset',
  form: 'form',
} as const
export type PlannerWizardRevenueFormDisplay =
  (typeof PlannerWizardRevenueFormDisplays)[keyof typeof PlannerWizardRevenueFormDisplays]
</script>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import ButtonComponent, {
  ButtonTypes,
  type ButtonType,
} from '@/core/components/ButtonComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { PlannerAssetDto, PlannerRevenueDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import { computed, ref } from 'vue'
import CloseIconComponent from '@/core/components/CloseIconComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import dayjs from 'dayjs'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import RecurrenceRuleInputComponent from './RecurrenceRuleInputComponent.vue'
import { FrequencyStrs } from '../models/recurrence_rule'
import type { AccountResponse } from '@/core/composables/useAPI'

const props = defineProps<{
  display?: PlannerWizardRevenueFormDisplay
  assets: PlannerAssetDto[]
  revenueAccounts: AccountResponse[]
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const values = defineModel<PlannerRevenueDto>({ required: true })
const form = ref<HTMLFormElement>()
const showFormInputs = ref<boolean>(false)
const presetOptions = computed<{ buttonType?: ButtonType; revenue: PlannerRevenueDto }[]>(() => [
  {
    buttonType: ButtonTypes.success,
    revenue: {
      scheduled_transaction_id: '',
      description: t('cashbunny.planner.revenue.presets.salary.description'),
      amount: 0,
      currency: store.userPreference?.userCurrencies[0] ?? '',
      source_revenue_account_id: '',
      source_revenue_account_name: t('cashbunny.planner.revenue.presets.salary.from'),
      destination_asset_account_id: '',
      destination_asset_account_name: '',
      recurrence_rule: {
        freq: FrequencyStrs.WEEKLY,
        dtstart: dayjs().toString(),
        count: 0,
        interval: 2,
        until: dayjs().add(3, 'year').toString(),
      },
    },
  },
  {
    buttonType: ButtonTypes.primary,
    revenue: {
      scheduled_transaction_id: '',
      description: t('cashbunny.planner.revenue.presets.pension.description'),
      amount: 0,
      currency: store.userPreference?.userCurrencies[0] ?? '',
      source_revenue_account_id: '',
      source_revenue_account_name: t('cashbunny.planner.revenue.presets.pension.from'),
      destination_asset_account_id: '',
      destination_asset_account_name: '',
      recurrence_rule: {
        freq: FrequencyStrs.MONTHLY,
        dtstart: dayjs().toString(),
        count: 0,
        interval: 1,
        until: dayjs().add(30, 'year').toString(),
      },
    },
  },
  {
    buttonType: ButtonTypes.warning,
    revenue: {
      scheduled_transaction_id: '',
      description: t('cashbunny.planner.revenue.presets.dividends.description'),
      amount: 0,
      currency: store.userPreference?.userCurrencies[0] ?? '',
      source_revenue_account_id: '',
      source_revenue_account_name: t('cashbunny.planner.revenue.presets.dividends.from'),
      destination_asset_account_id: '',
      destination_asset_account_name: '',
      recurrence_rule: {
        freq: FrequencyStrs.MONTHLY,
        dtstart: dayjs().toString(),
        count: 0,
        interval: 3,
        until: dayjs().add(5, 'year').toString(),
      },
    },
  },
])

const onClickPresetButton = (option: PlannerRevenueDto) => {
  values.value.scheduled_transaction_id = option.scheduled_transaction_id
  values.value.description = option.description
  values.value.currency = option.currency
  values.value.amount = option.amount
  values.value.source_revenue_account_id = option.source_revenue_account_id
  values.value.source_revenue_account_name = option.source_revenue_account_name
  values.value.destination_asset_account_id = option.destination_asset_account_id
  values.value.destination_asset_account_name = option.destination_asset_account_name
  values.value.recurrence_rule = option.recurrence_rule
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
