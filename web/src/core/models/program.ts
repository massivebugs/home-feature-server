export class Program {
  id: string
  name: string
  icon: string
  component: Object
  componentProps: Object

  constructor(id: string, name: string, icon: string, component: Object, componentProps: Object) {
    this.id = id
    this.name = name
    this.icon = icon
    this.component = component
    this.componentProps = componentProps
  }
}
