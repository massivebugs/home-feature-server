import { type ComposerTranslation } from 'vue-i18n'
import { Program } from './core/models/program'
import { CASHBUNNY_PROGRAM_ID } from './modules/cashbunny/constants'
import CashbunnyView from './modules/cashbunny/views/CashbunnyView.vue'
import { PORTFOLIO_PROGRAM_ID } from './modules/portfolio/constants'
import PortfolioView from './modules/portfolio/views/PortfolioView.vue'
import { SYSTEM_SETTINGS_PROGRAM_ID } from './modules/system_settings/constants'
import SystemSettingsView from './modules/system_settings/views/SystemSettingsView.vue'

// Returns all of the programs to be registered to the system,
// and options for window display
export const getPrograms = (t: ComposerTranslation): Program[] => [
  new Program(
    CASHBUNNY_PROGRAM_ID,
    t('cashbunny.name'),
    '/images/cashbunny_icon_round.svg',
    CashbunnyView,
    {},
  ),
  new Program(
    PORTFOLIO_PROGRAM_ID,
    t('portfolio.name'),
    '/images/portfolio_icon_normal.png',
    PortfolioView,
    {},
  ),
  new Program(
    SYSTEM_SETTINGS_PROGRAM_ID,
    t('systemSettings.name'),
    '/images/system_settings_icon.svg',
    SystemSettingsView,
    {},
  ),
]
