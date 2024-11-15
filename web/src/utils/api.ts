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
        const refreshRes = await api.put(APIEndpoints.v1.auth.token, null, {
          validateStatus: null,
        })

        if (refreshRes.status === 401) {
          throw new Error('Refresh token is either invalid or expired')
        }

        // Retrieve a new refresh token
        // TODO: Error handling when creating refresh token fails
        await api.post(APIEndpoints.v1.secure.auth.token, null, {
          validateStatus: null,
        })

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
    auth: {
      default: 'v1/auth',
      token: 'v1/auth/token',
    },
    secure: {
      auth: {
        token: 'v1/secure/auth/token',
      },
      user: {
        default: 'v1/secure/user',
      },
      systemPreferences: {
        default: 'v1/secure/system_preferences',
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
