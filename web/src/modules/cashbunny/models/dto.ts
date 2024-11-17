import type { CashbunnyRecurrenceRuleResponse } from '@/core/composables/useAPI'

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
  recurrence_rule: CashbunnyRecurrenceRuleResponse
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
  recurrence_rule: CashbunnyRecurrenceRuleResponse
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
  recurrence_rule: CashbunnyRecurrenceRuleResponse
  transaction_category: TransactionCategoryDto | null
}

export type PlannerParametersDto = {
  assets: PlannerAssetDto[]
  revenues: PlannerRevenueDto[]
  liabilities: PlannerLiabilityDto[]
  transaction_categories: TransactionCategoryDto[]
}

export type TransactionCategoryDto = {
  id: number
  name: string
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
