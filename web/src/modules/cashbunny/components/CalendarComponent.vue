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
  type EventInput,
  type MoreLinkAction,
} from '@fullcalendar/core/index.js'
import rrulePlugin from '@fullcalendar/rrule'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from '@fullcalendar/interaction'
import timeGridPlugin from '@fullcalendar/timegrid'
import FullCalendarComponent from '@fullcalendar/vue3'
import type { Dayjs } from 'dayjs'
import dayjs from 'dayjs'
import { debounce } from 'lodash'
import { computed, onMounted, ref } from 'vue'

export type CalendarSetDatesEvent = {
  dateStart: Dayjs
  dateEnd: Dayjs
}

export type CalendarSelectDatesEvent = CalendarSetDatesEvent | null

export type CalendarLoadedEvent = {
  resizeFunc: () => void
}

export type CalendarEvent = EventInput

const emit = defineEmits<{
  (e: 'selectDates', payload: CalendarSelectDatesEvent): void
  (e: 'setDates', payload: CalendarSetDatesEvent): void
  (e: 'loaded', payload: CalendarLoadedEvent): void
}>()

const props = defineProps<{
  tabs: CalendarTab[]
  height?: string
  events?: CalendarEvent
  moreLinkAction?: MoreLinkAction
}>()

const fullCalendar = ref()
let fullCalendarApi: CalendarApi
const resizeFunc = debounce(() => fullCalendarApi.updateSize(), 100)
let currentSelection: CalendarSelectDatesEvent | null = null

const onSelect = (arg: DateSelectArg) => {
  // We want dateEnd to not include the next day (we didn't select it!)
  const dateStart = dayjs(arg.start)
  const dateEnd = dayjs(arg.end).subtract(1, 's')

  // Clear date selection if the same date has been picked
  if (
    currentSelection &&
    currentSelection.dateStart?.isSame(dateStart) &&
    currentSelection.dateEnd?.isSame(dateEnd)
  ) {
    fullCalendarApi.unselect()
    currentSelection = null
  } else {
    currentSelection = { dateStart, dateEnd }
  }

  emit('selectDates', currentSelection)
}

const onDatesSet = (arg: DatesSetArg) => {
  // We want dateEnd to not include the next day (we didn't select it!)
  const dateStart = dayjs(arg.start)
  const dateEnd = dayjs(arg.end).subtract(1, 's')

  emit('setDates', { dateStart, dateEnd })
}

const calendarOptions = computed<CalendarOptions>(() => ({
  plugins: [rrulePlugin, dayGridPlugin, timeGridPlugin, interactionPlugin],
  height: props.height,
  initialView: 'dayGridMonth',
  customButtons: {
    myCustomButton: {
      text: 'custom!',
      click: function () {
        alert('clicked the custom button!')
      },
    },
  },
  headerToolbar: {
    start: 'myCustomButton',
    left: 'prev,next,today',
    center: 'title',
    right: props.tabs.join(','),
  },
  editable: true,
  selectable: true,
  selectMirror: true,
  dayMaxEvents: true,
  showNonCurrentDates: false,
  datesSet: onDatesSet,
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
  events: props.events,
  moreLinkClick: props.moreLinkAction,
  eventTimeFormat: {
    hour: '2-digit',
    minute: '2-digit',
    hour12: false,
  },
}))

onMounted(() => {
  fullCalendarApi = fullCalendar.value.getApi()
  emit('loaded', { resizeFunc })
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';
:deep() {
  user-select: none;

  .fc-header-toolbar {
    gap: 5px;
    flex-direction: column;
    .fc-button {
      background-color: colors.$dark-grey;
      text-transform: capitalize;
      padding: 0.1em 0.6em;

      &.fc-button-active {
        background-color: colors.$black;
      }
    }

    .fc-toolbar-title {
      font-size: 1.2em;
    }
  }

  .fc-daygrid-day {
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

    .fc-daygrid-day-top {
      display: flex;
      justify-content: center;
      font-size: 0.8em;
    }

    .fc-daygrid-day-events {
      overflow: hidden;
    }
  }
}
</style>
