import type { APIResponse } from '@/core/models/dto'

export type AccountDto = {
  id: number
  category: string
  name: string
  description: string
  balance: string
  currency: string
  type: string
  created_at: string
  updated_at: string
}
export type GetAccountsResponse = APIResponse<AccountDto[]>

export type CreateAccountDto = {
  name: string
  category: string
  description: string
  balance: number
  currency: string
  order_index: number
}
