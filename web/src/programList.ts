import { SYSTEM_SETTINGS_PROGRAM_ID } from './core/constants'
import { Program } from './core/models/program'
import SystemSettingsView from './core/views/SystemSettingsView.vue'
import { CASHBUNNY_PROGRAM_ID } from './modules/cashbunny/constants'
import CashbunnyView from './modules/cashbunny/views/CashbunnyView.vue'
import { PORTFOLIO_PROGRAM_ID } from './modules/portfolio/constants'
import PortfolioView from './modules/portfolio/views/PortfolioView.vue'
import { SANDMAN_PROGRAM_ID } from './modules/sandman/constants'
import SandmanView from './modules/sandman/views/SandmanView.vue'

// Returns all of the programs to be registered to the system,
// and options for window display
export const programList: Program[] = [
  new Program(CASHBUNNY_PROGRAM_ID, '/images/cashbunny_icon_round.svg', CashbunnyView, {}),
  new Program(PORTFOLIO_PROGRAM_ID, '/images/portfolio_icon_normal.png', PortfolioView, {}),
  new Program(SANDMAN_PROGRAM_ID, '/images/sandman_icon.jpeg', SandmanView, {}),
  new Program(
    SYSTEM_SETTINGS_PROGRAM_ID,
    '/images/system_settings_icon.svg',
    SystemSettingsView,
    {},
  ),
]
