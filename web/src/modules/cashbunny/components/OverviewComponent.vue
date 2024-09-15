<template>
  <div ref="overviewContainer" class="cashbunny-overview" :style="dragStyle">
    <CalendarComponent
      class="cashbunny-overview__calendar"
      height="100%"
      :style="{ width: 100 - currentSize.w + '%' }"
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
      :style="{ width: currentSize.w + '%' }"
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
          :options="{
            info: false,
            paging: false,
            searching: false,
            ordering: false,
            select: false,
          }"
          class="table display nowrap compact"
          @loaded="onNetWorthDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.overviewSummary') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'currency', title: 'Currency' },
            { data: 'revenue', title: 'Revenue' },
            { data: 'expense', title: 'Expense' },
            { data: 'profit', title: 'Profit' },
          ]"
          :data="summaryData"
          :options="{
            info: false,
            paging: false,
            searching: false,
            ordering: false,
            select: false,
          }"
          class="table display nowrap compact"
          @loaded="onSummaryDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.transactions') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'to', title: 'To' },
            { data: 'from', title: 'From' },
            { data: 'amount', title: 'Amount' },
          ]"
          :data="transactionData"
          :options="{
            info: false,
            paging: false,
            searching: false,
            ordering: false,
            select: false,
          }"
          class="table display nowrap compact"
          @loaded="onTransactionDataTableLoaded"
        />
      </section>
      <section class="cashbunny-overview__report__section">
        <h3 class="cashbunny-overview__report__section-header">
          {{ t('cashbunny.scheduledTransactions') }}
        </h3>
        <DataTableComponent
          :columns="[
            { data: 'to', title: 'To' },
            { data: 'from', title: 'From' },
            { data: 'amount', title: 'Amount' },
          ]"
          :data="scheduledTransactionData"
          :options="{
            info: false,
            paging: false,
            searching: false,
            ordering: false,
            select: false,
          }"
          class="table display nowrap compact"
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
import type { ToggleWindowResizeHandlerFunc } from '@/core/components/WindowComponent.vue'
import { ResizeDirection, useDragResize } from '@/core/composables/useDragResize'
import type { APIResponse } from '@/core/models/dto'
import { RelativePosition } from '@/core/models/relativePosition'
import { RelativeSize } from '@/core/models/relativeSize'
import type { OverviewDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import CalendarComponent, {
  type CalendarEvent,
  type CalendarLoadedEvent,
  type CalendarSelectDatesEvent,
  type CalendarSetDatesEvent,
  CalendarTabs,
} from './CalendarComponent.vue'

const { t } = useI18n()
const store = useCashbunnyStore()
const overviewContainer = ref()
const detailSection = ref()
const errorTitle = ref<string | null>(null)
const errorMessage = ref<string | null>(null)
const viewData = ref<OverviewDto | null>(null)
const viewDateStart = ref<Dayjs>()
const viewDateEnd = ref<Dayjs>()
const selectedData = ref<OverviewDto | null>(null)
const selectedDateStart = ref<Dayjs | null>()
const selectedDateEnd = ref<Dayjs | null>()
let netWorthDataTableResizeFunc: () => void
let summaryDataTableResizeFunc: () => void
let transactionDataTableResizeFunc: () => void
let scheduledTransactionDataTableResizeFunc: () => void
let calendarResizeFunc: () => void
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
    transactionDataTableResizeFunc()
    scheduledTransactionDataTableResizeFunc()
    calendarResizeFunc()
  },
)

const netWorthData = computed(() => {
  if (selectedData.value) {
    return Object.entries(selectedData.value.net_worth).map(([key, value]) => {
      return { currency: key, amount: value }
    })
  } else if (viewData.value) {
    return Object.entries(viewData.value.net_worth).map(([key, value]) => {
      return { currency: key, amount: value }
    })
  }
  return []
})

const summaryData = computed(() => {
  if (selectedData.value) {
    return Object.entries(selectedData.value.summaries).map(([key, value]) => {
      return { currency: key, ...value }
    })
  } else if (viewData.value) {
    return Object.entries(viewData.value.summaries).map(([key, value]) => {
      return { currency: key, ...value }
    })
  }
  return []
})

const transactionData = computed(() => {
  if (selectedData.value) {
    return Object.values(selectedData.value.transactions).map((value) => {
      return {
        to: value.destination_account_name,
        from: value.source_account_name,
        amount: value.amount_display,
      }
    })
  } else if (viewData.value) {
    return Object.values(viewData.value.transactions).map((value) => {
      return {
        to: value.destination_account_name,
        from: value.source_account_name,
        amount: value.amount_display,
      }
    })
  }
  return []
})

const scheduledTransactionData = computed(() => {
  if (selectedData.value) {
    return Object.values(selectedData.value.transactions_from_scheduled).map((value) => {
      return {
        to: value.destination_account_name,
        from: value.source_account_name,
        amount: value.amount_display,
      }
    })
  } else if (viewData.value) {
    return Object.values(viewData.value.transactions_from_scheduled).map((value) => {
      return {
        to: value.destination_account_name,
        from: value.source_account_name,
        amount: value.amount_display,
      }
    })
  }
  return []
})

const calendarEvents = computed<CalendarEvent[]>(() => {
  return [
    ...(viewData?.value?.transactions?.map(
      (transaction): CalendarEvent => ({
        title: transaction.description,
        start: transaction.transacted_at,
        allDay: false,
      }),
    ) ?? []),
    ...(viewData?.value?.transactions_from_scheduled?.map(
      (scheduled): CalendarEvent => ({
        title: scheduled.description,
        start: scheduled.transacted_at,
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
    const res = await store.getOverview({
      from: viewDateStart.value.startOf('day'),
      to: viewDateEnd.value.endOf('day'),
    })
    viewData.value = res.data.data
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

  const res = await store.getOverview({
    from: selectedDateStart.value.startOf('day'),
    to: selectedDateEnd.value.endOf('day'),
  })
  if (res.data.error === null) {
    selectedData.value = res.data.data
  }
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
  height: 100%;
  display: flex;
}

.cashbunny-overview__calendar {
  padding: 1em;
  margin-right: 5px;
  background-color: colors.$white;
}

.cashbunny-overview__report {
  padding: 1em;
  overflow: hidden;
  background-color: colors.$white;

  :deep(.datatable) {
    .dt-layout-row {
      margin: 0;
    }
  }
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
