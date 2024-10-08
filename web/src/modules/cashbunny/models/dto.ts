export type OverviewDto = {
  net_worth: { [key: string]: string }
  summaries: { [key: string]: { revenue: string; expense: string; profit: string } }
  transactions: TransactionDto[]
  transactions_from_scheduled: TransactionDto[]
}

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
  balance: number
  currency: string
  type: string
  created_at: string
  updated_at: string
}

export type CreateAccountDto = {
  name: string
  category: string
  description: string
  currency: string
  order_index: number
}

export type UpdateAccountDto = {
  name?: string
  description?: string
  order_index?: number
}

export type TransactionDto = {
  id: number
  description: string
  amount: number
  currency: string
  amount_display: string
  transacted_at: string
  created_at: string
  updated_at: string
  source_account_id: number
  source_account_name: string
  destination_account_id: number
  destination_account_name: string
  scheduled_transaction: ScheduledTransactionDto | null
}

export type ScheduledTransactionDto = {
  id: number
  description: string
  amount: number
  currency: string
  amount_display: string
  created_at: string
  updated_at: string
  recurrence_rule: RecurrenceRuleDto
  source_account_id: number
  source_account_name: string
  destination_account_id: number
  destination_account_name: string
}

export type RecurrenceRuleDto = {
  freq: string
  dtstart: string
  count: number
  interval: number
  until: string
}

export type CreateTransactionDto = {
  description: string
  amount: number
  currency: string
  source_account_id: number
  destination_account_id: number
  transacted_at: string
}

export type UpdateTransactionDto = {
  description?: string
  amount?: number
  transacted_at?: string
}
