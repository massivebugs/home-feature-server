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
      @click-close="onCloseConfirmDeleteDialog"
      @click-cancel="onCloseConfirmDeleteDialog"
      :pos="new RelativePosition(40, 40)"
      :size="new RelativeSize(20, 20)"
      :title="t('cashbunny.accountDeleteConfirmTitle')"
      :message="t('cashbunny.accountDeleteConfirmMessage')"
      :blocking="true"
    />
    <AccountFormDialogComponent
      v-if="showAccountFormDialog"
      :pos="new RelativePosition(25, 25)"
      :size="new RelativeSize(50, 50)"
      :title="t('cashbunny.addAccount')"
      @click-submit="onClickAddAccountSubmit"
      @click-cancel="onClickAddAccountCancel"
      @click-close="onClickAddAccountCancel"
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
    data: 'category.name',
    title: 'Category',
  },
  {
    data: 'name',
    title: 'Name',
  },
  {
    data: 'description',
    title: 'Description',
  },
  {
    data: 'balance',
    title: 'Balance',
  },
  {
    data: 'currency',
    title: 'Currency',
  },
  {
    data: 'created_at',
    title: 'Created at',
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'updated_at',
    title: 'Updated at',
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
]

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
    const clickedData: AccountDto = dt.row(this).data()
    const selectedData: AccountDto[] = dt.rows({ selected: true }).data() as any

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
                onRowClickEdit(clickedData)
              },
            },
            {
              label: 'Delete',
              shortcutKey: 'Del',
              isDisabled: false,
              onClick: () => {
                onRowClickDelete(selectedData.length ? selectedData : [clickedData])
                showConfirmDeleteDialog.value = true
              },
            },
          ],
        ],
      },
      contextMenuPos,
    )
  })
})

const onClickAddAccount = () => {
  showAccountFormDialog.value = true
}

const onClickAddAccountSubmit = () => {
  //
}

const onClickAddAccountCancel = () => {
  showAccountFormDialog.value = false
}

const onRowClickEdit = (accountDto: AccountDto) => {
  console.log('edit', accountDto)
}

const onRowClickDelete = (accountDtos: AccountDto[]) => {
  console.log('delete', accountDtos)
}

const onCloseConfirmDeleteDialog = () => {
  showConfirmDeleteDialog.value = false
}
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
