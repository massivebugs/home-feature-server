import type { FrequencyStr } from './recurrence_rule'

export type OverviewDto = {
  net_worth: { [key: string]: string }
  profit_loss_summary: { [key: string]: { revenue: string; expense: string; profit: string } }
  asset_accounts: AccountDto[]
  liability_accounts: AccountDto[]
  transactions: TransactionDto[]
  transactions_from_scheduled: TransactionDto[]
}

export type PlannerAssetDto = {
  asset_account_id: string
  name: string
  description: string
  amount: number
  currency: string
}

export type PlannerRevenueDto = {
  scheduled_transaction_id: string
  description: string
  amount: number
  currency: string
  source_revenue_account_id: string
  source_revenue_account_name: string
  destination_asset_account_id: string
  destination_asset_account_name: string
  recurrence_rule: RecurrenceRuleDto
  transaction_category: TransactionCategoryDto | null
}

export type PlannerLiabilityDto = {
  scheduled_transaction_id: string
  description: string
  amount: number
  currency: string
  source_asset_account_id: string
  source_asset_account_name: string
  destination_liability_account_id: string
  destination_liability_account_name: string
  recurrence_rule: RecurrenceRuleDto
  transaction_category: TransactionCategoryDto | null
}

export type PlannerExpenseDto = {
  scheduled_transaction_id: string
  description: string
  amount: number
  currency: string
  source_asset_account_id: string
  source_asset_account_name: string
  destination_expense_account_id: string
  destination_expense_account_name: string
  recurrence_rule: RecurrenceRuleDto
  transaction_category: TransactionCategoryDto | null
}

export type PlannerParametersDto = {
  assets: PlannerAssetDto[]
  revenues: PlannerRevenueDto[]
  liabilities: PlannerLiabilityDto[]
  transaction_categories: TransactionCategoryDto[]
}

export type AccountDto = {
  id: number
  category: string
  name: string
  description: string
  amount: number
  amount_display: string
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
  order_index?: number
}

export type UpdateAccountDto = {
  name?: string
  description?: string
  order_index?: number
}

export type TransactionCategoryDto = {
  id: number
  name: string
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
  freq: FrequencyStr
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
