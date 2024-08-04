import type { WindowOptions } from '../components/WindowComponent.vue'

export class Program {
  name: string
  component: Object
  componentProps: Object
  windowOptions: WindowOptions

  constructor(
    name: string,
    component: Object,
    componentProps: Object,
    windowOptions: WindowOptions,
  ) {
    this.name = name
    this.component = component
    this.componentProps = componentProps
    this.windowOptions = windowOptions
  }
}
