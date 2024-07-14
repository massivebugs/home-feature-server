import { type Ref } from 'vue'
import { AbsolutePosition } from '../models/absolute_position'
import { RelativePosition } from '../models/relative_position'

export function usePointerDown(el: Ref<HTMLElement | undefined>) {
  const absPosition = new AbsolutePosition(0, 0)
  const relPosition = new RelativePosition(0, 0)

  const updatePosition = (event: MouseEvent | Touch) => {
    if (!el.value) {
      return
    }
    const rect = el.value.getBoundingClientRect()
    absPosition.x = event.clientX
    absPosition.y = event.clientY
    relPosition.x = ((event.clientX - rect.left) / rect.width) * 100
    relPosition.y = ((event.clientY - rect.top) / rect.height) * 100
  }

  const _onPointerDown = (
    e: MouseEvent | TouchEvent,
    onPointerDown: (absPosition: AbsolutePosition, relPosition: RelativePosition) => void,
  ) => {
    if (e.type === 'mousedown') {
      e = e as MouseEvent
      updatePosition(e)
      onPointerDown(absPosition, relPosition)
    } else if (e.type === 'touchstart') {
      e = e as TouchEvent
      updatePosition(e.touches[0])
      onPointerDown(absPosition, relPosition)
    }
  }

  return {
    onPointerDown: _onPointerDown,
  }
}
