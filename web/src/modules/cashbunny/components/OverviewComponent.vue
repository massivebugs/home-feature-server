<template>
  <div
    ref="overviewContainer"
    class="cashbunny-overview"
    :class="{ 'cashbunny-overview_column': isColumnLayout }"
    :style="dragStyle"
  >
    <CalendarComponent
      class="cashbunny-overview__calendar"
      height="100%"
      :style="{ width: isColumnLayout ? undefined : 100 - currentSize.w + '%' }"
      :tabs="[CalendarTabs.year, CalendarTabs.month, CalendarTabs.week, CalendarTabs.day]"
      more-link-action="day"
      :events="calendarEvents"
      @loaded="onCalendarLoaded"
      @select-dates="onCalendarSelectDates"
      @set-dates="onCalendarSetDates"
    />
    <div
      ref="detailSection"
      class="cashbunny-overview__report"
      :class="{
        'cashbunny-overview__report_resizable': !isColumnLayout,
      }"
      :style="{ width: isColumnLayout ? undefined : currentSize.w + '%' }"
      @mousedown.stop="onDetailSectionResizeStart"
      @touchstart.stop="onDetailSectionResizeStart"
    >
      <h2 class="cashbunny-overview__report__title" v-if="selectedDateStart && selectedDateEnd">
        {{
          selectedDateEnd.diff(selectedDateStart, 'd') > 0
            ? t('cashbunny.overviewForDate', {
                v: `${dayjs(selectedDateStart).format('YYYY-MM-DD')} ~ ${dayjs(selectedDateEnd).format('YYYY-MM-DD')}`,
              })
            : t('cashbunny.overviewForDate', {
                v: dayjs(selectedDateStart).format('YYYY-MM-DD'),
              })
        }}
      </h2>
      <h2 class="cashbunny-overview__report__title" v-else-if="viewDateStart && viewDateEnd">
        {{
          viewDateEnd.diff(viewDateStart, 'd') > 0
            ? t('cashbunny.overviewForDate', {
                v: `${dayjs(viewDateStart).format('YYYY-MM-DD')} ~ ${dayjs(viewDateEnd).format('YYYY-MM-DD')}`,
              })
            : t('cashbunny.overviewForDate', {
                v: dayjs(viewDateStart).format('YYYY-MM-DD'),
              })
        }}
      </h2>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.netWorth') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'currency', title: 'Currency' },
            { data: 'amount', title: 'Amount' },
          ]"
          :data="netWorthData"
          @loaded="onNetWorthDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.overviewProfitLossSummary') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'currency', title: 'Currency' },
            { data: 'revenue', title: 'Revenue' },
            { data: 'expense', title: 'Expense' },
            { data: 'profit', title: 'Profit' },
          ]"
          :data="summaryData"
          @loaded="onSummaryDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.assets.name') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'name', title: 'Name' },
            { data: 'amountDisplay', title: 'Amount' },
          ]"
          :data="selectedData ? selectedData.assetAccounts : viewData?.assetAccounts"
          @loaded="onAssetAccountDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.liabilities.name') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'name', title: 'Name' },
            { data: 'amountDisplay', title: 'Amount' },
          ]"
          :data="selectedData ? selectedData.liabilityAccounts : viewData?.liabilityAccounts"
          @loaded="onLiabilityAccountDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.transactions') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'description', title: 'Description' },
            { data: 'destinationAccountName', title: 'To' },
            { data: 'sourceAccountName', title: 'From' },
            { data: 'amount', title: 'Amount' },
          ]"
          :data="selectedData ? selectedData.transactions : viewData?.transactions"
          @loaded="onTransactionDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.scheduledTransactions') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'description', title: 'Description' },
            { data: 'destinationAccountName', title: 'To' },
            { data: 'sourceAccountName', title: 'From' },
            { data: 'amount', title: 'Amount' },
          ]"
          :data="
            selectedData
              ? selectedData.transactionsFromScheduled
              : viewData?.transactionsFromScheduled
          "
          @loaded="onScheduledTransactionDataTableLoaded"
        />
      </section>
    </div>
    <ErrorDialogComponent
      v-if="errorTitle && errorMessage"
      pos="center"
      :title="errorTitle"
      :message="errorMessage"
      @click-cancel="onErrorDialogClickClose"
    />
  </div>
</template>

<script setup lang="ts">
import { AxiosError } from 'axios'
import dayjs, { Dayjs } from 'dayjs'
import { computed, inject, onBeforeUnmount, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import DataTableComponent, {
  type DataTableLoadedEvent,
} from '@/core/components/DataTableComponent.vue'
import ErrorDialogComponent from '@/core/components/ErrorDialogComponent.vue'
import type {
  ToggleWindowResizeHandlerFunc,
  WindowSizeQuery,
} from '@/core/components/WindowComponent.vue'
import type { API, GetOverviewResponse } from '@/core/composables/useAPI'
import { ResizeDirection, useDragResize } from '@/core/composables/useDragResize'
import type { APIResponse } from '@/core/models/dto'
import { RelativePosition } from '@/core/models/relativePosition'
import { RelativeSize } from '@/core/models/relativeSize'
import CalendarComponent, {
  type CalendarEvent,
  type CalendarLoadedEvent,
  type CalendarSelectDatesEvent,
  type CalendarSetDatesEvent,
  CalendarTabs,
} from './CalendarComponent.vue'

const props = defineProps<{
  api: API
}>()

const { t } = useI18n()
const overviewContainer = ref()
const detailSection = ref()
const errorTitle = ref<string | null>(null)
const errorMessage = ref<string | null>(null)
const viewData = ref<GetOverviewResponse | null>(null)
const viewDateStart = ref<Dayjs>()
const viewDateEnd = ref<Dayjs>()
const selectedData = ref<GetOverviewResponse | null>(null)
const selectedDateStart = ref<Dayjs | null>()
const selectedDateEnd = ref<Dayjs | null>()
let netWorthDataTableResizeFunc: () => void
let summaryDataTableResizeFunc: () => void
let assetAccountDataTableResizeFunc: () => void
let liabilityAccountDataTableResizeFunc: () => void
let transactionDataTableResizeFunc: () => void
let scheduledTransactionDataTableResizeFunc: () => void
let calendarResizeFunc: () => void
const windowSizeQuery = inject<WindowSizeQuery>('windowSizeQuery')
const addWindowResizeListener = inject('addWindowResizeListener') as ToggleWindowResizeHandlerFunc
const removeWindowResizeListener = inject(
  'removeWindowResizeListener',
) as ToggleWindowResizeHandlerFunc
const {
  currentSize,
  onResizeStart: onDetailSectionResizeStart,
  dragStyle,
} = useDragResize(
  detailSection,
  new RelativePosition(0, 0),
  new RelativeSize(30, 0),
  overviewContainer,
  {
    resize: {
      direction: ResizeDirection.Left,
    },
  },
  () => {
    netWorthDataTableResizeFunc()
    summaryDataTableResizeFunc()
    assetAccountDataTableResizeFunc()
    liabilityAccountDataTableResizeFunc()
    transactionDataTableResizeFunc()
    scheduledTransactionDataTableResizeFunc()
    calendarResizeFunc()
  },
)
const isColumnLayout = computed(() => {
  return !windowSizeQuery?.md
})

const netWorthData = computed(() => {
  if (selectedData.value) {
    return Object.entries(selectedData.value.netWorth).map(([key, value]) => {
      return { currency: key, amount: value }
    })
  } else if (viewData.value) {
    return Object.entries(viewData.value.netWorth).map(([key, value]) => {
      return { currency: key, amount: value }
    })
  }
  return []
})

const summaryData = computed(() => {
  if (selectedData.value) {
    return Object.entries(selectedData.value.profitLossSummary).map(([key, value]) => {
      return { currency: key, ...value }
    })
  } else if (viewData.value) {
    return Object.entries(viewData.value.profitLossSummary).map(([key, value]) => {
      return { currency: key, ...value }
    })
  }
  return []
})

const calendarEvents = computed<CalendarEvent[]>(() => {
  return [
    ...(viewData?.value?.transactions.map(
      (transaction): CalendarEvent => ({
        title: transaction.description,
        start: transaction.transactedAt,
        allDay: false,
      }),
    ) ?? []),
    ...(viewData?.value?.transactionsFromScheduled.map(
      (scheduled): CalendarEvent => ({
        title: scheduled.description,
        start: scheduled.transactedAt,
        allDay: true,
      }),
    ) ?? []),
  ]
})

const onNetWorthDataTableLoaded = (payload: DataTableLoadedEvent) => {
  netWorthDataTableResizeFunc = payload.resizeFunc
}

const onSummaryDataTableLoaded = (payload: DataTableLoadedEvent) => {
  summaryDataTableResizeFunc = payload.resizeFunc
}

const onAssetAccountDataTableLoaded = (payload: DataTableLoadedEvent) => {
  assetAccountDataTableResizeFunc = payload.resizeFunc
}

const onLiabilityAccountDataTableLoaded = (payload: DataTableLoadedEvent) => {
  liabilityAccountDataTableResizeFunc = payload.resizeFunc
}

const onTransactionDataTableLoaded = (payload: DataTableLoadedEvent) => {
  transactionDataTableResizeFunc = payload.resizeFunc
}

const onScheduledTransactionDataTableLoaded = (payload: DataTableLoadedEvent) => {
  scheduledTransactionDataTableResizeFunc = payload.resizeFunc
}

const clearSelectedDateData = () => {
  selectedData.value = null
  selectedDateStart.value = null
  selectedDateEnd.value = null
}

const onCalendarLoaded = (payload: CalendarLoadedEvent) => {
  calendarResizeFunc = payload.resizeFunc
  addWindowResizeListener(calendarResizeFunc)

  calendarResizeFunc()
}

const onCalendarSetDates = async (payload: CalendarSetDatesEvent) => {
  clearSelectedDateData()
  viewDateStart.value = payload.dateStart
  viewDateEnd.value = payload.dateEnd

  try {
    const res = await props.api.getOverview({
      from: viewDateStart.value.startOf('day'),
      to: viewDateEnd.value.endOf('day'),
    })
    viewData.value = res
  } catch (error) {
    if (error instanceof AxiosError) {
      const res = error.response?.data as APIResponse<null>
      errorTitle.value = 'An error occured while loading data'
      errorMessage.value = res.error?.message || ''
    }
  }
}

const onCalendarSelectDates = async (payload: CalendarSelectDatesEvent | null) => {
  if (payload === null) {
    clearSelectedDateData()
    return
  }

  selectedDateStart.value = payload.dateStart
  selectedDateEnd.value = payload.dateEnd

  const res = await props.api.getOverview({
    from: selectedDateStart.value.startOf('day'),
    to: selectedDateEnd.value.endOf('day'),
  })
  selectedData.value = res
}

const onErrorDialogClickClose = () => {
  errorTitle.value = ''
  errorMessage.value = ''
}

onBeforeUnmount(() => {
  removeWindowResizeListener(calendarResizeFunc)
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.cashbunny-overview {
  position: relative;
  display: flex;
  height: 100%;
  padding: 0.5em;
  gap: 0.5em;
  background-color: colors.$low-opacity-light-grey;
}

.cashbunny-overview_column {
  height: auto;
  flex-direction: column;

  .cashbunny-overview__calendar {
    height: 500px !important;
  }
}

.cashbunny-overview__calendar {
  background-color: colors.$white;
}

.cashbunny-overview__report {
  overflow-y: auto;
  background-color: colors.$white;

  :deep(.datatable) {
    .dt-layout-row {
      margin: 0;
    }
  }
}

.cashbunny-overview__report_resizable {
  border-left: 3px double colors.$light-grey;
}

.cashbunny-overview__report__title {
  user-select: none;
  text-align: center;
  margin-top: 0.3em;
}

.cashbunny-overview__report__section {
  margin-bottom: 2em;
  background-color: colors.$white;
  border-radius: 5px;
  padding: 0.7em;
}

.cashbunny-overview__report__section-header {
  user-select: none;
  font-weight: 700;
  font-size: 1.1em;
  margin-top: 0;
  margin-bottom: 0.2em;
}
</style>
