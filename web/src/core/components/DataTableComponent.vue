<template>
  <DataTable
    :columns="columns"
    :data="data"
    :options="dtConfig"
    ref="table"
    class="table display nowrap compact hfs-datatable"
  />
</template>

<script setup lang="ts">
import DataTablesCore from 'datatables.net'
import type { Api, Config, ConfigColumns } from 'datatables.net-dt'
import 'datatables.net-responsive'
import 'datatables.net-select'
import DataTable from 'datatables.net-vue3'
import { inject, onMounted, onUnmounted, ref } from 'vue'
import { type ToggleWindowResizeHandlerFunc } from '@/core/components/WindowComponent.vue'
import { AbsolutePosition } from '@/core/models/absolutePosition'
import type { SetContextMenu } from '@/core/views/DesktopView.vue'

DataTable.use(DataTablesCore)

const table = ref()
const setContextMenu = inject('setContextMenu') as SetContextMenu
let dt: Api
const showConfirmDeleteDialog = ref<boolean>(false)
const clickedData = ref<Object | null>(null)
const selectedData = ref<Object[]>([])
const addWindowResizeListener = inject('addWindowResizeListener') as ToggleWindowResizeHandlerFunc
const removeWindowResizeListener = inject(
  'removeWindowResizeListener',
) as ToggleWindowResizeHandlerFunc

export type DataTableClickEvent<T> = {
  row: T
}

export type DataTableSelectEvent<T> = {
  rows: T[]
}

export type DataTableRowEditEvent<T> = {
  row: T
}

export type DataTableRowDeleteEvent<T> = {
  rows: T[]
}

export type DataTableLoadedEvent = {
  resizeFunc: () => void
}

const emit = defineEmits<{
  (e: 'editRow', payload: DataTableRowEditEvent<any>): void
  (e: 'deleteRow', payload: DataTableRowDeleteEvent<any>): void
  (e: 'loaded', payload: DataTableLoadedEvent): void
}>()

const props = defineProps<{
  data?: Object[]
  columns?: ConfigColumns[]
  options?: Config
}>()

const dtConfig: Config = {
  drawCallback: (settings) => {
    settings.api.responsive.recalc()
  },
  responsive: true,
  select: true,
  layout: {
    topStart: {
      pageLength: {},
    },
    topEnd: {
      search: {},
    },
  },
  ...props.options,
}

const getTargetRowData = () => {
  return selectedData.value.length
    ? selectedData.value
    : clickedData.value
      ? [clickedData.value]
      : []
}

onMounted(async () => {
  dt = table.value.dt

  emit('loaded', { resizeFunc: dt.responsive.recalc })
  addWindowResizeListener(dt.responsive.recalc)

  // Prevent right click and display custom context menu
  dt.on('contextmenu', 'tbody tr', function (e) {
    e.preventDefault()

    // TODO: Types aren't exact here
    clickedData.value = dt.row(this).data()
    selectedData.value = dt.rows({ selected: true }).data().toArray()

    const contextMenuPos = new AbsolutePosition(
      (e as PointerEvent).clientX,
      (e as PointerEvent).clientY,
    )

    setContextMenu(
      {
        itemGroups: [
          [
            {
              label: 'Edit',
              isDisabled: false,
              onClick: () => {
                if (clickedData.value) {
                  emit('editRow', { row: clickedData.value })
                }
              },
            },
            {
              label: 'Delete',
              shortcutKey: 'Del',
              isDisabled: false,
              onClick: () => {
                showConfirmDeleteDialog.value = false
                const rows = getTargetRowData()
                emit('deleteRow', { rows })
              },
            },
          ],
        ],
      },
      contextMenuPos,
    )
  })
})

onUnmounted(() => {
  removeWindowResizeListener(dt.responsive.recalc)
})
</script>

<style lang="scss">
@import 'datatables.net-dt';
@import 'datatables.net-responsive-dt';
@import 'datatables.net-select-dt';
</style>
