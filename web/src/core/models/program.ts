export class Program {
  name: string
  icon: string
  component: Object
  componentProps: Object

  constructor(name: string, icon: string, component: Object, componentProps: Object) {
    this.name = name
    this.icon = icon
    this.component = component
    this.componentProps = componentProps
  }
}
