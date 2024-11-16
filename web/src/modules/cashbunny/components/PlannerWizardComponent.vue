<template>
  <div
    class="cashbunny-planner-wizard"
    :class="{ 'cashbunny-planner-wizard__column': isColumnLayout }"
  >
    <SimplePageList
      class="cashbunny-planner-wizard__simple-pagelist"
      :pages="
        Object.values(PlannerWizardPages).map((v) => ({
          key: v,
          name: t(`cashbunny.planner.wizard.${v}.name`),
        }))
      "
      :type="isColumnLayout ? 'tabs' : 'list'"
      :current-key="currentPage"
      @click-page="onClickPageList"
    />
    <div class="cashbunny-planner-wizard__page">
      <section
        class="cashbunny-planner-wizard__welcome"
        v-if="currentPage === PlannerWizardPages.welcome"
      >
        <p class="cashbunny-planner-wizard__question">
          {{ t('cashbunny.planner.wizard.welcome.question') }}
        </p>
        <p class="cashbunny-planner-wizard__question-info">
          {{ t('cashbunny.planner.wizard.welcome.info') }}
        </p>
      </section>
      <section
        class="cashbunny-planner-wizard__assets"
        v-else-if="currentPage === PlannerWizardPages.assets"
      >
        <p class="cashbunny-planner-wizard__question">
          {{ t('cashbunny.planner.wizard.assets.question') }}
        </p>
        <p class="cashbunny-planner-wizard__question-info">
          {{ t('cashbunny.planner.wizard.assets.info') }}
        </p>
        <div :class="{ 'cashbunny-planner-wizard__section-content_column': isColumnLayout }">
          <PlannerWizardAssetDataTableComponent
            :data="plannerParameters.assets"
            @edit-row="onClickAssetEdit"
            @delete-row="onClickAssetDelete"
          />
          <PlannerWizardAssetFormComponent
            :display="assetFormDisplay"
            v-model="assetFormValues"
            @submit="onClickAssetFormSubmit"
          />
        </div>
      </section>
      <section
        class="cashbunny-planner-wizard__revenues"
        v-else-if="currentPage === PlannerWizardPages.revenues"
      >
        <p class="cashbunny-planner-wizard__question">
          {{ t('cashbunny.planner.wizard.revenues.question') }}
        </p>
        <p class="cashbunny-planner-wizard__question-info">
          {{ t('cashbunny.planner.wizard.revenues.info') }}
        </p>
        <div :class="{ 'cashbunny-planner-wizard__section-content_column': isColumnLayout }">
          <PlannerWizardRevenueDataTableComponent
            :data="plannerParameters.revenues"
            @edit-row="onClickRevenueEdit"
            @delete-row="onClickRevenueDelete"
          />
          <PlannerWizardRevenueFormComponent
            :display="revenueFormDisplay"
            :assets="plannerParameters.assets"
            :revenue-accounts="revenueAccounts"
            v-model="revenueFormValues"
            @submit="onClickRevenueFormSubmit"
          />
        </div>
      </section>
      <section
        class="cashbunny-planner-wizard__liabilities"
        v-else-if="currentPage === PlannerWizardPages.liabilities"
      >
        <p class="cashbunny-planner-wizard__question">
          {{ t('cashbunny.planner.wizard.liabilities.question') }}
        </p>
        <p class="cashbunny-planner-wizard__question-info">
          {{ t('cashbunny.planner.wizard.liabilities.info') }}
        </p>
        <div :class="{ 'cashbunny-planner-wizard__section-content_column': isColumnLayout }">
          <PlannerWizardLiabilityDataTableComponent
            :data="plannerParameters.liabilities"
            @edit-row="onClickLiabilityEdit"
            @delete-row="onClickLiabilityDelete"
          />
          <PlannerWizardLiabilityFormComponent
            :display="liabilityFormDisplay"
            :assets="plannerParameters.assets"
            :liability-accounts="liabilityAccounts"
            v-model="liabilityFormValues"
            @submit="onClickLiabilityFormSubmit"
          />
        </div>
      </section>
      <section
        class="cashbunny-planner-wizard__expenses"
        v-else-if="currentPage === PlannerWizardPages.expenses"
      >
        <p class="cashbunny-planner-wizard__question">
          {{ t('cashbunny.planner.wizard.expenses.question') }}
        </p>
        <p class="cashbunny-planner-wizard__question-info">
          {{ t('cashbunny.planner.wizard.expenses.info') }}
        </p>
        <div :class="{ 'cashbunny-planner-wizard__section-content_column': isColumnLayout }">
          <div>
            <div>
              <p>Monthly revenues</p>
              <div
                v-for="[currencyCode, totalAmount] in Object.entries(
                  plannerParameters.revenues.reduce(
                    (acc, item) => {
                      if (!acc[item.currency]) {
                        acc[item.currency] = 0
                      }
                      acc[item.currency] += item.amount
                      return acc
                    },
                    {} as Record<string, number>,
                  ),
                )"
                :key="currencyCode"
              >
                {{ currencyCode }} {{ totalAmount.toLocaleString()
                }}{{ store.currencies[currencyCode] }}
              </div>
            </div>
            <div>Liability</div>
            <div>Expenses</div>
          </div>
          <div class="cashbunny-planner-wizard__categories">
            <p>Categories</p>
            <div
              class="cashbunny-planner-wizard__categories-row"
              v-for="{ id, name } in plannerParameters.transaction_categories"
              :key="id"
            >
              {{ name }} <RangeInputComponent :min="0" :max="100" :step="1" />
            </div>
          </div>
        </div>
      </section>
      <section
        class="cashbunny-planner-wizard__complete"
        v-else-if="currentPage === PlannerWizardPages.complete"
      >
        <p class="cashbunny-planner-wizard__question">
          {{ t('cashbunny.planner.wizard.complete.question') }}
        </p>
        <p class="cashbunny-planner-wizard__question-info">
          {{ t('cashbunny.planner.wizard.complete.info') }}
        </p>
      </section>
      <div class="cashbunny-planner-wizard__nav-buttons">
        <ButtonComponent @click="onClickBack">
          <template v-if="currentPage === PlannerWizardPages.welcome">
            {{ t('cashbunny.planner.wizard.welcome.no') }}
          </template>
          <template v-else>
            {{ t('common.back') }}
          </template>
        </ButtonComponent>
        <ButtonComponent
          type="success"
          @click="onClickContinue"
          :disabled="isContinueButtonDisabled"
        >
          <template v-if="currentPage === PlannerWizardPages.welcome">
            {{ t('cashbunny.planner.wizard.welcome.yes') }}
          </template>
          <template v-else>
            {{ t('common.continue') }}
          </template>
        </ButtonComponent>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { uniqueId } from 'lodash'
import { computed, inject, onMounted, ref, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import ButtonComponent from '@/core/components/ButtonComponent.vue'
import RangeInputComponent from '@/core/components/RangeInputComponent.vue'
import SimplePageList, {
  type SimplePageListClickPageEvent,
} from '@/core/components/SimplePageList.vue'
import type { WindowSizeQuery } from '@/core/components/WindowComponent.vue'
import type { AccountResponse } from '@/core/composables/useAPI'
import {
  type PlannerAssetDto,
  type PlannerLiabilityDto,
  type PlannerParametersDto,
  type PlannerRevenueDto,
} from '../models/dto'
import { useCashbunnyStore } from '../stores'
import PlannerWizardAssetDataTableComponent from './PlannerWizardAssetDataTableComponent.vue'
import PlannerWizardAssetFormComponent, {
  type PlannerWizardAssetFormDisplay,
  PlannerWizardAssetFormDisplays,
} from './PlannerWizardAssetFormComponent.vue'
import PlannerWizardLiabilityDataTableComponent from './PlannerWizardLiabilityDataTableComponent.vue'
import PlannerWizardLiabilityFormComponent, {
  type PlannerWizardLiabilityFormDisplay,
  PlannerWizardLiabilityFormDisplays,
} from './PlannerWizardLiabilityFormComponent.vue'
import PlannerWizardRevenueDataTableComponent from './PlannerWizardRevenueDataTableComponent.vue'
import PlannerWizardRevenueFormComponent, {
  type PlannerWizardRevenueFormDisplay,
  PlannerWizardRevenueFormDisplays,
} from './PlannerWizardRevenueFormComponent.vue'

const PlannerWizardPages = {
  welcome: 'welcome',
  assets: 'assets',
  revenues: 'revenues',
  liabilities: 'liabilities',
  expenses: 'expenses',
  complete: 'complete',
} as const
type PlannerWizardPage = (typeof PlannerWizardPages)[keyof typeof PlannerWizardPages]

const { t } = useI18n()
const store = useCashbunnyStore()
const windowSizeQuery = inject<WindowSizeQuery>('windowSizeQuery')
const currentPage = ref<PlannerWizardPage>(PlannerWizardPages.welcome)
const selectedAsset = ref<PlannerAssetDto | null>(null)
const selectedRevenue = ref<PlannerRevenueDto | null>(null)
const selectedLiability = ref<PlannerLiabilityDto | null>(null)
const assetFormDisplay = ref<PlannerWizardAssetFormDisplay>(PlannerWizardAssetFormDisplays.preset)
const revenueFormDisplay = ref<PlannerWizardRevenueFormDisplay>(
  PlannerWizardRevenueFormDisplays.preset,
)
const liabilityFormDisplay = ref<PlannerWizardLiabilityFormDisplay>(
  PlannerWizardLiabilityFormDisplays.preset,
)
const assetFormValues = ref<PlannerAssetDto>({
  asset_account_id: '',
  name: '',
  description: '',
  currency: '',
  amount: 0,
})
const revenueFormValues = ref<PlannerRevenueDto>({
  scheduled_transaction_id: '',
  description: '',
  currency: store.userPreference?.userCurrencies[0] ?? '',
  amount: 0,
  source_revenue_account_id: '',
  source_revenue_account_name: '',
  destination_asset_account_id: '',
  destination_asset_account_name: '',
  recurrence_rule: {
    freq: 'MONTHLY',
    dtstart: '',
    count: 0,
    interval: 1,
    until: '',
  },
  transaction_category: null,
})
const liabilityFormValues = ref<PlannerLiabilityDto>({
  scheduled_transaction_id: '',
  description: '',
  currency: store.userPreference?.userCurrencies[0] ?? '',
  amount: 0,
  source_asset_account_id: '',
  source_asset_account_name: '',
  destination_liability_account_id: '',
  destination_liability_account_name: '',
  recurrence_rule: {
    freq: 'MONTHLY',
    dtstart: '',
    count: 0,
    interval: 1,
    until: '',
  },
  transaction_category: null,
})
const plannerParameters = ref<PlannerParametersDto>({
  assets: [],
  revenues: [],
  liabilities: [],
  transaction_categories: [],
})
const allAccounts = ref<AccountResponse[]>([])
const isColumnLayout = computed(() => {
  return !windowSizeQuery?.md
})
const isContinueButtonDisabled = computed(() => {
  switch (currentPage.value) {
    case PlannerWizardPages.assets:
      // The user must have at least one asset account
      return plannerParameters.value.assets.length === 0
  }
  return false
})
const revenueAccounts = computed(() => allAccounts.value.filter((v) => v.category === 'revenues'))
const liabilityAccounts = computed(() =>
  allAccounts.value.filter((v) => v.category === 'liabilities'),
)

watch(currentPage, () => {
  resetAssetFormValues()
  resetRevenueFormValues()
  resetLiabilityFormValues()
})

const resetAssetFormValues = () => {
  assetFormValues.value = {
    asset_account_id: '',
    name: '',
    description: '',
    currency: store.userPreference?.userCurrencies[0] ?? '',
    amount: 0,
  }

  selectedAsset.value = null
  assetFormDisplay.value = PlannerWizardAssetFormDisplays.preset
}

const resetRevenueFormValues = () => {
  revenueFormValues.value = {
    scheduled_transaction_id: '',
    description: '',
    currency: store.userPreference?.userCurrencies[0] ?? '',
    amount: 0,
    source_revenue_account_id: '',
    source_revenue_account_name: '',
    destination_asset_account_id: '',
    destination_asset_account_name: '',
    recurrence_rule: {
      freq: 'MONTHLY',
      dtstart: dayjs().toString(),
      count: 0,
      interval: 1,
      until: dayjs().add(3, 'year').toString(),
    },
    transaction_category: null,
  }

  selectedRevenue.value = null
  revenueFormDisplay.value = PlannerWizardRevenueFormDisplays.preset
}

const resetLiabilityFormValues = () => {
  liabilityFormValues.value = {
    scheduled_transaction_id: '',
    description: '',
    currency: store.userPreference?.userCurrencies[0] ?? '',
    amount: 0,
    source_asset_account_id: '',
    source_asset_account_name: '',
    destination_liability_account_id: '',
    destination_liability_account_name: '',
    recurrence_rule: {
      freq: 'MONTHLY',
      dtstart: dayjs().toString(),
      count: 0,
      interval: 1,
      until: dayjs().add(3, 'year').toString(),
    },
    transaction_category: null,
  }

  selectedLiability.value = null
  liabilityFormDisplay.value = PlannerWizardLiabilityFormDisplays.preset
}

const onClickPageList = (data: SimplePageListClickPageEvent<PlannerWizardPage>) => {
  currentPage.value = data.key
}

const onClickAssetEdit = (data: PlannerAssetDto) => {
  assetFormValues.value = {
    asset_account_id: data.asset_account_id,
    name: data.name,
    description: data.description,
    currency: data.currency,
    amount: data.amount,
  }

  selectedAsset.value = data
  assetFormDisplay.value = PlannerWizardAssetFormDisplays.form
}

const onClickAssetDelete = (data: PlannerAssetDto) => {
  plannerParameters.value.assets.splice(plannerParameters.value.assets.indexOf(data), 1)
}

const onClickAssetFormSubmit = () => {
  if (selectedAsset.value) {
    selectedAsset.value.asset_account_id = assetFormValues.value.asset_account_id
    selectedAsset.value.name = assetFormValues.value.name
    selectedAsset.value.description = assetFormValues.value.description
    selectedAsset.value.currency = assetFormValues.value.currency
    selectedAsset.value.amount = assetFormValues.value.amount
  } else if (assetFormValues.value.asset_account_id === '') {
    assetFormValues.value.asset_account_id = uniqueId('new-')
    plannerParameters.value.assets.push(assetFormValues.value)
  }

  resetAssetFormValues()
}

const onClickRevenueEdit = (data: PlannerRevenueDto) => {
  revenueFormValues.value = {
    scheduled_transaction_id: data.scheduled_transaction_id,
    description: data.description,
    amount: data.amount,
    currency: data.currency,
    source_revenue_account_id: data.source_revenue_account_id,
    source_revenue_account_name: data.source_revenue_account_name,
    destination_asset_account_id: data.source_revenue_account_id,
    destination_asset_account_name: data.destination_asset_account_name,
    recurrence_rule: data.recurrence_rule,
    transaction_category: null,
  }

  selectedRevenue.value = data
  revenueFormDisplay.value = PlannerWizardRevenueFormDisplays.form
}

const onClickRevenueDelete = (data: PlannerRevenueDto) => {
  plannerParameters.value.revenues.splice(plannerParameters.value.revenues.indexOf(data), 1)
}

const onClickRevenueFormSubmit = () => {
  if (selectedRevenue.value) {
    selectedRevenue.value.scheduled_transaction_id =
      revenueFormValues.value.scheduled_transaction_id
    selectedRevenue.value.description = revenueFormValues.value.description
    selectedRevenue.value.amount = revenueFormValues.value.amount
    selectedRevenue.value.currency = revenueFormValues.value.currency
    selectedRevenue.value.source_revenue_account_id =
      revenueFormValues.value.source_revenue_account_id
    selectedRevenue.value.source_revenue_account_name =
      revenueFormValues.value.source_revenue_account_name
    selectedRevenue.value.destination_asset_account_id =
      revenueFormValues.value.destination_asset_account_id
    selectedRevenue.value.destination_asset_account_name =
      revenueFormValues.value.destination_asset_account_name
    selectedRevenue.value.recurrence_rule = revenueFormValues.value.recurrence_rule
  } else if (revenueFormValues.value.scheduled_transaction_id === '') {
    revenueFormValues.value.scheduled_transaction_id = uniqueId('new-')
    plannerParameters.value.revenues.push(revenueFormValues.value)
  }

  resetRevenueFormValues()
}

const onClickLiabilityEdit = (data: PlannerLiabilityDto) => {
  liabilityFormValues.value = {
    scheduled_transaction_id: data.scheduled_transaction_id,
    description: data.description,
    amount: data.amount,
    currency: data.currency,
    source_asset_account_id: data.source_asset_account_id,
    source_asset_account_name: data.source_asset_account_name,
    destination_liability_account_id: data.destination_liability_account_id,
    destination_liability_account_name: data.destination_liability_account_name,
    recurrence_rule: data.recurrence_rule,
    transaction_category: null,
  }

  selectedLiability.value = data
  liabilityFormDisplay.value = PlannerWizardLiabilityFormDisplays.form
}

const onClickLiabilityDelete = (data: PlannerLiabilityDto) => {
  plannerParameters.value.liabilities.splice(plannerParameters.value.liabilities.indexOf(data), 1)
}

const onClickLiabilityFormSubmit = () => {
  if (selectedLiability.value) {
    selectedLiability.value.scheduled_transaction_id =
      liabilityFormValues.value.scheduled_transaction_id
    selectedLiability.value.description = liabilityFormValues.value.description
    selectedLiability.value.amount = liabilityFormValues.value.amount
    selectedLiability.value.currency = liabilityFormValues.value.currency
    selectedLiability.value.source_asset_account_id =
      liabilityFormValues.value.source_asset_account_id
    selectedLiability.value.source_asset_account_name =
      liabilityFormValues.value.source_asset_account_name
    selectedLiability.value.destination_liability_account_id =
      liabilityFormValues.value.destination_liability_account_id
    selectedLiability.value.destination_liability_account_name =
      liabilityFormValues.value.destination_liability_account_name
    selectedLiability.value.recurrence_rule = liabilityFormValues.value.recurrence_rule
  } else if (liabilityFormValues.value.scheduled_transaction_id === '') {
    liabilityFormValues.value.scheduled_transaction_id = uniqueId('new-')
    plannerParameters.value.liabilities.push(liabilityFormValues.value)
  }

  resetLiabilityFormValues()
}

const onClickBack = () => {
  const pagesInOrder = [...Object.values(PlannerWizardPages)]
  const currentPageIdx = pagesInOrder.indexOf(currentPage.value)

  if (currentPageIdx !== 0) {
    currentPage.value = pagesInOrder[currentPageIdx - 1]
  }
}

const onClickContinue = () => {
  if (currentPage.value !== PlannerWizardPages.complete) {
    const pagesInOrder = [...Object.values(PlannerWizardPages)]
    const currentPageIdx = pagesInOrder.indexOf(currentPage.value)
    currentPage.value = pagesInOrder[currentPageIdx + 1]
  }
}

onMounted(async () => {
  const plannerParamRes = await store.getPlannerParameters()
  if (plannerParamRes.data.error === null) {
    plannerParameters.value = plannerParamRes.data.data
  }

  const accountsRes = await store.getAccounts()
  if (accountsRes.data.error === null) {
    allAccounts.value = accountsRes.data.data
  }
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.cashbunny-planner-wizard {
  height: 100%;
  display: flex;
  align-items: stretch;

  &.cashbunny-planner-wizard__column {
    flex-direction: column;
  }

  &:not(.cashbunny-planner-wizard__column) .cashbunny-planner-wizard__simple-pagelist {
    width: 200px;
  }

  .cashbunny-planner-wizard__simple-pagelist {
    margin: 1em;
  }

  .cashbunny-planner-wizard__page {
    position: relative;
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    overflow: auto;
    padding: 1em 1em 0 1em;
  }

  section {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
  }

  .cashbunny-planner-wizard__question {
    white-space: pre-line;
    text-align: center;
    line-height: 1.7em;
    font-size: 1.2em;
    user-select: none;
    margin-bottom: 0.7em;
    font-weight: 700;
  }

  .cashbunny-planner-wizard__question-info {
    white-space: pre-line;
    text-align: center;
    line-height: 1.5em;
    font-size: 1.1em;
    margin-top: 0;
    margin-bottom: 2em;
    user-select: none;
    color: colors.$dark-grey;
  }

  .cashbunny-planner-wizard__nav-buttons {
    position: sticky;
    bottom: 0;
    padding: 1em 0 3em 0;
    background-color: colors.$white;
    display: flex;
    gap: 0.5em;
    width: 100%;
    > button {
      flex: 1;
    }
  }

  .cashbunny-planner-wizard__section-content_column {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 1em;

    > * {
      width: 100%;
    }
  }

  section.cashbunny-planner-wizard__welcome {
    justify-content: center;
  }

  section.cashbunny-planner-wizard__assets,
  section.cashbunny-planner-wizard__revenues,
  section.cashbunny-planner-wizard__liabilities,
  section.cashbunny-planner-wizard__expenses {
    > div {
      width: 100%;
      display: flex;
      gap: 1em;

      > * {
        flex: 1;
      }
    }
  }
}
</style>
