import { capitalize } from 'lodash'
import { Frequency, RRule } from 'rrule'
import type { RecurrenceRuleDto } from './dto'

export type FrequencyStr = keyof typeof Frequency
export const FrequencyStrs: { [key in FrequencyStr]: FrequencyStr } = {
  YEARLY: 'YEARLY',
  MONTHLY: 'MONTHLY',
  WEEKLY: 'WEEKLY',
  DAILY: 'DAILY',
  HOURLY: 'HOURLY',
  MINUTELY: 'MINUTELY',
  SECONDLY: 'SECONDLY',
}

export class RecurrenceRule {
  rule: RRule
  constructor(data: RecurrenceRuleDto) {
    this.rule = new RRule({
      freq: RecurrenceRule.frequencyStrToRRuleEnum(data.freq),
      dtstart: new Date(data.dtstart),
      count: data.count,
      interval: data.interval,
      until: new Date(data.until),
    })
  }

  toHumanFriendlyString(): string {
    return capitalize(this.rule.toText())
  }

  static frequencyStrToRRuleEnum(freq: FrequencyStr): Frequency {
    switch (freq) {
      case 'YEARLY':
        return Frequency.YEARLY
      case 'MONTHLY':
        return Frequency.MONTHLY
      case 'WEEKLY':
        return Frequency.WEEKLY
      case 'DAILY':
        return Frequency.DAILY
      case 'HOURLY':
        return Frequency.HOURLY
      case 'MINUTELY':
        return Frequency.MINUTELY
      default:
        return Frequency.SECONDLY
    }
  }
}
