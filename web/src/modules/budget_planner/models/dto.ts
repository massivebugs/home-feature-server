export type BalanceDto = {
  id: number
  name: string
  amount: number
  currency: string
  description: string
}

export type AccountingSummaryDto = {
  balances: BalanceDto[]
}
