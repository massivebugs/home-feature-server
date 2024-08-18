<template>
  <div ref="calendar"></div>
</template>

<script setup lang="ts">
import VC from 'vanilla-calendar-pro'
import type { IOptions, IVanillaCalendar } from 'vanilla-calendar-pro/types'
import { onMounted, ref } from 'vue'

export type CalendarChangeDateEvent = {
  year: number
  month: number
  dates: string[]
}

const emit = defineEmits<{
  (e: 'changeDate', payload: CalendarChangeDateEvent): void
}>()

const props = defineProps<{
  config?: IOptions
}>()

const calendar = ref()
const vanillaCalendar = ref<VC | null>(null)

const onDate = (c: IVanillaCalendar) => {
  emit('changeDate', {
    year: c.selectedYear,
    month: c.selectedMonth,
    dates: c.selectedDates.filter((v) => v !== undefined),
  })
}

onMounted(() => {
  vanillaCalendar.value = new VC(calendar.value, {
    ...{
      actions: {
        clickDay: (_, c) => onDate(c),
        clickMonth: (_, c) => onDate(c),
        clickYear: (_, c) => onDate(c),
        clickArrow: (_, c) => onDate(c),
        initCalendar: onDate,
      },
    },
    ...props.config,
  })
  vanillaCalendar.value.init()
})
</script>

<style lang="scss">
@import 'vanilla-calendar-pro/build/vanilla-calendar.min.css';
</style>

<style scoped lang="scss">
:deep(.vanilla-calendar-day__btn) {
  border: 0;
  border-radius: 0 !important;
}

:deep(.vanilla-calendar-day__popup_custom) {
  //
}
</style>
