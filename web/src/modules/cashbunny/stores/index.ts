import { defineStore } from 'pinia'
import { ref } from 'vue'
import type {
  CashbunnyUserPreferenceResponse,
  GetCashbunnySupportedCurrenciesResponse,
} from '@/core/composables/useAPI'
import { APIEndpoints, api } from '@/utils/api'
import { type PlannerParametersDto } from '../models/dto'

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
    api.get<PlannerParametersDto>(APIEndpoints.v1.secure.cashbunny.planner.parameters.path)

  return {
    currencies,
    userPreference,
    setCurrencies,
    getPlannerParameters,
  }
})
