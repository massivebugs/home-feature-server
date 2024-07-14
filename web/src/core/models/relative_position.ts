export class RelativePosition {
  x: number
  y: number
  constructor(x: number, y: number) {
    this.x = x
    this.y = y
  }

  add(otherX: number, otherY: number) {
    return new RelativePosition(this.x + otherX, this.y + otherY)
  }
}
