import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { APIResponse } from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import type { Currency } from '../models/currency'
import {
  type AccountDto,
  type CreateAccountDto,
  type CreateTransactionDto,
  type GetAllCurrenciesDto,
  type TransactionDto,
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

  const deleteAccount = (accountId: number) =>
    api.delete<APIResponse<null>>(APIEndpoints.v1.secure.cashbunny.accounts + `/${accountId}`)

  const getTransactions = () =>
    api.get<APIResponse<TransactionDto[]>>(APIEndpoints.v1.secure.cashbunny.transactions)

  const createTransaction = (data: CreateTransactionDto) =>
    api.post(APIEndpoints.v1.secure.cashbunny.transactions, data)

  const deleteTransaction = (transactionId: number) =>
    api.delete<APIResponse<null>>(
      APIEndpoints.v1.secure.cashbunny.transactions + `/${transactionId}`,
    )

  return {
    currencies,
    userPreferences,
    setCurrencies,
    getAllCurrencies,
    getUserPreferences,
    createUserPreferences,
    getAccounts,
    createAccount,
    deleteAccount,
    getTransactions,
    createTransaction,
    deleteTransaction,
  }
})
