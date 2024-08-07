import { defineStore } from 'pinia'
import type { APIResponse } from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import {
  type CreateAccountDto,
  type GetAccountsResponse,
  type GetCategoriesResponse,
} from '../models/dto'

export const useCashbunnyStore = defineStore('cashbunny', () => {
  // const summary = ref<AccountingSummary | null>(null)
  // const fetchSummary = async () => {
  // const accountingSummary = (
  //   await api.get<V1AccountingSummaryResponse>(APIEndpoints.v1.accountingSummary)
  // ).data
  // summary.value = new AccountingSummary(accountingSummary)
  // }

  const getAccounts = () => api.get<GetAccountsResponse>(APIEndpoints.v1.secure.cashbunny.accounts)
  const getAccountCategories = () =>
    api.get<GetCategoriesResponse>(APIEndpoints.v1.secure.cashbunny.categories)
  const createAccount = (data: CreateAccountDto) =>
    api.post<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts, data)
  const deleteAccount = (accountId: number) =>
    api.delete<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts + `/${accountId}`)

  return { getAccounts, getAccountCategories, createAccount, deleteAccount }
})
