import { AbsolutePosition } from '../models/absolute_position'
import { AbsoluteSize } from '../models/absolute_size'
import { RelativePosition } from '../models/relative_position'
import { RelativeSize } from '../models/relative_size'

export const getRelativeClientPosition = (absPos: AbsolutePosition) => {
  const relPos = new RelativePosition(0, 0)
  relPos.x = (absPos.x / document.body.clientWidth) * 100
  relPos.y = (absPos.y / document.body.clientHeight) * 100

  return relPos
}

export const getRelativeParentPosition = (absPos: AbsolutePosition, parentEl: HTMLElement) => {
  const parentRect = parentEl.getBoundingClientRect()

  const relPos = new RelativePosition(0, 0)

  if (parentRect) {
    relPos.x = ((absPos.x - parentRect.left) / parentRect.width) * 100
    relPos.y = ((absPos.y - parentRect.top) / parentRect.height) * 100
  }

  return relPos
}

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

export const getMouseMetrics = (clientX: number, clientY: number, parentEl: HTMLElement) => {
  const parentRect = parentEl.getBoundingClientRect()

  const absPos = new AbsolutePosition(clientX + window.scrollX, clientY + window.scrollY)
  const absSize = new AbsoluteSize(clientX, clientY)

  const relPos = new RelativePosition(0, 0)
  const relSize = new RelativeSize(0, 0)
  if (parentRect) {
    relPos.x = ((clientX - parentRect.left) / parentRect.width) * 100
    relPos.y = ((clientY - parentRect.top) / parentRect.height) * 100
    relSize.width = (clientX / parentRect.width) * 100
    relSize.height = (clientY / parentRect.height) * 100
  }

  return {
    absPos,
    absSize,
    relPos,
    relSize,
  }
}
