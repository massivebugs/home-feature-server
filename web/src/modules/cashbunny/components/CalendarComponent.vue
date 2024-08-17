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
      CSSClasses: {
        dayBtn: 'vanilla-calendar-day__btn_custom',
      },
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
:deep(.vanilla-calendar-day__btn_custom) {
  border: 0;
  border-radius: 0 !important;
  flex: 1;
  min-height: 1.875rem;
  width: 100%;
  min-width: 1.875rem;
  cursor: pointer;
  padding: 0;
  font-size: 0.75rem;
  line-height: 1rem;
  font-weight: 400;
}

:deep(.vanilla-calendar-day__btn_custom:not([class$='_hover']):not([class$='_selected'])) {
  background-color: transparent;
}
</style>
