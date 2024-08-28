<template>
  <div class="cashbunny__schedule-container">
    <CalendarComponent
      class="cashbunny__schedule-calendar"
      :tabs="[CalendarTabs.year, CalendarTabs.month, CalendarTabs.week, CalendarTabs.day]"
      @loaded="onCalendarLoaded"
    />
  </div>
</template>

<script setup lang="ts">
import { inject, onBeforeUnmount } from 'vue'
import type { ToggleWindowResizeHandlerFunc } from '@/core/components/WindowComponent.vue'
import CalendarComponent, { type CalendarLoadedEvent, CalendarTabs } from './CalendarComponent.vue'

let calendarResizeFunc: () => void
const addWindowResizeListener = inject('addWindowResizeListener') as ToggleWindowResizeHandlerFunc
const removeWindowResizeListener = inject(
  'removeWindowResizeListener',
) as ToggleWindowResizeHandlerFunc

const onCalendarLoaded = (payload: CalendarLoadedEvent) => {
  calendarResizeFunc = payload.resizeFunc
  addWindowResizeListener(calendarResizeFunc)
}

onBeforeUnmount(() => {
  removeWindowResizeListener(calendarResizeFunc)
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.cashbunny__schedule-calendar {
  height: 500px;
}
</style>
