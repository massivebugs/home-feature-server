<template>
  <div class="cashbunny-transactions-datatable">
    <div class="controls">
      <div></div>
      <button @click="onClickAddTransaction">{{ t('cashbunny.createTransaction') }}</button>
    </div>
    <DataTableComponent
      :columns="columns"
      :data="data"
      :info="true"
      :paging="true"
      :searching="true"
      :ordering="true"
      :select="true"
      @edit-row="onRowEdit"
      @delete-rows="onRowsDelete"
    />
    <ConfirmDialogComponent
      v-if="rowsToDelete"
      @click-success="onSuccessConfirmDeleteDialog"
      @click-close="onCloseConfirmDeleteDialog"
      @click-cancel="onCloseConfirmDeleteDialog"
      pos="center"
      :title="t('cashbunny.transactionDeleteConfirmTitle')"
      :message="t('cashbunny.transactionDeleteConfirmMessage', rowsToDelete.length)"
      :blocking="true"
    />
    <TransactionFormDialogComponent
      v-if="isCreate || rowToEdit"
      :api="props.api"
      pos="center"
      :title="t('cashbunny.createTransaction')"
      :transaction="rowToEdit ?? undefined"
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
  type DataTableRowEditEvent,
  type DataTableRowsDeleteEvent,
} from '@/core/components/DataTableComponent.vue'
import type { API, CashbunnyTransactionResponse } from '@/core/composables/useAPI'
import TransactionFormDialogComponent from './TransactionFormDialogComponent.vue'

const props = defineProps<{
  api: API
}>()

DataTable.use(DataTablesCore)

const { t } = useI18n()
const data = ref<CashbunnyTransactionResponse[]>([])
const isCreate = ref<boolean>(false)
const rowsToDelete = ref<CashbunnyTransactionResponse[] | null>(null)
const rowToEdit = ref<CashbunnyTransactionResponse | null>(null)

const columns: ConfigColumns[] = [
  {
    data: 'id',
    title: 'ID',
  },
  {
    data: 'transactedAt',
    title: 'Transacted at',
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'description',
    title: t('cashbunny.transaction.description'),
  },
  {
    data: 'amountDisplay',
    title: t('cashbunny.transaction.amount'),
  },
  {
    data: 'currency',
    title: t('cashbunny.transaction.currency'),
  },
  {
    data: 'sourceAccountId',
    title: t('cashbunny.transaction.sourceAccount'),
    render: function (data: string, _, row: CashbunnyTransactionResponse) {
      return `${row.sourceAccountName} (${data})`
    },
  },
  {
    data: 'destinationAccountId',
    title: t('cashbunny.transaction.destinationAccount'),
    render: function (data: string, _, row: CashbunnyTransactionResponse) {
      return `${row.destinationAccountName} (${data})`
    },
  },
  {
    data: 'createdAt',
    title: t('cashbunny.transaction.createdAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'updatedAt',
    title: t('cashbunny.transaction.updatedAt'),
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

  const res = await props.api.getCashbunnyTransactions()
  data.value = res.transactions
}

const onTransactionFormCancel = () => {
  isCreate.value = false
  rowToEdit.value = null
}

const onRowEdit = ({ row }: DataTableRowEditEvent<CashbunnyTransactionResponse>) => {
  rowToEdit.value = row
}

const onRowsDelete = ({ rows }: DataTableRowsDeleteEvent<CashbunnyTransactionResponse>) => {
  rowsToDelete.value = rows
}

const onCloseConfirmDeleteDialog = async () => {
  rowsToDelete.value = null
}

const onSuccessConfirmDeleteDialog = async () => {
  if (!rowsToDelete.value) {
    return
  }

  await Promise.all([
    ...rowsToDelete.value.map((info) => props.api.deleteCashbunnyTransaction(info.id)),
  ])

  rowsToDelete.value = null

  const res = await props.api.getCashbunnyTransactions()
  data.value = res.transactions
}

onMounted(async () => {
  const res = await props.api.getCashbunnyTransactions()
  data.value = res.transactions
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.cashbunny-transactions-datatable {
  background-color: colors.$high-opacity-white;
  height: 100%;
  padding: 1em;
}

.controls {
  display: flex;
  justify-content: space-between;
}
</style>
