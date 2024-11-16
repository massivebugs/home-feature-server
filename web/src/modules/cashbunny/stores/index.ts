import type { Dayjs } from 'dayjs'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import type {
  CashbunnyUserPreferenceResponse,
  GetCashbunnySupportedCurrenciesResponse,
} from '@/core/composables/useAPI'
import type { APIResponse } from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import {
  type AccountDto,
  type CreateAccountDto,
  type CreateTransactionDto,
  type OverviewDto,
  type PlannerParametersDto,
  type TransactionDto,
  type UpdateAccountDto,
  type UpdateTransactionDto,
} from '../models/dto'

export const useCashbunnyStore = defineStore('cashbunny', () => {
  // Currency Code and Grapheme pair
  const currencies = ref<Record<string, string>>({})
  const userPreference = ref<CashbunnyUserPreferenceResponse | null>(null)

  const setCurrencies = (res: GetCashbunnySupportedCurrenciesResponse) => {
    for (const [code, grapheme] of Object.entries(res.currenciesAndGrapheme)) {
      currencies.value[code] = grapheme
    }
  }

  const getOverview = (dateRange?: { from: Dayjs; to: Dayjs }) =>
    api.get<APIResponse<OverviewDto>>(APIEndpoints.v1.secure.cashbunny.overview.path, {
      params: dateRange
        ? {
            from: dateRange.from.unix(),
            to: dateRange.to.unix(),
          }
        : undefined,
    })

  const getPlannerParameters = () =>
    api.get<APIResponse<PlannerParametersDto>>(
      APIEndpoints.v1.secure.cashbunny.planner.parameters.path,
    )

  const getAccounts = () =>
    api.get<APIResponse<AccountDto[]>>(APIEndpoints.v1.secure.cashbunny.accounts.path)

  const createAccount = (data: CreateAccountDto) =>
    api.post<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts.path, data)

  const updateAccount = (accountId: number, data: UpdateAccountDto) =>
    api.put<APIResponse<null>>(
      APIEndpoints.v1.secure.cashbunny.accounts.path + `/${accountId}`,
      data,
    )

  const deleteAccount = (accountId: number) =>
    api.delete<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts.path + `/${accountId}`)

  const getTransactions = () =>
    api.get<APIResponse<TransactionDto[]>>(APIEndpoints.v1.secure.cashbunny.transactions.path)

  const createTransaction = (data: CreateTransactionDto) =>
    api.post(APIEndpoints.v1.secure.cashbunny.transactions.path, data)

  const updateTransaction = (transactionId: number, data: UpdateTransactionDto) =>
    api.put(APIEndpoints.v1.secure.cashbunny.transactions.path + `/${transactionId}`, data)

  const deleteTransaction = (transactionId: number) =>
    api.delete<APIResponse<null>>(
      APIEndpoints.v1.secure.cashbunny.transactions.path + `/${transactionId}`,
    )

  return {
    currencies,
    userPreference,
    setCurrencies,
    getOverview,
    getPlannerParameters,
    getAccounts,
    createAccount,
    updateAccount,
    deleteAccount,
    getTransactions,
    createTransaction,
    updateTransaction,
    deleteTransaction,
  }
})
