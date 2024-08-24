import { AbsolutePosition } from '../models/absolutePosition'
import { AbsoluteSize } from '../models/absoluteSize'
import { RelativePosition } from '../models/relativePosition'
import { RelativeSize } from '../models/relativeSize'

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
    relSize.w = (rect.width / parentRect.width) * 100
    relSize.h = (rect.height / parentRect.height) * 100
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
    relSize.w = (clientX / parentRect.width) * 100
    relSize.h = (clientY / parentRect.height) * 100
  }

  return {
    absPos,
    absSize,
    relPos,
    relSize,
  }
}
