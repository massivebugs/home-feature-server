import { type ComposerTranslation } from 'vue-i18n'
import { Program } from './core/models/program'
import type { useCoreStore } from './core/stores'
import { CASHBUNNY_PROGRAM_ID } from './modules/cashbunny/constants'
import CashbunnyView from './modules/cashbunny/views/CashbunnyView.vue'
import { PORTFOLIO_PROGRAM_ID } from './modules/portfolio/constants'
import PortfolioView from './modules/portfolio/views/PortfolioView.vue'
import type { PiniaStore } from './utils/pinia'

// Returns all of the programs to be registered to the system,
// and options for window display
export const getPrograms = (
  coreStore: PiniaStore<typeof useCoreStore>,
  t: ComposerTranslation,
): Program[] => [
  new Program(CASHBUNNY_PROGRAM_ID, t('cashbunny.name'), 'images/file_icon.svg', CashbunnyView, {}),
  new Program(PORTFOLIO_PROGRAM_ID, t('portfolio.name'), 'images/file_icon.svg', PortfolioView, {}),
]
