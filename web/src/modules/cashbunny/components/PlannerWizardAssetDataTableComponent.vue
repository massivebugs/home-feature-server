<template>
  <div class="cashbunny-planner-wizard__datatable">
    <DataTableComponent
      :columns="[
        { data: 'name', title: t('cashbunny.planner.asset.name'), responsivePriority: 1 },
        {
          data: 'description',
          title: t('cashbunny.planner.asset.description'),
          responsivePriority: 3,
        },
        {
          data: 'amount',
          title: t('cashbunny.planner.asset.amount'),
          responsivePriority: 2,
          render: (amount: number) => amount.toLocaleString(),
        },
        { data: 'currency', title: t('cashbunny.planner.asset.currency'), responsivePriority: 2 },
      ]"
      :data="props.data"
      :action-column="true"
      @edit-row="
        (payload: DataTableRowEditEvent<PlannerAssetDto>) => {
          emit('editRow', payload.row)
        }
      "
      @delete-rows="
        (payload: DataTableRowsDeleteEvent<PlannerAssetDto>) => {
          emit('deleteRow', payload.rows[0])
        }
      "
    />
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DataTableComponent, {
  type DataTableRowEditEvent,
  type DataTableRowsDeleteEvent,
} from '@/core/components/DataTableComponent.vue'
import type { PlannerAssetDto } from '../models/dto'

const emit = defineEmits<{
  (e: 'editRow', payload: PlannerAssetDto): void
  (e: 'deleteRow', payload: PlannerAssetDto): void
}>()

const props = defineProps<{
  data: PlannerAssetDto[]
}>()

const { t } = useI18n()
</script>
