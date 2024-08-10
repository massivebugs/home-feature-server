export type GetAllCurrenciesDto = {
  currencies_and_grapheme: { [key: string]: string }
}

export type UserPreferencesDto = {
  user_currencies: string[]
}

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

export type CreateAccountDto = {
  name: string
  category: string
  description: string
  balance: number
  currency: string
  order_index: number
}

export type TransactionDto = {
  id: number
  description: string
  amount: string
  currency: string
  transacted_at: string
  created_at: string
  updated_at: string
  source_account_id: number
  source_account_name: string
  destination_account_id: number
  destination_account_name: string
}

export type CreateTransactionDto = {
  description: string
  amount: number
  currency: string
  source_account_id: number
  destination_account_id: number
  transacted_at: string
}
