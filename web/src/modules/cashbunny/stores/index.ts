import type { Dayjs } from 'dayjs'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { APIResponse } from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import {
  type AccountDto,
  type CreateAccountDto,
  type CreateTransactionDto,
  type GetAllCurrenciesDto,
  type OverviewDto,
  type PlannerParametersDto,
  type TransactionDto,
  type UpdateAccountDto,
  type UpdateTransactionDto,
  type UserPreferencesDto,
} from '../models/dto'

export const useCashbunnyStore = defineStore('cashbunny', () => {
  // Currency Code and Grapheme pair
  const currencies = ref<Record<string, string>>({})
  const userPreferences = ref<UserPreferencesDto | null>(null)

  const setCurrencies = (dto: GetAllCurrenciesDto) => {
    for (const [code, grapheme] of Object.entries(dto.currencies_and_grapheme)) {
      currencies.value[code] = grapheme
    }
  }

  const getOverview = (dateRange?: { from: Dayjs; to: Dayjs }) =>
    api.get<APIResponse<OverviewDto>>(APIEndpoints.v1.secure.cashbunny.overview, {
      params: dateRange
        ? {
            from: dateRange.from.unix(),
            to: dateRange.to.unix(),
          }
        : undefined,
    })

  const getPlannerParameters = () =>
    api.get<APIResponse<PlannerParametersDto>>(APIEndpoints.v1.secure.cashbunny.plannerParameters)

  const getAllCurrencies = () =>
    api.get<APIResponse<GetAllCurrenciesDto>>(APIEndpoints.v1.secure.cashbunny.currencies)

  const getUserPreferences = () =>
    api.get<APIResponse<UserPreferencesDto>>(APIEndpoints.v1.secure.cashbunny.userPreferences)

  const createUserPreferences = () =>
    api.post<APIResponse<UserPreferencesDto>>(APIEndpoints.v1.secure.cashbunny.userPreferences)

  const getAccounts = () =>
    api.get<APIResponse<AccountDto[]>>(APIEndpoints.v1.secure.cashbunny.accounts)

  const createAccount = (data: CreateAccountDto) =>
    api.post<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts, data)

  const updateAccount = (accountId: number, data: UpdateAccountDto) =>
    api.put<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts + `/${accountId}`, data)

  const deleteAccount = (accountId: number) =>
    api.delete<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts + `/${accountId}`)

  const getTransactions = () =>
    api.get<APIResponse<TransactionDto[]>>(APIEndpoints.v1.secure.cashbunny.transactions)

  const createTransaction = (data: CreateTransactionDto) =>
    api.post(APIEndpoints.v1.secure.cashbunny.transactions, data)

  const updateTransaction = (transactionId: number, data: UpdateTransactionDto) =>
    api.put(APIEndpoints.v1.secure.cashbunny.transactions + `/${transactionId}`, data)

  const deleteTransaction = (transactionId: number) =>
    api.delete<APIResponse<null>>(
      APIEndpoints.v1.secure.cashbunny.transactions + `/${transactionId}`,
    )

  return {
    currencies,
    userPreferences,
    setCurrencies,
    getOverview,
    getPlannerParameters,
    getAllCurrencies,
    getUserPreferences,
    createUserPreferences,
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
