import type { APIResponse } from '@/core/models/dto'

export type AccountCategoryDto = {
  id: number
  name: string
  description: string
  created_at: string
  updated_at: string
}
export type GetCategoriesResponse = APIResponse<AccountCategoryDto[]>

export type AccountDto = {
  id: number
  name: string
  description: string
  balance: string
  currency: string
  created_at: string
  updated_at: string
  category: AccountCategoryDto
}
export type GetAccountsResponse = APIResponse<AccountDto[]>

export type CreateAccountDto = {
  name: string
  category_name: string
  description: string
  balance: number
  currency: string
  type: string
  order_index: number
}
