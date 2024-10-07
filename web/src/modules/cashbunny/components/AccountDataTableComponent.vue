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
import type { AccountDto } from '../models/dto'
import { useCashbunnyStore } from '../stores'
import AccountFormDialogComponent from './AccountFormDialogComponent.vue'

const { t } = useI18n()
const store = useCashbunnyStore()
const data = ref<AccountDto[]>([])
const isCreate = ref<boolean>(false)
const rowsToDelete = ref<AccountDto[] | null>(null)
const rowToEdit = ref<AccountDto | null>(null)

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
    data: 'amount',
    title: t('cashbunny.accountAmount'),
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

const onClickAddAccount = () => {
  isCreate.value = true
}

const onAccountFormSuccess = async () => {
  isCreate.value = false
  rowToEdit.value = null
  const res = await store.getAccounts()
  if (res.data.error === null) {
    data.value = res.data.data
  }
}

const onAccountFormCancel = () => {
  isCreate.value = false
  rowToEdit.value = null
}

const onRowEdit = ({ row }: DataTableRowEditEvent<AccountDto>) => {
  rowToEdit.value = row
}

const onRowsDelete = ({ rows }: DataTableRowsDeleteEvent<AccountDto>) => {
  rowsToDelete.value = rows
}

const onCloseConfirmDeleteDialog = async () => {
  rowsToDelete.value = null
}

const onSuccessConfirmDeleteDialog = async () => {
  if (!rowsToDelete.value) {
    return
  }

  await Promise.all([...rowsToDelete.value.map((info: AccountDto) => store.deleteAccount(info.id))])

  rowsToDelete.value = null

  const res = await store.getAccounts()
  if (res.data.error === null) {
    data.value = res.data.data
  }
}

onMounted(async () => {
  const res = await store.getAccounts()
  if (res.data.error === null) {
    data.value = res.data.data
  }
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
