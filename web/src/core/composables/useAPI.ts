import axios, {
  AxiosError,
  type AxiosInstance,
  type AxiosRequestConfig,
  type AxiosResponse,
} from 'axios'
import type { Dayjs } from 'dayjs'
import type { FrequencyStr } from '@/modules/cashbunny/models/recurrence_rule'
import { nestedSnakeToCamel } from '../utils/object'

// A nested dictionary including all paths and their subpaths
const Endpoints = {
  v1: {
    ping: {},
    repeat: {},
    auth: {
      token: {},
    },
    secure: {
      auth: {
        token: {},
        refresh_token: {},
      },
      user: {},
      system_preferences: {},
      cashbunny: {
        overview: {},
        planner: {
          parameters: {},
        },
        currencies: {},
        user_preferences: {},
        accounts: {},
        transactions: {},
      },
    },
  },
}

// This adds a "path" property to all subpaths, which returns
// the entire path from root to said subpath.
type AugmentedProxy<T> = T extends object
  ? { [K in keyof T]: AugmentedProxy<T[K]> } & { path: string }
  : T

function makeEndpointsProxy<T extends Object>(target: T, traversed: string[] = []) {
  const trap: ProxyHandler<T> = {
    get(target, prop, receiver) {
      const value = Reflect.get(target, prop, receiver)

      if (prop === 'path') {
        return traversed.join('/')
      }

      if (value !== null && typeof value === 'object' && value.constructor.name === 'Object') {
        return makeEndpointsProxy(value, traversed.concat(prop.toString()))
      }

      return value
    },
  }

  return new Proxy(target, trap) as AugmentedProxy<T>
}

// A endpoint path builder helper proxy object (why did I make this)
const APIEndpoints = makeEndpointsProxy(Endpoints)

export class APIError {
  message: string
  validationMessages: { [key: string]: string }
  private handled: boolean

  constructor(message: string, validationMessages: { [key: string]: string }) {
    this.message = message
    this.validationMessages = validationMessages
    this.handled = false
  }

  setHandled() {
    this.handled = true
  }

  get isHandled() {
    return this.handled
  }

  toString() {
    return this.message
  }
}

export type RepeatRequest = {
  message: string
}

export type RepeatResponse = RepeatRequest

export type CreateUserRequestDto = {
  email: string
  username: string
  password: string
}

export type CreateJWTTokenRequest = {
  username: string
  password: string
}

export type UserResponse = {
  id: number
  name: string
  loggedInAt: string // time
  createdAt: string // time
}

export type GetUserResponse = {
  user: UserResponse
}

export type UserSystemPreferenceResponse = {
  language: string | null
}

export type GetUserSystemPreferenceResponse = {
  userSystemPreference: UserSystemPreferenceResponse
}

export type UpdateUserSystemPreferenceRequest = UserSystemPreferenceResponse

export type CashbunnyUserPreferenceResponse = {
  userCurrencies: string[]
}

export type GetCashbunnyUserPreferenceResponse = {
  userPreference: CashbunnyUserPreferenceResponse
}

export type GetCashbunnySupportedCurrenciesResponse = {
  currenciesAndGrapheme: { [key: string]: string }
}

export type APIErrorHandlers<T extends number[]> = {
  [key in T[number]]: (error: APIError) => void
}

export type GetOverviewResponse = {
  netWorth: { [key: string]: string }
  profitLossSummary: { [key: string]: { revenue: string; expense: string; profit: string } }
  assetAccounts: AccountResponse[]
  liabilityAccounts: AccountResponse[]
  transactions: TransactionResponse[]
  transactionsFromScheduled: TransactionResponse[]
}

export type AccountResponse = {
  id: number
  category: string
  name: string
  description: string
  amount: number
  amountDisplay: string
  currency: string
  type: string
  createdAt: string
  updatedAt: string
}

export type TransactionResponse = {
  id: number
  description: string
  amount: number
  currency: string
  amountDisplay: string
  transactedAt: string
  createdAt: string
  updatedAt: string
  sourceAccountId: number
  sourceAccountName: string
  destinationAccountId: number
  destinationAccountName: string
  scheduledTransaction: ScheduledTransactionResponse | null
}

export type ScheduledTransactionResponse = {
  id: number
  description: string
  amount: number
  currency: string
  amountDisplay: string
  createdAt: string
  updatedAt: string
  recurrenceRule: RecurrenceRuleResponse
  sourceAccountId: number
  sourceAccountName: string
  destinationAccountId: number
  destinationAccountName: string
}

export type RecurrenceRuleResponse = {
  freq: FrequencyStr
  dtstart: string
  count: number
  interval: number
  until: string
}

export class API {
  ax: AxiosInstance
  constructor(ax: AxiosInstance) {
    this.ax = ax
  }

  isError<T>(res: T | APIError) {
    return res instanceof APIError
  }

  ping() {
    return this.wrapRequest(this.ax.get(APIEndpoints.v1.ping.path))
  }

  repeat(data: RepeatRequest, errorHandlers: APIErrorHandlers<[400]>) {
    return this.wrapRequest(
      this.ax.post<RepeatResponse>(APIEndpoints.v1.repeat.path, data),
      errorHandlers,
    )
  }

  createUser(data: CreateUserRequestDto, errorHandlers: APIErrorHandlers<[400]>) {
    return this.wrapRequest(this.ax.post(APIEndpoints.v1.auth.path, data), errorHandlers)
  }

  createJWTToken(data: CreateJWTTokenRequest, errorHandlers: APIErrorHandlers<[400, 403]>) {
    return this.wrapRequest(this.ax.post(APIEndpoints.v1.auth.token.path, data), errorHandlers)
  }

  deleteJWTToken() {
    return this.wrapRequest(this.ax.delete(APIEndpoints.v1.secure.auth.token.path))
  }

  refreshJWTToken() {
    return this.wrapRequest(this.ax.put(APIEndpoints.v1.secure.auth.token.path))
  }

  createJWTRefreshToken(errorHandlers: APIErrorHandlers<[403]>) {
    return this.wrapRequest(
      this.ax.post(APIEndpoints.v1.secure.auth.refresh_token.path),
      errorHandlers,
    )
  }

  deleteJWTRefreshToken() {
    return this.wrapRequest(this.ax.delete(APIEndpoints.v1.secure.auth.refresh_token.path))
  }

  getUser() {
    return this.wrapRequest(this.ax.get<GetUserResponse>(APIEndpoints.v1.secure.user.path))
  }

  getUserSystemPreference(errorHandlers: APIErrorHandlers<[404]>) {
    return this.wrapRequest(
      this.ax.get<GetUserSystemPreferenceResponse>(APIEndpoints.v1.secure.system_preferences.path),
      errorHandlers,
    )
  }

  createDefaultUserSystemPreference() {
    return this.wrapRequest(
      this.ax.post<GetUserSystemPreferenceResponse>(APIEndpoints.v1.secure.system_preferences.path),
    )
  }

  updateUserSystemPreference(data: UpdateUserSystemPreferenceRequest) {
    return this.wrapRequest(
      this.ax.put<GetUserSystemPreferenceResponse>(
        APIEndpoints.v1.secure.system_preferences.path,
        data,
      ),
    )
  }

  getCashbunnyUserPreference(errorHandlers: APIErrorHandlers<[404]>) {
    return this.wrapRequest(
      this.ax.get<GetCashbunnyUserPreferenceResponse>(
        APIEndpoints.v1.secure.cashbunny.user_preferences.path,
      ),
      errorHandlers,
    )
  }

  createCashbunnyDefaultUserPreference() {
    return this.wrapRequest(
      this.ax.post<GetCashbunnyUserPreferenceResponse>(
        APIEndpoints.v1.secure.cashbunny.user_preferences.path,
      ),
    )
  }

  getCashbunnySupportedCurrencies() {
    return this.wrapRequest(
      this.ax.get<GetCashbunnySupportedCurrenciesResponse>(
        APIEndpoints.v1.secure.cashbunny.currencies.path,
      ),
    )
  }

  getOverview(dateRange?: { from: Dayjs; to: Dayjs }) {
    return this.wrapRequest(
      this.ax.get<GetOverviewResponse>(APIEndpoints.v1.secure.cashbunny.overview.path, {
        params: dateRange
          ? {
              from: dateRange.from.unix(),
              to: dateRange.to.unix(),
            }
          : undefined,
      }),
    )
  }

  // Enforces error handling of known error responses, and converts snake_case response to camelCase
  private async wrapRequest<T extends Object, K extends number[]>(
    promise: Promise<AxiosResponse<T, any>>,
    errorHandlers?: APIErrorHandlers<K>,
  ) {
    try {
      const res = await promise
      return nestedSnakeToCamel(res.data) as T
    } catch (error) {
      if (
        error instanceof AxiosError &&
        error.status &&
        errorHandlers &&
        error.status in errorHandlers
      ) {
        const errorData = error.response?.data
        if ('message' in errorData && 'validation_messages' in errorData) {
          const apiErr = new APIError(errorData.message, errorData.validation_messages)
          const handler = errorHandlers[error.status as K[number]]
          if (handler) {
            handler(apiErr)
            apiErr.setHandled()
          }
          throw apiErr
        }
      }
      throw error
    }
  }
}

export function useAPI(baseURL: string) {
  const ax = axios.create({
    baseURL,
    timeout: 10000,
    withCredentials: true,
  })

  ax.interceptors.response.use(
    (res) => {
      return res
    },
    async (err: AxiosError) => {
      if (err.response?.status === 401 && location.pathname !== '/login') {
        try {
          // If a token has expired, try refreshing it first
          const refreshRes = await ax.put(APIEndpoints.v1.auth.token.path, null, {
            validateStatus: null,
          })

          if (refreshRes.status === 401) {
            throw new Error('Refresh token is either invalid or expired')
          }

          // Retrieve a new refresh token
          // TODO: Error handling when creating refresh token fails
          await ax.post(APIEndpoints.v1.secure.auth.refresh_token.path, null, {
            validateStatus: null,
          })

          // Retry original request
          return ax.request(err.config as AxiosRequestConfig)
        } catch (refreshErr) {
          // Refresh token is invalid
          window.location.replace('/login')
        }
      }
      return Promise.reject(err)
    },
  )

  return new API(ax)
}
