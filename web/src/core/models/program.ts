// A Program represents each application listed on the desktop interface.
// It contains a reference to it's Vue component, which is then dynamically rendered
// during runtime.
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
