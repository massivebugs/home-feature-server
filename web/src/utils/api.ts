import axios, { AxiosError } from 'axios'

export const isAPIError = (e: any) => {
  return axios.isAxiosError(e)
}

const api = axios.create({
  baseURL: 'http://localhost:1323/api/',
  timeout: 10000,
  headers: {},
})

api.interceptors.request.use(
  (config) => {
    config.headers['Authorization'] = `Bearer ${localStorage.getItem('token')}`
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

api.interceptors.response.use(
  (res) => {
    return res
  },
  (err: AxiosError) => {
    if (err.response?.status === 401 && location.pathname !== '/login') {
      window.location.replace('/login')
    }
    return Promise.reject(err)
  },
)

export const APIEndpoints = {
  v1: {
    auth: 'v1/auth',
    authToken: 'v1/auth/token',
    secure: {
      authUser: 'v1/secure/auth',
      cashbunny: {
        overview: 'v1/secure/cashbunny/overview',
        currencies: 'v1/secure/cashbunny/currencies',
        userPreferences: 'v1/secure/cashbunny/user_preferences',
        accounts: 'v1/secure/cashbunny/accounts',
        transactions: 'v1/secure/cashbunny/transactions',
      },
    },
  },
}

export { api }
