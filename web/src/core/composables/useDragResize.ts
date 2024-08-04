import { ref, onMounted, onUnmounted, type Ref, type CSSProperties } from 'vue'
import type { RelativePosition } from '../models/relative_position'
import type { RelativeSize } from '../models/relative_size'

const ResizeCursor = {
  ew: 'ew-resize',
  ns: 'ns-resize',
  nesw: 'nesw-resize',
  nwse: 'nwse-resize',
} as const

export type ResizeCursor = (typeof ResizeCursor)[keyof typeof ResizeCursor]

const DRAG_CORNER_OFFSET_PX = 5

export function useDraggableResizable(
  initialPos: RelativePosition,
  initialSize: RelativeSize,
  el?: Ref<HTMLElement | undefined>,
) {
  // Drag states
  let isDragging = false
  let isResizing = false
  let startMouseX = 0
  let startMouseY = 0
  let startWidth = 0
  let startHeight = 0
  let startTop = 0
  let startLeft = 0
  const dragCorners = {
    top: false,
    bottom: false,
    left: false,
    right: false,
  }

  // Result
  const boxWidth = ref(initialSize.width) // percentage
  const boxHeight = ref(initialSize.height) // percentage
  const boxLeft = ref(initialPos.x) // percentage
  const boxTop = ref(initialPos.y) // percentage
  const dragStyle = ref<CSSProperties>({
    cursor: 'auto',
    userSelect: 'auto',
  })

  const onDragStart = (e: MouseEvent | TouchEvent) => {
    if (e.type === 'mousedown') {
      e = e as MouseEvent
      isDragging = true
      startMouseX = e.clientX
      startMouseY = e.clientY
      startTop = boxTop.value
      startLeft = boxLeft.value
      window.addEventListener('mousemove', onMouseMove)
      window.addEventListener('mouseup', onMouseUp)
    } else {
      e = e as TouchEvent
      isDragging = true
      startMouseX = e.touches[0].clientX
      startMouseY = e.touches[0].clientY
      startTop = boxTop.value
      startLeft = boxLeft.value
      window.addEventListener('touchmove', onTouchMove)
      window.addEventListener('touchend', onTouchEnd)
    }
  }

  const onResizeStart = (e: MouseEvent | TouchEvent) => {
    if (e.type === 'mousedown') {
      e = e as MouseEvent
      const target = el?.value ? el.value : (e.target as HTMLElement)
      // We need to use getBoundingClientRect() for pixel values
      const viewportRect = target.getBoundingClientRect()
      dragCorners.top, dragCorners.bottom, dragCorners.left, (dragCorners.right = false)

      // Check which corners are being selected at the moment
      dragCorners.top = viewportRect.top > e.clientY - DRAG_CORNER_OFFSET_PX
      dragCorners.bottom =
        viewportRect.top + viewportRect.height < e.clientY + DRAG_CORNER_OFFSET_PX
      dragCorners.left = viewportRect.left > e.clientX - DRAG_CORNER_OFFSET_PX
      dragCorners.right = viewportRect.left + viewportRect.width < e.clientX + DRAG_CORNER_OFFSET_PX

      // If no corners are being selected, ignore and return
      if (!dragCorners.top && !dragCorners.bottom && !dragCorners.left && !dragCorners.right) {
        return
      }

      // Set which cursor should be displayed, based on the selected corners
      if ((dragCorners.top && dragCorners.left) || (dragCorners.bottom && dragCorners.right)) {
        dragStyle.value.cursor = ResizeCursor.nwse
      } else if (
        (dragCorners.top && dragCorners.right) ||
        (dragCorners.bottom && dragCorners.left)
      ) {
        dragStyle.value.cursor = ResizeCursor.nesw
      } else if (dragCorners.top || dragCorners.bottom) {
        dragStyle.value.cursor = ResizeCursor.ns
      } else {
        dragStyle.value.cursor = ResizeCursor.ew
      }

      isResizing = true
      isDragging = false // Resizing overrides dragging behaviors
      startMouseX = e.clientX
      startMouseY = e.clientY
      startTop = boxTop.value
      startLeft = boxLeft.value
      startWidth = boxWidth.value
      startHeight = boxHeight.value
      window.addEventListener('mousemove', onMouseMove)
      window.addEventListener('mouseup', onMouseUp)
    } else {
      e = e as TouchEvent
      const target = el?.value ? el.value : (e.target as HTMLElement)
      // We need to use getBoundingClientRect() for pixel values
      const viewportRect = target.getBoundingClientRect()
      dragCorners.top, dragCorners.bottom, dragCorners.left, (dragCorners.right = false)

      // Check which corners are being selected at the moment
      dragCorners.top = viewportRect.top > e.touches[0].clientY - DRAG_CORNER_OFFSET_PX
      dragCorners.bottom =
        viewportRect.top + viewportRect.height < e.touches[0].clientY + DRAG_CORNER_OFFSET_PX
      dragCorners.left = viewportRect.left > e.touches[0].clientX - DRAG_CORNER_OFFSET_PX
      dragCorners.right =
        viewportRect.left + viewportRect.width < e.touches[0].clientX + DRAG_CORNER_OFFSET_PX

      // If no corners are being selected, ignore and return
      if (!dragCorners.top && !dragCorners.bottom && !dragCorners.left && !dragCorners.right) {
        return
      }

      // Set which cursor should be displayed, based on the selected corners
      if ((dragCorners.top && dragCorners.left) || (dragCorners.bottom && dragCorners.right)) {
        dragStyle.value.cursor = ResizeCursor.nwse
      } else if (
        (dragCorners.top && dragCorners.right) ||
        (dragCorners.bottom && dragCorners.left)
      ) {
        dragStyle.value.cursor = ResizeCursor.nesw
      } else if (dragCorners.top || dragCorners.bottom) {
        dragStyle.value.cursor = ResizeCursor.ns
      } else {
        dragStyle.value.cursor = ResizeCursor.ew
      }

      isResizing = true
      isDragging = false // Resizing overrides dragging behaviors
      startMouseX = e.touches[0].clientX
      startMouseY = e.touches[0].clientY
      startTop = boxTop.value
      startLeft = boxLeft.value
      startWidth = boxWidth.value
      startHeight = boxHeight.value
      window.addEventListener('touchmove', onTouchMove)
      window.addEventListener('touchend', onTouchEnd)
    }
    dragStyle.value.userSelect = 'none'
  }

  const onMouseMove = (e: MouseEvent) => {
    if (isDragging) {
      const dx = ((e.clientX - startMouseX) / window.innerWidth) * 100
      const dy = ((e.clientY - startMouseY) / window.innerHeight) * 100
      boxTop.value = startTop + dy
      boxLeft.value = startLeft + dx
    } else if (isResizing) {
      const dx = ((e.clientX - startMouseX) / window.innerWidth) * 100
      const dy = ((e.clientY - startMouseY) / window.innerHeight) * 100
      if (dragCorners.top) {
        boxTop.value = startTop + dy
        boxHeight.value = startHeight - dy
      }
      if (dragCorners.bottom) {
        boxHeight.value = startHeight + dy
      }
      if (dragCorners.left) {
        boxLeft.value = startLeft + dx
        boxWidth.value = startWidth - dx
      }
      if (dragCorners.right) {
        boxWidth.value = startWidth + dx
      }
    }
  }

  const onTouchMove = (e: TouchEvent) => {
    if (isDragging) {
      const dx = ((e.touches[0].clientX - startMouseX) / window.innerWidth) * 100
      const dy = ((e.touches[0].clientY - startMouseY) / window.innerHeight) * 100
      boxTop.value = startTop + dy
      boxLeft.value = startLeft + dx
    } else if (isResizing) {
      const dx = ((e.touches[0].clientX - startMouseX) / window.innerWidth) * 100
      const dy = ((e.touches[0].clientY - startMouseY) / window.innerHeight) * 100
      if (dragCorners.top) {
        boxTop.value = startTop + dy
        boxHeight.value = startHeight - dy
      }
      if (dragCorners.bottom) {
        boxHeight.value = startHeight + dy
      }
      if (dragCorners.left) {
        boxLeft.value = startLeft + dx
        boxWidth.value = startWidth - dx
      }
      if (dragCorners.right) {
        boxWidth.value = startWidth + dx
      }
    }
  }

  const onMouseUp = () => {
    isDragging = false
    isResizing = false
    dragStyle.value.cursor = 'auto'
    dragStyle.value.userSelect = 'auto'
    window.removeEventListener('mousemove', onMouseMove)
    window.removeEventListener('mouseup', onMouseUp)
  }

  const onTouchEnd = () => {
    isDragging = false
    isResizing = false
    dragStyle.value.cursor = 'auto'
    dragStyle.value.userSelect = 'auto'
    window.removeEventListener('touchmove', onTouchMove)
    window.removeEventListener('touchend', onTouchEnd)
  }

  onMounted(() => {
    window.addEventListener('mouseup', onMouseUp)
  })

  onUnmounted(() => {
    window.removeEventListener('mouseup', onMouseUp)
  })

  return {
    boxWidth,
    boxHeight,
    boxTop,
    boxLeft,
    dragStyle,
    onDragStart,
    onResizeStart,
  }
}
