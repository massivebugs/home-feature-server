<template>
  <div class="cashbunny__schedule-container">
    <CalendarComponent
      class="cashbunny__schedule__calendar"
      :config="{
        settings: {
          visibility: {
            theme: 'light',
          },
        },
      }"
      @change-date="onCalendarChangeDate"
    />
    <div class="cashbunny__schedule__controls">
      <div></div>
      <div>Foo</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import dayjs, { Dayjs } from 'dayjs'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCashbunnyStore } from '../stores'
import CalendarComponent, { type CalendarChangeDateEvent } from './CalendarComponent.vue'

const { t } = useI18n()
const store = useCashbunnyStore()
const calendarData = ref<CalendarChangeDateEvent>()

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

  console.log(dateFrom, dateTo)
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.cashbunny__schedule__controls {
  display: flex;
  justify-content: space-between;
}

.cashbunny__schedule__calendar {
  border-radius: 0;
  width: 100%;
}
</style>
