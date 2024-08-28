<template>
  <div>
    <div class="controls">
      <div></div>
      <button @click="onClickAddTransaction">{{ t('cashbunny.createTransaction') }}</button>
    </div>
    <DataTableComponent
      :columns="columns"
      :data="data"
      @edit-row="onRowEdit"
      @delete-row="onRowDelete"
    />
    <ConfirmDialogComponent
      v-if="rowsToDelete"
      @click-success="onSuccessConfirmDeleteDialog"
      @click-close="onCloseConfirmDeleteDialog"
      @click-cancel="onCloseConfirmDeleteDialog"
      :pos="new RelativePosition(40, 40)"
      :size="new RelativeSize(20, 20)"
      :title="t('cashbunny.transactionDeleteConfirmTitle')"
      :message="t('cashbunny.transactionDeleteConfirmMessage', rowsToDelete.length)"
      :blocking="true"
    />
    <TransactionFormDialogComponent
      v-if="isCreate || rowToEdit"
      :pos="new RelativePosition(25, 25)"
      :size="new RelativeSize(50, 50)"
      :title="t('cashbunny.addTransaction')"
      :transaction="clickedData ?? undefined"
      @success="onTransactionFormSuccess"
      @click-cancel="onTransactionFormCancel"
      @click-close="onTransactionFormCancel"
    />
  </div>
</template>

<script setup lang="ts">
import DataTablesCore from 'datatables.net'
import type { ConfigColumns } from 'datatables.net-dt'
import 'datatables.net-responsive'
import 'datatables.net-select'
import DataTable from 'datatables.net-vue3'
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ConfirmDialogComponent from '@/core/components/ConfirmDialogComponent.vue'
import DataTableComponent, {
  type DataTableRowDeleteEvent,
  type DataTableRowEditEvent,
} from '@/core/components/DataTableComponent.vue'
import { RelativePosition } from '@/core/models/relativePosition'
import { RelativeSize } from '@/core/models/relativeSize'
import type { TransactionDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import TransactionFormDialogComponent from './TransactionFormDialogComponent.vue'

DataTable.use(DataTablesCore)

const { t } = useI18n()
const store = useCashbunnyStore()
const data = ref<TransactionDto[]>([])
const clickedData = ref<TransactionDto | null>(null)
const isCreate = ref<boolean>(false)
const rowsToDelete = ref<TransactionDto[] | null>(null)
const rowToEdit = ref<TransactionDto | null>(null)

const columns: ConfigColumns[] = [
  {
    data: 'id',
    title: 'ID',
  },
  {
    data: 'transacted_at',
    title: 'Transacted at',
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'description',
    title: t('cashbunny.transactionDescription'),
  },
  {
    data: 'amount_display',
    title: t('cashbunny.transactionAmount'),
  },
  {
    data: 'currency',
    title: t('cashbunny.transactionCurrency'),
  },
  {
    data: 'source_account_id',
    title: t('cashbunny.transactionSourceAccount'),
    render: function (data: string, _, row: TransactionDto) {
      return `${row.source_account_name} (${data})`
    },
  },
  {
    data: 'destination_account_id',
    title: t('cashbunny.transactionDestinationAccount'),
    render: function (data: string, _, row: TransactionDto) {
      return `${row.destination_account_name} (${data})`
    },
  },
  {
    data: 'transacted_at',
    title: t('cashbunny.transactionTransactedAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'created_at',
    title: t('cashbunny.transactionCreatedAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'updated_at',
    title: t('cashbunny.transactionUpdatedAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
]

const onClickAddTransaction = () => {
  isCreate.value = true
}

const onTransactionFormSuccess = async () => {
  isCreate.value = false
  rowToEdit.value = null

  const res = await store.getTransactions()
  if (res.data.error === null) {
    data.value = res.data.data
  }
}

const onTransactionFormCancel = () => {
  isCreate.value = false
  rowToEdit.value = null
}

const onRowEdit = ({ row }: DataTableRowEditEvent<TransactionDto>) => {
  rowToEdit.value = row
}

const onRowDelete = ({ rows }: DataTableRowDeleteEvent<TransactionDto>) => {
  rowsToDelete.value = rows
}

const onCloseConfirmDeleteDialog = async () => {
  rowsToDelete.value = null
}

const onSuccessConfirmDeleteDialog = async () => {
  if (!rowsToDelete.value) {
    return
  }

  await Promise.all([...rowsToDelete.value.map((info) => store.deleteTransaction(info.id))])

  rowsToDelete.value = null

  const res = await store.getTransactions()
  if (res.data.error === null) {
    data.value = res.data.data
  }
}

onMounted(async () => {
  const res = await store.getTransactions()
  if (res.data.error === null) {
    data.value = res.data.data
  }
})
</script>

<style scoped lang="scss">
.controls {
  display: flex;
  justify-content: space-between;
}
</style>
