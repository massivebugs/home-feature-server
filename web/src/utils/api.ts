import axios, { AxiosError, type AxiosRequestConfig } from 'axios'
import { snakeCase } from 'lodash'

export const isAPIError = (e: any) => {
  return axios.isAxiosError(e)
}

const api = axios.create({
  baseURL: `https://${window.location.hostname}:${import.meta.env.VITE_PORT}/api/`,
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
        const refreshRes = await api.put(APIEndpoints.v1.auth.token.path, null, {
          validateStatus: null,
        })

        if (refreshRes.status === 401) {
          throw new Error('Refresh token is either invalid or expired')
        }

        // Retrieve a new refresh token
        // TODO: Error handling when creating refresh token fails
        await api.post(APIEndpoints.v1.secure.auth.token.path, null, {
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
        return makeEndpointsProxy(value, traversed.concat(snakeCase(prop.toString())))
      }

      return value
    },
  }

  return new Proxy(target, trap) as AugmentedProxy<T>
}

export const APIEndpoints = makeEndpointsProxy(Endpoints)

export { api }
