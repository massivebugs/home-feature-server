import { defineStore } from 'pinia'
import { ref } from 'vue'
import type {
  CashbunnyTransactionResponse,
  CashbunnyUserPreferenceResponse,
  GetCashbunnySupportedCurrenciesResponse,
} from '@/core/composables/useAPI'
import type { APIResponse } from '@/core/models/dto'
import { APIEndpoints, api } from '@/utils/api'
import {
  type CreateTransactionDto,
  type PlannerParametersDto,
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

  const getTransactions = () =>
    api.get<APIResponse<CashbunnyTransactionResponse[]>>(
      APIEndpoints.v1.secure.cashbunny.transactions.path,
    )

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
    getTransactions,
    createTransaction,
    updateTransaction,
    deleteTransaction,
  }
})
