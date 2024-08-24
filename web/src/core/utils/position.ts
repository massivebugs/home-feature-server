import { AbsolutePosition } from '../models/absolutePosition'

export const ClientArea = {
  top: 'top',
  topLeft: 'topLeft',
  topRight: 'topRight',
  center: 'center',
  centerLeft: 'centerLeft',
  centerRight: 'centerRight',
  bottom: 'bottom',
  bottomLeft: 'bottomLeft',
  bottomRight: 'bottomRight',
} as const

export type ClientArea = (typeof ClientArea)[keyof typeof ClientArea]

export const getClientPosition = (area: ClientArea): AbsolutePosition => {
  switch (area) {
    case 'top':
      return new AbsolutePosition(0, 0)
    case 'topLeft':
      return new AbsolutePosition(0, 0)
    case 'topRight':
      return new AbsolutePosition(0, 0)
    case 'center':
      return new AbsolutePosition(0, 0)
    case 'centerLeft':
      return new AbsolutePosition(0, 0)
    case 'centerRight':
      return new AbsolutePosition(0, 0)
    case 'bottom':
      return new AbsolutePosition(0, 0)
    case 'bottomLeft':
      return new AbsolutePosition(0, 0)
    case 'bottomRight':
      return new AbsolutePosition(0, 0)
  }
  return new AbsolutePosition(0, 0)
}
