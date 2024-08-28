import { type App, reactive } from 'vue'

export type MediaQueryBreakpoints = {
  xl: number
  lg: number
  m: number
  sm: number
}

export type MediaQuery = {
  isXL: boolean
  isLG: boolean
  isM: boolean
  isSM: boolean
}

export default {
  install: (app: App, options: { breakpoints: MediaQueryBreakpoints }) => {
    const result = reactive<MediaQuery>({
      isXL: false,
      isLG: false,
      isM: false,
      isSM: false,
    })

    const update = () => {
      result.isSM = window.matchMedia(`(max-width: ${options.breakpoints.m - 1}px)`).matches
      result.isM = window.matchMedia(`(min-width: ${options.breakpoints.lg}px )`).matches
      result.isSM = window.matchMedia(`(min-width: ${options.breakpoints.m}px)`).matches
    }
    // window.addEventListener('resize', () => {
    //   const width = window.innerWidth

    //   if (width < options.breakpoints.m) {
    //     result.deviceSize = DeviceSize.s
    //   } else if (width < options.breakpoints.l) {
    //     result.deviceSize = DeviceSize.m
    //   } else if (width < options.breakpoints.xl) {
    //     result.deviceSize = DeviceSize.l
    //   } else {
    //     result.deviceSize = DeviceSize.xl
    //   }
    // })

    app.config.globalProperties.$mediaQuery = result
    app.provide('$mediaQuery', result)
  },
}

// const mediaQueryPlugin = {
//   install(app: any) {
//     const state = reactive<MediaQueryState>({
//       isMobile: window.matchMedia('(max-width: 600px)').matches,
//       isTablet: window.matchMedia('(min-width: 601px) and (max-width: 1024px)').matches,
//       isDesktop: window.matchMedia('(min-width: 1025px)').matches,
//     });

//     const updateState = () => {
//       state.isMobile = window.matchMedia('(max-width: 600px)').matches;
//       state.isTablet = window.matchMedia('(min-width: 601px) and (max-width: 1024px)').matches;
//       state.isDesktop = window.matchMedia('(min-width: 1025px)').matches;
//     };

//     window.addEventListener('resize', updateState);

//     app.config.globalProperties.$mediaQuery = readonly(state);
//   },
// };

// export default mediaQueryPlugin;
