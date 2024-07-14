import type { AccountingSummaryDto } from '@/utils/dto'
import { Balance } from './balance'

export class AccountingSummary {
  balances: Balance[]

  constructor(dto: AccountingSummaryDto) {
    this.balances = dto.balances.map((moneyDto) => new Balance(moneyDto))
  }
}
