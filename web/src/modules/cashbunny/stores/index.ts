import { defineStore } from 'pinia'
import type { APIResponse } from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import { type CreateAccountDto, type GetAccountsResponse } from '../models/dto'

export const useCashbunnyStore = defineStore('cashbunny', () => {
  const getAccounts = () => api.get<GetAccountsResponse>(APIEndpoints.v1.secure.cashbunny.accounts)

  const createAccount = (data: CreateAccountDto) =>
    api.post<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts, data)

  const deleteAccount = (accountId: number) =>
    api.delete<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts + `/${accountId}`)

  return { getAccounts, createAccount, deleteAccount }
})
