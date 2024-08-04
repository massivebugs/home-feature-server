import { type ComposerTranslation } from 'vue-i18n'
import { Program } from './core/models/program'
import type { useStore } from './core/stores'
import CashbunnyView from './modules/cashbunny/views/CashbunnyView.vue'
import PortfolioView from './modules/portfolio/views/PortfolioView.vue'
import type { PiniaStore } from './utils/pinia'

// Returns all of the programs to be registered to the system,
// and options for window display
export const getPrograms = (
  coreStore: PiniaStore<typeof useStore>,
  t: ComposerTranslation,
): Program[] => [
  new Program(t('cashbunny.name'), 'images/file_icon.svg', CashbunnyView, {}),
  new Program(t('portfolio.name'), 'images/file_icon.svg', PortfolioView, {}),
]
