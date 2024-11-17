<template>
  <div class="cashbunny-account-datatable">
    <div class="controls">
      <div></div>
      <button @click="onClickAddAccount">{{ t('cashbunny.addAccount') }}</button>
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
      :title="t('cashbunny.accountDeleteConfirmTitle')"
      :message="t('cashbunny.accountDeleteConfirmMessage', rowsToDelete.length)"
      :blocking="true"
    />
    <AccountFormDialogComponent
      v-if="isCreate || rowToEdit"
      :api="props.api"
      pos="center"
      :title="t('cashbunny.addAccount')"
      :next-account-index="data.length"
      :account="rowToEdit ?? undefined"
      @success="onAccountFormSuccess"
      @click-cancel="onAccountFormCancel"
      @click-close="onAccountFormCancel"
    />
  </div>
</template>

<script setup lang="ts">
import type { ConfigColumns } from 'datatables.net-dt'
import 'datatables.net-responsive'
import 'datatables.net-select'
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ConfirmDialogComponent from '@/core/components/ConfirmDialogComponent.vue'
import DataTableComponent, {
  type DataTableRowEditEvent,
  type DataTableRowsDeleteEvent,
} from '@/core/components/DataTableComponent.vue'
import type { API, CashbunnyAccountResponse } from '@/core/composables/useAPI'
import AccountFormDialogComponent from './AccountFormDialogComponent.vue'

const props = defineProps<{
  api: API
}>()

const { t } = useI18n()
const data = ref<CashbunnyAccountResponse[]>([])
const isCreate = ref<boolean>(false)
const rowsToDelete = ref<CashbunnyAccountResponse[] | null>(null)
const rowToEdit = ref<CashbunnyAccountResponse | null>(null)

const columns: ConfigColumns[] = [
  {
    data: 'id',
    title: 'ID',
  },
  {
    data: 'category',
    title: 'Category',
    render: function (data: string, _, row: CashbunnyAccountResponse) {
      return `${data} (${row.type})`
    },
  },
  {
    data: 'name',
    title: t('cashbunny.account.name'),
  },
  {
    data: 'description',
    title: t('cashbunny.account.description'),
  },
  {
    data: 'amountDisplay',
    title: t('cashbunny.account.amount'),
  },
  {
    data: 'currency',
    title: t('cashbunny.account.currency'),
  },
  {
    data: 'createdAt',
    title: t('cashbunny.account.createdAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
  {
    data: 'updatedAt',
    title: t('cashbunny.account.updatedAt'),
    render: function (data: string) {
      return new Date(data).toLocaleString(navigator.language)
    },
  },
]

const onClickAddAccount = () => {
  isCreate.value = true
}

const onAccountFormSuccess = async () => {
  isCreate.value = false
  rowToEdit.value = null
  const res = await props.api.getCashbunnyAccounts()
  data.value = res.accounts
}

const onAccountFormCancel = () => {
  isCreate.value = false
  rowToEdit.value = null
}

const onRowEdit = ({ row }: DataTableRowEditEvent<CashbunnyAccountResponse>) => {
  rowToEdit.value = row
}

const onRowsDelete = ({ rows }: DataTableRowsDeleteEvent<CashbunnyAccountResponse>) => {
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
    ...rowsToDelete.value.map((info: CashbunnyAccountResponse) =>
      props.api.deleteCashbunnyAccount(info.id),
    ),
  ])

  rowsToDelete.value = null

  const res = await props.api.getCashbunnyAccounts()
  data.value = res.accounts
}

onMounted(async () => {
  const res = await props.api.getCashbunnyAccounts()
  data.value = res.accounts
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';
.cashbunny-account-datatable {
  background-color: colors.$high-opacity-white;
  height: 100%;
  padding: 1em;
}

.controls {
  display: flex;
  justify-content: space-between;
}
</style>
