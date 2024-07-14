import type { BalanceDto } from './dto'

export class Balance {
  id: number
  name: string
  amount: number
  currency: string
  description: string

  constructor(dto: BalanceDto) {
    this.id = dto.id
    this.name = dto.name
    this.amount = dto.amount
    this.currency = dto.currency
    this.description = dto.description
  }
}
