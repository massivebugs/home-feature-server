<template>
  <!-- https://vue-land.github.io/faq/forwarding-slots#passing-all-slots -->
  <DataTable
    :columns="tableColumns"
    :data="data"
    :options="dtConfig"
    ref="table"
    class="table display nowrap compact hfs-datatable"
  >
    <template v-for="(_, slotName) in $slots" v-slot:[slotName]="slotProps">
      <slot :name="slotName" v-bind="slotProps ?? {}" />
    </template>
    <template #column-action="row">
      <div class="datatable__action-group">
        <ButtonComponent
          class="datatable__action-button"
          type="warning"
          @click="emit('editRow', { row: row.rowData })"
        >
          <EditIconComponent />
        </ButtonComponent>
        <ButtonComponent
          class="datatable__action-button"
          type="danger"
          @click="emit('deleteRows', { rows: [row.rowData] })"
        >
          <TrashIconComponent />
        </ButtonComponent>
      </div>
    </template>
  </DataTable>
</template>

<script setup lang="ts">
import DataTablesCore from 'datatables.net-dt'
import type { Api, Config, ConfigColumns } from 'datatables.net-dt'
import 'datatables.net-responsive'
import 'datatables.net-select'
import DataTable from 'datatables.net-vue3'
import { computed, inject, onMounted, onUnmounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { type ToggleWindowResizeHandlerFunc } from '@/core/components/WindowComponent.vue'
import { AbsolutePosition } from '@/core/models/absolutePosition'
import type { SetContextMenu } from '@/core/views/DesktopView.vue'
import ButtonComponent from './ButtonComponent.vue'
import type { ContextMenuOptions } from './ContextMenuComponent.vue'
import EditIconComponent from './EditIconComponent.vue'
import TrashIconComponent from './TrashIconComponent.vue'

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

export type DataTableRowsDeleteEvent<T> = {
  rows: T[]
}

export type DataTableLoadedEvent = {
  resizeFunc: () => void
}

const emit = defineEmits<{
  (e: 'editRow', payload: DataTableRowEditEvent<any>): void
  (e: 'deleteRows', payload: DataTableRowsDeleteEvent<any>): void
  (e: 'loaded', payload: DataTableLoadedEvent): void
}>()

const props = defineProps<{
  data?: Object[]
  columns: ConfigColumns[]
  contextMenu?: ContextMenuOptions | { edit: boolean; delete: boolean }
  info?: boolean
  paging?: boolean
  searching?: boolean
  ordering?: boolean
  select?: boolean
  actionColumn?: boolean
}>()

const { t } = useI18n()
const tableColumns = computed(() => {
  if (props.actionColumn) {
    return [
      ...props.columns,
      { data: null, name: 'action', title: t('common.action'), responsivePriority: 1 },
    ]
  }
  return props.columns
})

const dtConfig: Config = {
  drawCallback: (settings) => {
    settings.api.responsive.recalc()
  },
  responsive: true,
  layout: {
    topStart: {
      pageLength: {},
    },
    topEnd: {
      search: {},
    },
  },
  info: props.info,
  paging: props.paging,
  searching: props.searching,
  ordering: props.ordering,
  select: props.select,
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
                emit('deleteRows', { rows })
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

<style scoped lang="scss">
.datatable__action-group {
  display: flex;
  gap: 5px;
}

.datatable__action-button {
  padding: 0.1em 0.4em;
}
</style>
