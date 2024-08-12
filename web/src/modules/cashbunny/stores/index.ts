import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { APIResponse } from '@/core/models/dto'
import { formatDateToDateOnly } from '@/core/utils/time'
import { APIEndpoints, api } from '@/utils/api'
import type { Currency } from '../models/currency'
import {
  type AccountDto,
  type CreateAccountDto,
  type CreateTransactionDto,
  type GetAllCurrenciesDto,
  type OverviewDto,
  type TransactionDto,
  type UpdateAccountDto,
  type UpdateTransactionDto,
  type UserPreferencesDto,
} from '../models/dto'

export const useCashbunnyStore = defineStore('cashbunny', () => {
  const currencies = ref<Currency[]>([])
  const userPreferences = ref<UserPreferencesDto | null>(null)

  const setCurrencies = (dto: GetAllCurrenciesDto) => {
    for (const [key, value] of Object.entries(dto.currencies_and_grapheme)) {
      currencies.value.push({ code: key, grapheme: value })
    }
  }

  const getOverview = (dateRange?: { from: Date; to: Date }) =>
    api.get<APIResponse<OverviewDto>>(APIEndpoints.v1.secure.cashbunny.overview, {
      params: dateRange
        ? {
            from_date: formatDateToDateOnly(dateRange.from),
            to_date: formatDateToDateOnly(dateRange.to),
          }
        : undefined,
    })

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
