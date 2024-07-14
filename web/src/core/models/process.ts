import type { WindowOptions } from '../components/WindowsXPWindowComponent.vue'
import type { Program } from './program'

export class Process {
  id: string
  program: Program
  windowOptions: WindowOptions

  constructor(id: string, program: Program, windowOptions: WindowOptions) {
    this.id = id
    this.program = program
    this.windowOptions = windowOptions
  }
}
