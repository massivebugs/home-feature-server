import { defineStore } from 'pinia'
import { ref } from 'vue'
import type {
  AccountResponse,
  CashbunnyUserPreferenceResponse,
  GetCashbunnySupportedCurrenciesResponse,
  TransactionResponse,
} from '@/core/composables/useAPI'
import type { APIResponse } from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import {
  type CreateAccountDto,
  type CreateTransactionDto,
  type PlannerParametersDto,
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

  const getPlannerParameters = () =>
    api.get<APIResponse<PlannerParametersDto>>(
      APIEndpoints.v1.secure.cashbunny.planner.parameters.path,
    )

  const getAccounts = () =>
    api.get<APIResponse<AccountResponse[]>>(APIEndpoints.v1.secure.cashbunny.accounts.path)

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
    api.get<APIResponse<TransactionResponse[]>>(APIEndpoints.v1.secure.cashbunny.transactions.path)

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
