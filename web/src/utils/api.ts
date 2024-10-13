import axios, { AxiosError, type AxiosRequestConfig } from 'axios'

export const isAPIError = (e: any) => {
  return axios.isAxiosError(e)
}

const api = axios.create({
  baseURL: `https://${window.location.hostname}:1323/api/`,
  timeout: 10000,
  withCredentials: true,
})

api.interceptors.response.use(
  (res) => {
    return res
  },
  async (err: AxiosError) => {
    if (err.response?.status === 401 && location.pathname !== '/login') {
      try {
        // If a token has expired, try refreshing it first
        const refreshRes = await api.post(APIEndpoints.v1.authRefresh, null, {
          validateStatus: null,
        })

        if (refreshRes.status === 401) {
          throw new Error('Refresh token is either invalid or expired')
        }

        // Retry original request
        return api.request(err.config as AxiosRequestConfig)
      } catch (refreshErr) {
        // Refresh token is invalid
        window.location.replace('/login')
      }
    }
    return Promise.reject(err)
  },
)

export const APIEndpoints = {
  v1: {
    auth: 'v1/auth',
    authToken: 'v1/auth/token',
    authRefresh: 'v1/auth/refresh',
    secure: {
      user: {
        default: 'v1/secure/user',
        systemPreferences: 'v1/secure/user/system_preferences',
      },
      cashbunny: {
        overview: 'v1/secure/cashbunny/overview',
        plannerParameters: 'v1/secure/cashbunny/planner/parameters',
        currencies: 'v1/secure/cashbunny/currencies',
        userPreferences: 'v1/secure/cashbunny/user_preferences',
        accounts: 'v1/secure/cashbunny/accounts',
        transactions: 'v1/secure/cashbunny/transactions',
      },
    },
  },
}

export { api }
