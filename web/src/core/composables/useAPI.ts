import axios, {
  AxiosError,
  type AxiosInstance,
  type AxiosRequestConfig,
  type AxiosResponse,
} from 'axios'
import type { Dayjs } from 'dayjs'
import type { FrequencyStr } from '@/modules/cashbunny/models/recurrence_rule'
import { camelToKebab, nestedCamelToSnake, nestedSnakeToCamel } from '../utils/object'

//API request/response models

export type RepeatRequest = {
  message: string
}

export type RepeatResponse = RepeatRequest

export type CreateUserRequest = {
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

export type GetCashbunnyOverviewResponse = {
  netWorth: { [key: string]: string }
  profitLossSummary: { [key: string]: { revenue: string; expense: string; profit: string } }
  assetAccounts: CashbunnyAccountResponse[]
  liabilityAccounts: CashbunnyAccountResponse[]
  transactions: CashbunnyTransactionResponse[]
  transactionsFromScheduled: CashbunnyTransactionResponse[]
}

export type CashbunnyAccountResponse = {
  id: number
  category: string
  name: string
  description: string
  amount: number
  amountDisplay: string
  currency: string
  type: string
  orderIndex: number
  createdAt: string
  updatedAt: string
}

export type CashbunnyTransactionResponse = {
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
  scheduledTransaction: CashbunnyScheduledTransactionResponse | null
}

export type CashbunnyScheduledTransactionResponse = {
  id: number
  description: string
  amount: number
  currency: string
  amountDisplay: string
  createdAt: string
  updatedAt: string
  recurrenceRule: CashbunnyRecurrenceRuleResponse
  sourceAccountId: number
  sourceAccountName: string
  destinationAccountId: number
  destinationAccountName: string
}

export type CashbunnyRecurrenceRuleResponse = {
  freq: FrequencyStr
  dtstart: string
  count: number
  interval: number
  until: string
}

export type GetCashbunnyAccountsResponse = {
  accounts: CashbunnyAccountResponse[]
}

export type CreateCashbunnyAccountRequest = Pick<
  CashbunnyAccountResponse,
  'name' | 'category' | 'description' | 'currency'
> &
  Partial<Pick<CashbunnyAccountResponse, 'orderIndex'>>

export type UpdateCashbunnyAccountRequest = Pick<
  CashbunnyAccountResponse,
  'name' | 'description' | 'orderIndex'
>

export type GetCashbunnyTransactionsResponse = {
  transactions: CashbunnyTransactionResponse[]
}

export type CreateCashbunnyTransactionRequest = Pick<
  CashbunnyTransactionResponse,
  | 'description'
  | 'amount'
  | 'currency'
  | 'sourceAccountId'
  | 'destinationAccountId'
  | 'transactedAt'
>

export type UpdateCashbunnyTransactionRequest = Pick<
  CashbunnyTransactionResponse,
  'description' | 'amount' | 'transactedAt'
>

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
        refreshToken: {},
      },
      user: {},
      systemPreferences: {},
      cashbunny: {
        overview: {},
        planner: {
          parameters: {},
        },
        currencies: {},
        userPreferences: {},
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
        return makeEndpointsProxy(value, traversed.concat(camelToKebab(prop.toString())))
      }

      return value
    },
  }

  return new Proxy(target, trap) as AugmentedProxy<T>
}

// A endpoint path builder helper proxy object (why did I make this)
const APIEndpoints = makeEndpointsProxy(Endpoints)

export class APIError {
  status: number
  message: string
  validationMessages: { [key: string]: string }
  private handled: boolean

  constructor(status: number, message: string, validationMessages: { [key: string]: string }) {
    this.status = status
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

export type APIErrorHandlers<T extends number[]> = {
  [key in T[number]]: (error: APIError) => void
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

  createUser(data: CreateUserRequest, errorHandlers: APIErrorHandlers<[400]>) {
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
      this.ax.post(APIEndpoints.v1.secure.auth.refreshToken.path),
      errorHandlers,
    )
  }

  deleteJWTRefreshToken() {
    return this.wrapRequest(this.ax.delete(APIEndpoints.v1.secure.auth.refreshToken.path))
  }

  getUser() {
    return this.wrapRequest(this.ax.get<GetUserResponse>(APIEndpoints.v1.secure.user.path))
  }

  getUserSystemPreference(errorHandlers: APIErrorHandlers<[404]>) {
    return this.wrapRequest(
      this.ax.get<GetUserSystemPreferenceResponse>(APIEndpoints.v1.secure.systemPreferences.path),
      errorHandlers,
    )
  }

  createDefaultUserSystemPreference() {
    return this.wrapRequest(
      this.ax.post<GetUserSystemPreferenceResponse>(APIEndpoints.v1.secure.systemPreferences.path),
    )
  }

  updateUserSystemPreference(data: UpdateUserSystemPreferenceRequest) {
    return this.wrapRequest(
      this.ax.put<GetUserSystemPreferenceResponse>(
        APIEndpoints.v1.secure.systemPreferences.path,
        data,
      ),
    )
  }

  getCashbunnyUserPreference(errorHandlers: APIErrorHandlers<[404]>) {
    return this.wrapRequest(
      this.ax.get<GetCashbunnyUserPreferenceResponse>(
        APIEndpoints.v1.secure.cashbunny.userPreferences.path,
      ),
      errorHandlers,
    )
  }

  createCashbunnyDefaultUserPreference() {
    return this.wrapRequest(
      this.ax.post<GetCashbunnyUserPreferenceResponse>(
        APIEndpoints.v1.secure.cashbunny.userPreferences.path,
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

  getCashbunnyOverview(dateRange?: { from: Dayjs; to: Dayjs }) {
    return this.wrapRequest(
      this.ax.get<GetCashbunnyOverviewResponse>(APIEndpoints.v1.secure.cashbunny.overview.path, {
        params: dateRange
          ? {
              from: dateRange.from.unix(),
              to: dateRange.to.unix(),
            }
          : undefined,
      }),
    )
  }

  getCashbunnyAccounts() {
    return this.wrapRequest(
      this.ax.get<GetCashbunnyAccountsResponse>(APIEndpoints.v1.secure.cashbunny.accounts.path),
    )
  }

  createCashbunnyAccount(
    data: CreateCashbunnyAccountRequest,
    errorHandlers: APIErrorHandlers<[400]>,
  ) {
    return this.wrapRequest(
      this.ax.post(APIEndpoints.v1.secure.cashbunny.accounts.path, data),
      errorHandlers,
    )
  }

  deleteCashbunnyAccount(accountId: number) {
    return this.wrapRequest(
      this.ax.delete(APIEndpoints.v1.secure.cashbunny.accounts.path + `/${accountId}`),
    )
  }

  updateCashbunnyAccount(
    accountId: number,
    data: UpdateCashbunnyAccountRequest,
    errorHandlers: APIErrorHandlers<[400]>,
  ) {
    return this.wrapRequest(
      this.ax.put(APIEndpoints.v1.secure.cashbunny.accounts.path + `/${accountId}`, data),
      errorHandlers,
    )
  }

  getCashbunnyTransactions() {
    return this.wrapRequest(
      this.ax.get<GetCashbunnyTransactionsResponse>(
        APIEndpoints.v1.secure.cashbunny.transactions.path,
      ),
    )
  }

  createCashbunnyTransaction(
    data: CreateCashbunnyTransactionRequest,
    errorHandlers: APIErrorHandlers<[400]>,
  ) {
    return this.wrapRequest(
      this.ax.post(APIEndpoints.v1.secure.cashbunny.transactions.path, data),
      errorHandlers,
    )
  }

  deleteCashbunnyTransaction(transactionId: number) {
    return this.wrapRequest(
      this.ax.delete(APIEndpoints.v1.secure.cashbunny.transactions.path + `/${transactionId}`),
    )
  }

  updateCashbunnyTransaction(
    transactionId: number,
    data: UpdateCashbunnyTransactionRequest,
    errorHandlers: APIErrorHandlers<[400]>,
  ) {
    return this.wrapRequest(
      this.ax.put(APIEndpoints.v1.secure.cashbunny.transactions.path + `/${transactionId}`, data),
      errorHandlers,
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
        const errorData = nestedSnakeToCamel(error.response?.data)
        if ('message' in errorData && 'validationMessages' in errorData) {
          const apiErr = new APIError(error.status, errorData.message, errorData.validationMessages)
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

  ax.interceptors.request.use((req) => {
    // If we have any data we are sending, convert it to snake_case.
    // Even if req.data is a number or string, nestedCamelToSnake will handle that and return as is
    // so all we need is "if (req.data)" to check for data.
    if (req.data) {
      try {
        req.data = nestedCamelToSnake(req.data)
      } catch {
        // Do nothing
      }
    }
    return req
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
          await ax.post(APIEndpoints.v1.secure.auth.refreshToken.path, null, {
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
