<template>
  <div ref="overviewContainer" class="overview-container" :style="dragStyle">
    <h2 v-if="calendarData">
      {{
        calendarData.dates.length === 1
          ? t('cashbunny.overviewForDate', {
              v: dayjs(calendarData.dates[0]).format('YYYY-MM-DD'),
            })
          : calendarData.dates.length > 1
            ? t('cashbunny.overviewForDate', {
                v: `${dayjs(calendarData.dates[0]).format('YYYY-MM-DD')} ~ ${dayjs(calendarData.dates[calendarData.dates.length - 1]).format('YYYY-MM-DD')}`,
              })
            : t('cashbunny.overviewForMonth', {
                m: calendarData.month + 1,
                y: calendarData.year,
              })
      }}
    </h2>
    <section>
      <div :style="{ width: 100 - rightSectionWidth + '%' }">
        <div v-if="overviewData" class="revenue-expense">
          <div>
            <div class="overview-section-header">
              {{ t('cashbunny.overviewRevenueAndExpense') }}
            </div>
            <div
              v-for="currency in new Set([
                ...Object.keys(overviewData.revenues),
                ...Object.keys(overviewData.expenses),
              ])"
              :key="currency"
            >
              {{ currency }}
              {{ overviewData.revenues[currency] }} / {{ overviewData.expenses[currency] }}
            </div>
          </div>
          <div>
            <div class="overview-section-header">{{ t('cashbunny.overviewProfit') }}</div>
            <div v-for="(sum, currency) in overviewData.sums" :key="currency">
              {{ currency }}
              {{ sum }}
            </div>
          </div>
        </div>
      </div>
      <div
        ref="rightSection"
        :style="{ width: rightSectionWidth + '%' }"
        @mousedown.stop="onRightSectionResizeStart"
        @touchstart.stop="onRightSectionResizeStart"
      >
        <CalendarComponent
          class="calendar"
          :config="{
            settings: {
              selection: {
                day: 'multiple-ranged',
              },
              visibility: {
                theme: 'light',
              },
            },
          }"
          @change-date="onCalendarChangeDate"
        />
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import dayjs, { Dayjs } from 'dayjs'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ResizeDirection, useDraggableResizable } from '@/core/composables/useDragResize'
import { RelativePosition } from '@/core/models/relative_position'
import { RelativeSize } from '@/core/models/relative_size'
import type { OverviewDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import CalendarComponent, { type CalendarChangeDateEvent } from './CalendarComponent.vue'

const { t } = useI18n()
const store = useCashbunnyStore()
const overviewContainer = ref()
const overviewData = ref<OverviewDto | null>(null)
const calendarData = ref<CalendarChangeDateEvent>()
const {
  boxWidth: rightSectionWidth,
  onResizeStart: onRightSectionResizeStart,
  dragStyle,
} = useDraggableResizable(
  new RelativePosition(0, 0),
  new RelativeSize(30, 0),
  undefined,
  overviewContainer,
  {
    resize: {
      direction: ResizeDirection.Left,
    },
  },
)

const onCalendarChangeDate = async (payload: CalendarChangeDateEvent) => {
  calendarData.value = payload
  let dateFrom: Dayjs
  let dateTo: Dayjs

  if (payload.dates.length) {
    dateFrom = dayjs(payload.dates[0])
    dateTo = dayjs(payload.dates[payload.dates.length > 1 ? payload.dates.length - 1 : 0]).add(
      1,
      'day',
    )
  } else {
    dateFrom = dayjs({ year: payload.year, month: payload.month, day: 1 })
    dateTo = dayjs({ year: payload.year, month: payload.month + 1 })
  }

  const res = await store.getOverview({ from: dateFrom, to: dateTo })
  if (res.data.error === null) {
    overviewData.value = res.data.data
  }
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.overview-container {
  position: relative;
  display: flex;
  flex-direction: column;
  width: 100%;
  min-height: 100%;
  padding: 0.5em;

  h2 {
    margin-top: 0;
  }

  > section {
    display: flex;
  }

  > section > div {
    &:nth-child(2) {
      margin-left: 0.5em;
      padding-left: 0.5em;
      border-left: 3px double colors.$black;
      min-width: 272px;
    }
  }
}

.overview-section-header {
  user-select: none;
  font-weight: 700;
  font-size: 1.1em;
}

.revenue-expense {
  display: flex;
  justify-content: space-between;
}

.calendar {
  padding: 0;
  border-radius: 0;
  width: 100%;
  background: none;
}
</style>
