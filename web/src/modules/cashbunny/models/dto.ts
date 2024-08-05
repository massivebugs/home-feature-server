import type { APIResponse } from '@/core/models/dto'

export type CategoryDto = {
  id: number
  name: string
  description: string
  created_at: string
  updated_at: string
}

export type AccountDto = {
  id: number
  name: string
  description: string
  balance: string
  currency: string
  created_at: string
  updated_at: string
  category: CategoryDto
}
export type GetAccountsResponse = APIResponse<AccountDto[]>
