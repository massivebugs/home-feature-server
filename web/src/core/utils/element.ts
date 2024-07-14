import { AbsolutePosition } from '../models/absolute_position'
import { AbsoluteSize } from '../models/absolute_size'
import { RelativePosition } from '../models/relative_position'
import { RelativeSize } from '../models/relative_size'

export const getElementMetrics = (htmlEl: HTMLElement, parentEl: HTMLElement) => {
  const rect = htmlEl.getBoundingClientRect()
  const parentRect = parentEl.getBoundingClientRect()

  const absPos = new AbsolutePosition(rect.left + window.scrollX, rect.top + window.scrollY)
  const absSize = new AbsoluteSize(rect.width, rect.height)

  const relPos = new RelativePosition(0, 0)
  const relSize = new RelativeSize(0, 0)
  if (parentRect) {
    relPos.x = ((rect.left - parentRect.left) / parentRect.width) * 100
    relPos.y = ((rect.top - parentRect.top) / parentRect.height) * 100
    relSize.width = (rect.width / parentRect.width) * 100
    relSize.height = (rect.height / parentRect.height) * 100
  }

  return {
    absPos,
    absSize,
    relPos,
    relSize,
  }
}
