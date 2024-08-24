export abstract class Size {
  w: number
  h: number

  constructor(w: number, h: number) {
    this.w = w
    this.h = h
  }

  clone<T extends Size>(this: T): T {
    return new (this.constructor as new (w: number, h: number) => T)(this.w, this.h)
  }
}
