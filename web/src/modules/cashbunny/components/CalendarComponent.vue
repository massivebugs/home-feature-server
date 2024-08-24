<template>
  <FullCalendarComponent ref="fullCalendar" :options="calendarOptions" />
</template>

<script lang="ts">
export const CalendarTabs = {
  year: 'dayGridYear',
  month: 'dayGridMonth',
  week: 'timeGridWeek',
  day: 'timeGridDay',
} as const
export type CalendarTab = (typeof CalendarTabs)[keyof typeof CalendarTabs]
</script>

<script setup lang="ts">
import {
  type CalendarApi,
  type CalendarOptions,
  type DateSelectArg,
  type DatesSetArg,
} from '@fullcalendar/core/index.js'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from '@fullcalendar/interaction'
import timeGridPlugin from '@fullcalendar/timegrid'
import FullCalendarComponent from '@fullcalendar/vue3'
import type { Dayjs } from 'dayjs'
import dayjs from 'dayjs'
import { debounce } from 'lodash'
import { onMounted, ref } from 'vue'

export type CalendarSelectDatesEvent = {
  dateStart: Dayjs
  dateEnd: Dayjs
}

export type CalendarLoadedEvent = {
  resizeFunc: () => void
}

const emit = defineEmits<{
  (e: 'selectDates', payload: CalendarSelectDatesEvent): void
  (e: 'loaded', payload: CalendarLoadedEvent): void
}>()

const props = defineProps<{
  tabs: CalendarTab[]
  height?: string
}>()

const fullCalendar = ref()
let fullCalendarApi: CalendarApi
const resizeFunc = debounce(() => fullCalendarApi.updateSize(), 100)
let currentSelection: CalendarSelectDatesEvent | null = null

const onSelect = (arg: DatesSetArg | DateSelectArg) => {
  // We want dateEnd to not include the next day (we didn't select it!)
  const dateStart = dayjs(arg.start)
  const dateEnd = dayjs(arg.end).subtract(1, 's')

  // Clear date selection if the same date has been picked
  if (
    currentSelection &&
    currentSelection.dateStart.isSame(dateStart) &&
    currentSelection.dateEnd.isSame(dateEnd)
  ) {
    fullCalendarApi.unselect()
    currentSelection = null
    return
  }

  currentSelection = { dateStart, dateEnd }
  emit('selectDates', currentSelection)
}

const calendarOptions: CalendarOptions = {
  height: props.height,
  plugins: [dayGridPlugin, timeGridPlugin, interactionPlugin],
  initialView: 'dayGridMonth',
  headerToolbar: {
    left: 'prev,next,today',
    center: 'title',
    right: props.tabs.join(','),
  },
  allDaySlot: false,
  editable: true,
  selectable: true,
  selectMirror: true,
  dayMaxEvents: true,
  showNonCurrentDates: false,
  datesSet: onSelect,
  select: onSelect,
  eventClick: (arg) => {
    console.log('eventClick', arg)
  },
  eventAdd: (arg) => {
    console.log('eventAdd', arg)
  },
  eventChange: (arg) => {
    console.log('eventChange', arg)
  },
  eventRemove: (arg) => {
    console.log('eventRemove', arg)
  },
  events: [
    { title: 'event 1', date: '2019-04-01' },
    { title: 'event 2', date: '2019-04-02' },
  ],
}

onMounted(() => {
  fullCalendarApi = fullCalendar.value.getApi()
  emit('loaded', { resizeFunc })
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';
:deep() {
  user-select: none;
}

:deep(.fc-header-toolbar) {
  flex-direction: column;
  gap: 5px;

  .fc-button {
    padding: 0 0.4em;
    font-size: 0.9em;
    background-color: colors.$dark-grey;
    text-transform: capitalize;

    &.fc-button-active {
      background-color: colors.$black;
    }
  }

  .fc-toolbar-title {
    font-size: 1em;
  }
}

:deep(.fc-daygrid-day) {
  .fc-highlight {
    background-color: colors.$high-opacity-viridian;
  }

  &.fc-day-today {
    background-color: colors.$high-opacity-peach;
  }

  .fc-daygrid-month-start {
    font-size: 0.8em;
    text-align: center;
    width: 100%;
    background-color: colors.$dark-grey;
    color: colors.$white;
  }
}
</style>
