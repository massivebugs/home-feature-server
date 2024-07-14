import { onMounted, onUnmounted, provide, ref, type Ref } from 'vue'
import { AbsolutePosition } from '../models/absolute_position'
import { RelativePosition } from '../models/relative_position'
import { usePointerDown } from './usePointerDown'
import type { ContextMenuOptions } from '../components/WindowsXPContextMenuComponent.vue'
import { getElementMetrics } from '../utils/element'
import type { AbsoluteSize } from '../models/absolute_size'
import type { RelativeSize } from '../models/relative_size'

export type ContextOptions = {
  pointerDown?: {
    onPointerDown: (absPosition: AbsolutePosition, relPosition: RelativePosition) => void
  }
}

export type SetContextMenu = (
  newContextMenu: ContextMenuOptions | null,
  position?: RelativePosition,
) => void

export type GetContextualMetrics = (targetEl: HTMLElement) => {
  absPos: AbsolutePosition
  absSize: AbsoluteSize
  relPos: RelativePosition
  relSize: RelativeSize
}

export function useContext(el: Ref<HTMLElement | undefined>, options: ContextOptions) {
  const { onPointerDown } = usePointerDown(el)
  const contextMenu = ref<ContextMenuOptions | null>(null)
  const contextMenuRelativePosition = ref<RelativePosition>(new RelativePosition(0, 0))

  const setContextMenu: SetContextMenu = (
    newContextMenu: ContextMenuOptions | null,
    position?: RelativePosition,
  ) => {
    contextMenu.value = newContextMenu
    if (position) {
      contextMenuRelativePosition.value = position
    }
  }

  const getContextualMetrics: GetContextualMetrics = (targetEl: HTMLElement) => {
    return getElementMetrics(targetEl, el.value as HTMLElement)
  }

  provide('contextMenu', contextMenu)
  provide('setContextMenu', setContextMenu)
  provide('getContextualMetrics', getContextualMetrics)

  onMounted(() => {
    console.debug('Context created', el)
    if (options.pointerDown) {
      const func = options.pointerDown.onPointerDown
      el.value?.addEventListener('mousedown', (e) => onPointerDown(e, func))
      el.value?.addEventListener('touchstart', (e) => onPointerDown(e, func))
    }
    window.addEventListener('click', () => {
      setContextMenu(null)
    })
  })

  onUnmounted(() => {
    console.debug('Context removed', el)
    if (options.pointerDown) {
      const func = options.pointerDown.onPointerDown
      el.value?.removeEventListener('mousedown', (e) => onPointerDown(e, func))
      el.value?.removeEventListener('touchstart', (e) => onPointerDown(e, func))
    }
  })

  return {
    contextMenu,
    contextMenuRelativePosition,
    setContextMenu,
  }
}
