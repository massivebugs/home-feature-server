export abstract class Position {
  x: number
  y: number

  constructor(x: number, y: number) {
    this.x = x
    this.y = y
  }

  clone<T extends Position>(this: T): T {
    return new (this.constructor as new (x: number, y: number) => T)(this.x, this.y)
  }
}
