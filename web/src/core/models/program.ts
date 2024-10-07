export class Program {
  id: string
  icon: string
  component: Object
  componentProps: Object

  constructor(id: string, icon: string, component: Object, componentProps: Object) {
    this.id = id
    this.icon = icon
    this.component = component
    this.componentProps = componentProps
  }
}
