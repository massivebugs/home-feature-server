<template>
  <div class="cashbunny-planner-wizard__datatable">
    <DataTableComponent
      :columns="[
        {
          data: 'description',
          title: t('cashbunny.planner.liability.description'),
          responsivePriority: 1,
        },
        {
          data: 'amount',
          title: t('cashbunny.planner.liability.amount'),
          responsivePriority: 2,
          render: (amount: number) => amount.toLocaleString(),
        },
        {
          data: 'currency',
          title: t('cashbunny.planner.liability.currency'),
          responsivePriority: 2,
        },
        {
          data: 'source_asset_account_name',
          title: t('cashbunny.planner.liability.from'),
        },
        {
          data: 'destination_liability_account_name',
          title: t('cashbunny.planner.liability.to'),
        },
        {
          data: 'recurrence_rule',
          title: t('cashbunny.planner.liability.recurrence'),
          render: (recurrenceRule: RecurrenceRuleDto) =>
            new RecurrenceRule(recurrenceRule).toHumanFriendlyString(),
        },
      ]"
      :data="props.data"
      :action-column="true"
      @edit-row="
        (payload: DataTableRowEditEvent<PlannerLiabilityDto>) => {
          emit('editRow', payload.row)
        }
      "
      @delete-rows="
        (payload: DataTableRowsDeleteEvent<PlannerLiabilityDto>) => {
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
import type { PlannerLiabilityDto, RecurrenceRuleDto } from '../models/dto'
import { RecurrenceRule } from '../models/recurrence_rule'

const emit = defineEmits<{
  (e: 'editRow', payload: PlannerLiabilityDto): void
  (e: 'deleteRow', payload: PlannerLiabilityDto): void
}>()

const { t } = useI18n()
const props = defineProps<{
  data: PlannerLiabilityDto[]
}>()
</script>
