import { type App, reactive } from 'vue'

export type MediaQueryBreakpoints = {
  md: number
  lg: number
  xl: number
}

export type MediaQuery = {
  isSM: boolean
  isXL: boolean
  isLG: boolean
  isMD: boolean
}

export const mediaQueryPlugin = {
  install: (app: App, options: { breakpoints: MediaQueryBreakpoints }) => {
    const result = reactive<MediaQuery>({
      isSM: false,
      isMD: false,
      isLG: false,
      isXL: false,
    })

    const isSM = window.matchMedia(`(max-width: ${options.breakpoints.md - 1}px)`)
    const isMD = window.matchMedia(
      `(min-width: ${options.breakpoints.md}) and (max-width: ${options.breakpoints.lg - 1}px )`,
    )
    const isLG = window.matchMedia(
      `(min-width: ${options.breakpoints.lg}) and (max-width: ${options.breakpoints.xl - 1}px)`,
    )
    const isXL = window.matchMedia(`(min-width: ${options.breakpoints.xl}px)`)

    isSM.addEventListener('change', (e) => {
      result.isSM = e.matches
    })

    isMD.addEventListener('change', (e) => {
      result.isMD = e.matches
    })

    isLG.addEventListener('change', (e) => {
      result.isLG = e.matches
    })

    isXL.addEventListener('change', (e) => {
      result.isXL = e.matches
    })

    app.config.globalProperties.$mediaQuery = result
    app.provide('mediaQuery', result)
  },
}
