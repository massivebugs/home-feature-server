<template>
  <div ref="overviewContainer" class="overview-container" :style="dragStyle">
    <h2 v-if="dateStart && dateEnd">
      {{
        dateEnd.diff(dateStart, 'd') > 0
          ? t('cashbunny.overviewForDate', {
              v: `${dayjs(dateStart).format('YYYY-MM-DD')} ~ ${dayjs(dateEnd).format('YYYY-MM-DD')}`,
            })
          : t('cashbunny.overviewForDate', {
              v: dayjs(dateStart).format('YYYY-MM-DD'),
            })
      }}
    </h2>
    <section>
      <div :style="{ width: 100 - currentSize.w + '%' }">
        <h3 class="overview-section-header">
          {{ t('cashbunny.overviewSummary') }}
        </h3>
        <div v-if="overviewData">
          <DataTable
            :columns="[
              { data: 'currency', title: 'Currency' },
              { data: 'revenue', title: 'Revenue' },
              { data: 'expense', title: 'Expense' },
              { data: 'profit', title: 'Profit' },
            ]"
            :data="
              Object.entries(overviewData.summaries).map(([key, value]) => {
                return { currency: key, ...value }
              })
            "
            :options="{
              info: false,
              paging: false,
              searching: false,
              ordering: false,
            }"
            ref="summaryTable"
            class="table display nowrap compact"
          />
        </div>
      </div>
      <div
        ref="rightSection"
        :style="{ width: currentSize.w + '%' }"
        @mousedown.stop="onRightSectionResizeStart"
        @touchstart.stop="onRightSectionResizeStart"
      >
        <CalendarComponent
          class="cashbunny__overview__calendar"
          :tabs="[CalendarTabs.year, CalendarTabs.month]"
          :small="true"
          @loaded="onCalendarLoaded"
          @select-dates="onCalendarSelectDates"
        />
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import DataTablesCore from 'datatables.net'
import { DataTable } from 'datatables.net-vue3'
import dayjs, { Dayjs } from 'dayjs'
import { inject, onBeforeUnmount, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import type { ToggleWindowResizeHandlerFunc } from '@/core/components/WindowComponent.vue'
import { ResizeDirection, useDragResize } from '@/core/composables/useDragResize'
import { RelativePosition } from '@/core/models/relativePosition'
import { RelativeSize } from '@/core/models/relativeSize'
import type { OverviewDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import CalendarComponent, {
  type CalendarLoadedEvent,
  type CalendarSelectDatesEvent,
  CalendarTabs,
} from './CalendarComponent.vue'

DataTable.use(DataTablesCore)

const { t } = useI18n()
const store = useCashbunnyStore()
const overviewContainer = ref()
const rightSection = ref()
const overviewData = ref<OverviewDto | null>(null)
const dateStart = ref<Dayjs>()
const dateEnd = ref<Dayjs>()
let calendarResizeFunc: () => void
const addWindowResizeListener = inject('addWindowResizeListener') as ToggleWindowResizeHandlerFunc
const removeWindowResizeListener = inject(
  'removeWindowResizeListener',
) as ToggleWindowResizeHandlerFunc
const {
  currentSize,
  onResizeStart: onRightSectionResizeStart,
  dragStyle,
} = useDragResize(
  new RelativePosition(0, 0),
  new RelativeSize(30, 0),
  rightSection,
  overviewContainer,
  {
    resize: {
      direction: ResizeDirection.Left,
    },
  },
  () => {
    calendarResizeFunc()
  },
)

const onCalendarLoaded = (payload: CalendarLoadedEvent) => {
  calendarResizeFunc = payload.resizeFunc
  addWindowResizeListener(calendarResizeFunc)
}

const onCalendarSelectDates = async (payload: CalendarSelectDatesEvent) => {
  dateStart.value = dayjs(payload.dateStart)
  dateEnd.value = dayjs(payload.dateEnd)

  const res = await store.getOverview({
    from: dateStart.value.startOf('day'),
    to: dateEnd.value.endOf('day'),
  })
  if (res.data.error === null) {
    overviewData.value = res.data.data
  }
}

onBeforeUnmount(() => {
  removeWindowResizeListener(calendarResizeFunc)
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.overview-container {
  position: relative;
  display: flex;
  flex-direction: column;
  width: 100%;
  min-height: 100%;

  h2 {
    user-select: none;
    margin-top: 0;
  }

  > section {
    display: flex;
  }

  > section > div {
    overflow: hidden;

    &:nth-child(2) {
      margin-left: 1em;
      padding-left: 1em;
      border-left: 3px double colors.$black;
      min-width: 272px;
    }
  }
}

.overview-section-header {
  user-select: none;
  font-weight: 700;
  font-size: 1.1em;
  margin-bottom: 0.3em;
}

.cashbunny__overview__calendar {
  height: 400px;
}
</style>
