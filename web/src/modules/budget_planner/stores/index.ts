import { defineStore } from 'pinia'
import { ref } from 'vue'
import { AccountingSummary } from '@/modules/budget_planner/models/accounting_summary'
import { APIEndpoints, api, type V1AccountingSummaryResponse } from '../../../utils/api'

export const useStore = defineStore('budget_planner', () => {
  const summary = ref<AccountingSummary | null>(null)
  const fetchSummary = async () => {
    const accountingSummary = (
      await api.get<V1AccountingSummaryResponse>(APIEndpoints.v1.accountingSummary)
    ).data

    summary.value = new AccountingSummary(accountingSummary)
  }

  return { summary, fetchSummary }
})
