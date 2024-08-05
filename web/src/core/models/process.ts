import type { Program } from './program'

export class Process {
  id: string
  program: Program

  constructor(id: string, program: Program) {
    this.id = id
    this.program = program
  }
}
