import { api, APIEndpoints } from '@/utils/api'
import { defineStore } from 'pinia'
import { type GetAccountsResponse } from '../models/dto'

export const useStore = defineStore('cashbunny', () => {
  // const summary = ref<AccountingSummary | null>(null)
  // const fetchSummary = async () => {
  // const accountingSummary = (
  //   await api.get<V1AccountingSummaryResponse>(APIEndpoints.v1.accountingSummary)
  // ).data
  // summary.value = new AccountingSummary(accountingSummary)
  // }

  const getAccounts = () => api.get<GetAccountsResponse>(APIEndpoints.v1.secure.cashbunny.accounts)

  return { getAccounts }
})
