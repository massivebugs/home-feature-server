import { camelize } from '@/core/utils/object'
import axios from 'axios'
import type { AccountingSummaryDto } from '../modules/budget_planner/models/dto'

const api = axios.create({
  baseURL: 'http://localhost:8888/api/',
  timeout: 1000,
})

api.interceptors.response.use(
  (res) => {
    res.data = camelize(res.data.Data)
    return res
  },
  (err) => {
    return Promise.reject(err)
  },
)

export const APIEndpoints = {
  v1: { accountingSummary: 'v1/accounting' },
}

export type V1AccountingSummaryResponse = AccountingSummaryDto

export { api }
