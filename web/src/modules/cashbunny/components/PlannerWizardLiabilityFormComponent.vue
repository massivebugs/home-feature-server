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
      v-if="!showFormInputs && props.display === PlannerWizardLiabilityFormDisplays.preset"
    >
      <ButtonComponent
        v-for="option in presetOptions"
        :key="option.liability.scheduled_transaction_id"
        :type="option.buttonType"
        @click.prevent="onClickPresetButton(option.liability)"
      >
        {{ option.liability.description }}
      </ButtonComponent>
      <ButtonComponent @click="onClickCustomButton"> Custom </ButtonComponent>
    </div>
    <div class="cashbunny-planner-wizard-form__inputs" v-else>
      <TextInputComponent
        name="liabilityDescription"
        :label="t('cashbunny.planner.liability.description')"
        :placeholder="t('cashbunny.planner.liability.descriptionPlaceholder')"
        :max="255"
        :required="true"
        v-model="values.description"
      />
      <NumberInputComponent
        name="liabilityAddAmount"
        :label="t('cashbunny.planner.liability.amount')"
        placeholder="1"
        :min="1"
        :units="store.userPreference?.userCurrencies"
        :required="true"
        v-model:value="values.amount"
        v-model:unit="values.currency"
      />
      <SelectInputComponent
        name="sourceAccount"
        :label="t('cashbunny.planner.liability.from')"
        :options="
          props.assets
            .filter((v) => v.currency === values.currency)
            .map(({ asset_account_id, name }) => ({
              label: name,
              value: asset_account_id,
            }))
        "
        :required="true"
        v-model="values.source_asset_account_id"
      />
      <SelectInputComponent
        name="destinationAccount"
        :label="t('cashbunny.planner.liability.to')"
        :options="
          props.liabilityAccounts
            .filter((v) => v.currency === values.currency)
            .map(({ id, name }) => ({ label: name, value: id.toString() }))
        "
        :required="true"
        v-model="values.destination_liability_account_id"
      />
      <RecurrenceRuleInputComponent v-model="values.recurrence_rule" />
      <ButtonComponent type="success">
        {{ t('common.save') }}
      </ButtonComponent>
    </div>
  </form>
</template>

<script lang="ts">
export const PlannerWizardLiabilityFormDisplays = {
  preset: 'preset',
  form: 'form',
} as const
export type PlannerWizardLiabilityFormDisplay =
  (typeof PlannerWizardLiabilityFormDisplays)[keyof typeof PlannerWizardLiabilityFormDisplays]
</script>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import ButtonComponent, {
  ButtonTypes,
  type ButtonType,
} from '@/core/components/ButtonComponent.vue'
import TextInputComponent from '@/core/components/TextInputComponent.vue'
import type { PlannerAssetDto, PlannerLiabilityDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import { computed, ref } from 'vue'
import CloseIconComponent from '@/core/components/CloseIconComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import dayjs from 'dayjs'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import RecurrenceRuleInputComponent from './RecurrenceRuleInputComponent.vue'
import { FrequencyStrs } from '../models/recurrence_rule'
import type { CashbunnyAccountResponse } from '@/core/composables/useAPI'

const props = defineProps<{
  display?: PlannerWizardLiabilityFormDisplay
  assets: PlannerAssetDto[]
  liabilityAccounts: CashbunnyAccountResponse[]
}>()

const { t } = useI18n()
const store = useCashbunnyStore()
const values = defineModel<PlannerLiabilityDto>({ required: true })
const form = ref<HTMLFormElement>()
const showFormInputs = ref<boolean>(false)
const presetOptions = computed<{ buttonType?: ButtonType; liability: PlannerLiabilityDto }[]>(
  () => [
    {
      buttonType: ButtonTypes.success,
      liability: {
        scheduled_transaction_id: '',
        description: t('cashbunny.planner.liability.presets.studentLoans.description'),
        amount: 0,
        currency: store.userPreference?.userCurrencies[0] ?? '',
        source_asset_account_id: '',
        source_asset_account_name: t('cashbunny.planner.liability.presets.studentLoans.from'),
        destination_liability_account_id: '',
        destination_liability_account_name: '',
        recurrence_rule: {
          freq: FrequencyStrs.MONTHLY,
          dtstart: dayjs().toString(),
          count: 0,
          interval: 1,
          until: dayjs().add(10, 'years').toString(),
        },
      },
    },
    {
      buttonType: ButtonTypes.primary,
      liability: {
        scheduled_transaction_id: '',
        description: t('cashbunny.planner.liability.presets.carLoans.description'),
        amount: 0,
        currency: store.userPreference?.userCurrencies[0] ?? '',
        source_asset_account_id: '',
        source_asset_account_name: t('cashbunny.planner.liability.presets.carLoans.from'),
        destination_liability_account_id: '',
        destination_liability_account_name: '',
        recurrence_rule: {
          freq: FrequencyStrs.MONTHLY,
          dtstart: dayjs().toString(),
          count: 0,
          interval: 1,
          until: dayjs().add(5, 'years').toString(),
        },
      },
    },
    {
      buttonType: ButtonTypes.warning,
      liability: {
        scheduled_transaction_id: '',
        description: t('cashbunny.planner.liability.presets.mortgage.description'),
        amount: 0,
        currency: store.userPreference?.userCurrencies[0] ?? '',
        source_asset_account_id: '',
        source_asset_account_name: t('cashbunny.planner.liability.presets.mortgage.from'),
        destination_liability_account_id: '',
        destination_liability_account_name: '',
        recurrence_rule: {
          freq: FrequencyStrs.MONTHLY,
          dtstart: dayjs().toString(),
          count: 0,
          interval: 1,
          until: dayjs().add(30, 'years').toString(),
        },
      },
    },
  ],
)

const onClickPresetButton = (option: PlannerLiabilityDto) => {
  values.value.scheduled_transaction_id = option.scheduled_transaction_id
  values.value.description = option.description
  values.value.currency = option.currency
  values.value.amount = option.amount
  values.value.source_asset_account_id = option.source_asset_account_id
  values.value.source_asset_account_name = option.source_asset_account_name
  values.value.destination_liability_account_id = option.destination_liability_account_id
  values.value.destination_liability_account_name = option.destination_liability_account_name
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
