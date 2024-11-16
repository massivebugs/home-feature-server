<template>
  <div class="cashbunny-planner-wizard__datatable">
    <DataTableComponent
      :columns="[
        {
          data: 'description',
          title: t('cashbunny.planner.revenue.description'),
          responsivePriority: 1,
        },
        {
          data: 'amount',
          title: t('cashbunny.planner.revenue.amount'),
          responsivePriority: 2,
          render: (amount: number) => amount.toLocaleString(),
        },
        { data: 'currency', title: t('cashbunny.planner.revenue.currency'), responsivePriority: 2 },
        {
          data: 'source_revenue_account_name',
          title: t('cashbunny.planner.revenue.from'),
        },
        {
          data: 'destination_asset_account_name',
          title: t('cashbunny.planner.revenue.to'),
        },
        {
          data: 'recurrence_rule',
          title: t('cashbunny.planner.revenue.recurrence'),
          render: (recurrenceRule: RecurrenceRuleResponse) =>
            new RecurrenceRule(recurrenceRule).toHumanFriendlyString(),
        },
      ]"
      :data="props.data"
      :action-column="true"
      @edit-row="
        (payload: DataTableRowEditEvent<PlannerRevenueDto>) => {
          emit('editRow', payload.row)
        }
      "
      @delete-rows="
        (payload: DataTableRowsDeleteEvent<PlannerRevenueDto>) => {
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
import type { RecurrenceRuleResponse } from '@/core/composables/useAPI'
import type { PlannerRevenueDto } from '../models/dto'
import { RecurrenceRule } from '../models/recurrence_rule'

const emit = defineEmits<{
  (e: 'editRow', payload: PlannerRevenueDto): void
  (e: 'deleteRow', payload: PlannerRevenueDto): void
}>()

const { t } = useI18n()
const props = defineProps<{
  data: PlannerRevenueDto[]
}>()
</script>
