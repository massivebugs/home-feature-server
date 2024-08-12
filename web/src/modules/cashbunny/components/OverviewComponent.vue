<template>
  <div ref="overviewContainer" class="overview-container" :style="dragStyle">
    <section :style="{ width: 100 - rightSectionWidth + '%' }">
      <div v-if="overviewData" class="revenue-expense">
        <div>
          <div class="overview-section-header">Revenue & Expense</div>
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
          <div class="overview-section-header">Profit</div>
          <div v-for="(sum, currency) in overviewData.sums" :key="currency">
            {{ currency }}
            {{ sum }}
          </div>
        </div>
      </div>
    </section>
    <section
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
          actions: {
            initCalendar: (c) => {
              console.log(c.selectedDates)
              console.log(c.selectedMonth)
              console.log(c.selectedYear)
            },
          },
        }"
      />
    </section>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { ResizeDirection, useDraggableResizable } from '@/core/composables/useDragResize'
import { RelativePosition } from '@/core/models/relative_position'
import { RelativeSize } from '@/core/models/relative_size'
import type { OverviewDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import CalendarComponent from './CalendarComponent.vue'

const store = useCashbunnyStore()
const overviewContainer = ref()
const overviewData = ref<OverviewDto | null>(null)
const selectedDate = ref<{ from: Date; to: Date } | null>(null)
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

onMounted(async () => {
  const res = await store.getOverview(selectedDate.value ?? undefined)
  if (res.data.error === null) {
    overviewData.value = res.data.data
  }
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.overview-container {
  position: relative;
  display: flex;
  width: 100%;
  min-height: 100%;

  > section {
    padding: 0.5em;
    &:nth-child(2) {
      border-left: 1px solid colors.$black;
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
