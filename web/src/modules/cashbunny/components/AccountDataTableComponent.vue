<template>
  <div>
    <div class="controls">
      <div></div>
      <button @click="onClickAddAccount">{{ t('cashbunny.addAccount') }}</button>
    </div>
    <DataTable
      :columns="columns"
      :data="data"
      :options="options"
      ref="table"
      class="table display nowrap compact"
    />
    <ConfirmDialogComponent
      v-if="showConfirmDeleteDialog"
      @click-success="onSuccessConfirmDeleteDialog"
      @click-close="onCloseConfirmDeleteDialog"
      @click-cancel="onCloseConfirmDeleteDialog"
      :pos="new RelativePosition(40, 40)"
      :size="new RelativeSize(20, 20)"
      :title="t('cashbunny.accountDeleteConfirmTitle')"
      :message="t('cashbunny.accountDeleteConfirmMessage', getTargetRowData().length)"
      :blocking="true"
    />
    <AccountFormDialogComponent
      v-if="showAccountFormDialog"
      :pos="new RelativePosition(25, 25)"
      :size="new RelativeSize(50, 50)"
      :title="t('cashbunny.addAccount')"
      :next-account-index="data.length"
      @success="onAccountFormSuccess"
      @click-cancel="onAccountFormCancel"
      @click-close="onAccountFormCancel"
    />
  </div>
</template>

<script setup lang="ts">
import DataTablesCore from 'datatables.net'
import type { Api, Config, ConfigColumns } from 'datatables.net-dt'
import 'datatables.net-responsive'
import 'datatables.net-select'
import DataTable from 'datatables.net-vue3'
import { inject, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ConfirmDialogComponent from '@/core/components/ConfirmDialogComponent.vue'
import { AbsolutePosition } from '@/core/models/absolute_position'
import { RelativePosition } from '@/core/models/relative_position'
import { RelativeSize } from '@/core/models/relative_size'
import type { SetContextMenu } from '@/core/views/DesktopView.vue'
import type { AccountDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import AccountFormDialogComponent from './AccountFormDialogComponent.vue'

DataTable.use(DataTablesCore)

const { t } = useI18n()
const store = useCashbunnyStore()
const table = ref()
const setContextMenu = inject('setContextMenu') as SetContextMenu
let dt: Api
const data = ref<AccountDto[]>([])
const showConfirmDeleteDialog = ref<boolean>(false)
const showAccountFormDialog = ref<boolean>(false)
const clickedData = ref<AccountDto>()
const selectedData = ref<AccountDto[]>([])

const layoutOptions = {
  topStart: {
    pageLength: {},
  },
  topEnd: {
    search: {},
  },
}

const options: Config = {
  responsive: true,
  select: true,
  layout: {
    ...(layoutOptions as any),
  },
  columnDefs: [
    // {
    //   className: 'dt-head-right',
    //   targets: '_all',
    // },
    // {
    //   className: 'dt-body-left',
    //   targets: '_all',
    // },
  ],
}

const columns: ConfigColumns[] = [
  {
    data: 'id',
    title: 'ID',
  },
  {
    data: 'category',
    title: 'Category',
    render: function (data: string, _, row: AccountDto) {
      return `${data} (${row.type})`
    },
  },
  {
    data: 'name',
    title: t('cashbunny.accountName'),
  },
  {
    data: 'description',
    title: t('cashbunny.accountDescription'),
  },
  {
    data: 'balance',
    title: t('cashbunny.accountBalance'),
  },
  {
    data: 'currency',
    title: t('cashbunny.accountCurrency'),
  },
  {
    data: 'created_at',
    title: t('cashbunny.accountCreatedAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'updated_at',
    title: t('cashbunny.accountUpdatedAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
]

const getTargetRowData = () => {
  return selectedData.value.length
    ? selectedData.value
    : clickedData.value
      ? [clickedData.value]
      : []
}

const onClickAddAccount = () => {
  showAccountFormDialog.value = true
}

const onAccountFormSuccess = async () => {
  showAccountFormDialog.value = false
  const res = await store.getAccounts()
  if (res.data.error === null) {
    data.value = res.data.data
  }
}

const onAccountFormCancel = () => {
  showAccountFormDialog.value = false
}

const onRowClickEdit = () => {
  console.log('edit', clickedData.value)
}

const onRowClickDelete = () => {
  showConfirmDeleteDialog.value = true
}

const onSuccessConfirmDeleteDialog = async () => {
  showConfirmDeleteDialog.value = false

  // TODO
  const rows = getTargetRowData()
  await Promise.all([...rows.map((info) => store.deleteAccount(info.id))])

  const res = await store.getAccounts()
  if (res.data.error === null) {
    data.value = res.data.data
  }
}

const onCloseConfirmDeleteDialog = async () => {
  showConfirmDeleteDialog.value = false
}

onMounted(async () => {
  const res = await store.getAccounts()
  if (res.data.error === null) {
    data.value = res.data.data
  }
  dt = table.value.dt
  // TODO: Add a hook to the window's context
  dt.responsive.recalc()

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
              onClick: onRowClickEdit,
            },
            {
              label: 'Delete',
              shortcutKey: 'Del',
              isDisabled: false,
              onClick: onRowClickDelete,
            },
          ],
        ],
      },
      contextMenuPos,
    )
  })
})
</script>

<style lang="scss">
@import 'datatables.net-dt';
@import 'datatables.net-responsive-dt';
@import 'datatables.net-select-dt';
</style>

<style scoped lang="scss">
:deep(.table) {
  max-width: 100%;
}

.table-action-btn {
  margin-right: 5px;
}

.controls {
  display: flex;
  justify-content: space-between;
}
</style>
