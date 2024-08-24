import { type CSSProperties, type Ref, onMounted, onUnmounted, ref } from 'vue'
import { AbsoluteSize } from '../models/absoluteSize'
import type { RelativePosition } from '../models/relativePosition'
import { RelativeSize } from '../models/relativeSize'

const ResizeCursor = {
  ew: 'ew-resize',
  ns: 'ns-resize',
  nesw: 'nesw-resize',
  nwse: 'nwse-resize',
} as const
export type ResizeCursor = (typeof ResizeCursor)[keyof typeof ResizeCursor]

export enum ResizeDirection {
  Top = 1 << 0,
  Bottom = 1 << 1,
  Left = 1 << 2,
  Right = 1 << 3,
  All = ~(~0 << 4),
}

export type DragResizeOptions = {
  resize?: {
    direction?: ResizeDirection
  }
}

const DRAG_CORNER_OFFSET_PX = 5

export function useDragResize(
  initialPos: RelativePosition,
  initialSize: RelativeSize,
  el?: Ref<HTMLElement | undefined>,
  parentEl?: Ref<HTMLElement | undefined>,
  options: DragResizeOptions = {
    resize: {
      direction: ResizeDirection.All,
    },
  },
  onResizeEnd?: () => void,
) {
  // Drag states
  let isDragging = false
  let isResizing = false
  let startMouseX = 0
  let startMouseY = 0
  let startSize = initialSize.clone()
  let startPos = initialPos.clone()
  const dragCorners = {
    top: false,
    bottom: false,
    left: false,
    right: false,
  }
  const originalBoxDimensions = ref({
    size: initialSize.clone(),
    pos: initialPos.clone(),
  })
  const isMaximized = ref<boolean>(false)

  // Result
  const currentSize = ref(initialSize.clone())
  const currentPos = ref(initialPos.clone())
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
      startPos = currentPos.value.clone()
      window.addEventListener('mousemove', onMouseMove)
      window.addEventListener('mouseup', onMouseUp)
    } else {
      e = e as TouchEvent
      isDragging = true
      startMouseX = e.touches[0].clientX
      startMouseY = e.touches[0].clientY
      startPos = currentPos.value.clone()
      window.addEventListener('touchmove', onMouseMove)
      window.addEventListener('touchend', onMouseUp)
    }
  }

  const onResizeStart = (e: MouseEvent | TouchEvent) => {
    const target = el?.value ? el.value : (e.target as HTMLElement)
    let clientX: number
    let clientY: number
    if (e.type === 'mousedown') {
      e = e as MouseEvent
      clientX = e.clientX
      clientY = e.clientY
    } else {
      e = e as TouchEvent
      clientX = e.touches[0].clientX
      clientY = e.touches[0].clientY
    }

    // We need to use getBoundingClientRect() for pixel values
    const viewportRect = target.getBoundingClientRect()
    dragCorners.top, dragCorners.bottom, dragCorners.left, (dragCorners.right = false)

    // Check which corners are being selected at the moment
    dragCorners.top = viewportRect.top > clientY - DRAG_CORNER_OFFSET_PX
    dragCorners.bottom = viewportRect.top + viewportRect.height < clientY + DRAG_CORNER_OFFSET_PX
    dragCorners.left = viewportRect.left > clientX - DRAG_CORNER_OFFSET_PX
    dragCorners.right = viewportRect.left + viewportRect.width < clientX + DRAG_CORNER_OFFSET_PX

    // If no corners are being selected, ignore and return
    if (!dragCorners.top && !dragCorners.bottom && !dragCorners.left && !dragCorners.right) {
      return
    }

    if (
      options.resize?.direction &&
      !(
        (dragCorners.top && !!(options.resize.direction & ResizeDirection.Top)) ||
        (dragCorners.bottom && !!(options.resize.direction & ResizeDirection.Bottom)) ||
        (dragCorners.left && !!(options.resize.direction & ResizeDirection.Left)) ||
        (dragCorners.right && !!(options.resize.direction & ResizeDirection.Right))
      )
    ) {
      return
    }

    // Set which cursor should be displayed, based on the selected corners
    if ((dragCorners.top && dragCorners.left) || (dragCorners.bottom && dragCorners.right)) {
      dragStyle.value.cursor = ResizeCursor.nwse
    } else if ((dragCorners.top && dragCorners.right) || (dragCorners.bottom && dragCorners.left)) {
      dragStyle.value.cursor = ResizeCursor.nesw
    } else if (dragCorners.top || dragCorners.bottom) {
      dragStyle.value.cursor = ResizeCursor.ns
    } else {
      dragStyle.value.cursor = ResizeCursor.ew
    }
    dragStyle.value.userSelect = 'none'

    isResizing = true
    isDragging = false // Resizing overrides dragging behaviors
    startMouseX = clientX
    startMouseY = clientY
    startPos = currentPos.value.clone()
    startSize = currentSize.value.clone()

    if (e.type === 'mousedown') {
      window.addEventListener('mousemove', onMouseMove)
      window.addEventListener('mouseup', onMouseUp)
    } else {
      window.addEventListener('touchmove', onMouseMove)
      window.addEventListener('touchend', onMouseUp)
    }
  }

  const onMouseMove = (e: MouseEvent | TouchEvent) => {
    let clientX: number
    let clientY: number
    if (e.type === 'mousemove') {
      e = e as MouseEvent
      clientX = e.clientX
      clientY = e.clientY
    } else {
      e = e as TouchEvent
      clientX = e.touches[0].clientX
      clientY = e.touches[0].clientY
    }

    let parentElSize: AbsoluteSize
    if (parentEl && parentEl.value) {
      parentElSize = new AbsoluteSize(parentEl.value.clientWidth, parentEl.value.clientHeight)
    } else {
      parentElSize = new AbsoluteSize(window.innerWidth, window.innerHeight)
    }

    const dx = ((clientX - startMouseX) / parentElSize.w) * 100
    const dy = ((clientY - startMouseY) / parentElSize.h) * 100

    if (isDragging) {
      currentPos.value.y = startPos.y + dy
      currentPos.value.x = startPos.x + dx
    } else if (isResizing) {
      if (dragCorners.top) {
        currentPos.value.y = startPos.y + dy
        currentSize.value.h = startSize.h - dy
      }
      if (dragCorners.bottom) {
        currentSize.value.h = startSize.h + dy
      }
      if (dragCorners.left) {
        currentPos.value.x = startPos.x + dx
        currentSize.value.w = startSize.w - dx
      }
      if (dragCorners.right) {
        currentSize.value.w = startSize.w + dx
      }
    }
  }

  const onMouseUp = () => {
    if (isResizing && onResizeEnd) {
      onResizeEnd()
    }

    isDragging = false
    isResizing = false
    dragStyle.value.cursor = 'auto'
    dragStyle.value.userSelect = 'auto'
    window.removeEventListener('mousemove', onMouseMove)
    window.removeEventListener('mouseup', onMouseUp)
    window.removeEventListener('touchmove', onMouseMove)
    window.removeEventListener('touchend', onMouseUp)
  }

  const maximizeSize = (direction = ResizeDirection.All, lengthSeconds?: number) => {
    if (direction === ResizeDirection.All) {
      isMaximized.value = true

      originalBoxDimensions.value = {
        size: currentSize.value.clone(),
        pos: currentPos.value.clone(),
      }
    }

    if (direction & ResizeDirection.Left) {
      currentSize.value.w += currentPos.value.x
      currentPos.value.x = 0
    }
    if (direction & ResizeDirection.Right) {
      currentSize.value.w += 100 - (currentPos.value.x + currentSize.value.w)
    }
    if (direction & ResizeDirection.Top) {
      currentSize.value.h += currentPos.value.y
      currentPos.value.y = 0
    }
    if (direction & ResizeDirection.Bottom) {
      currentSize.value.h += 100 - (currentPos.value.y + currentSize.value.h)
    }

    if (lengthSeconds) {
      dragStyle.value.transition = `
      width ${lengthSeconds}s, 
      height ${lengthSeconds}s, 
      top ${lengthSeconds}s, 
      left ${lengthSeconds}s`

      setTimeout(() => {
        dragStyle.value.transition = undefined
        if (onResizeEnd) {
          onResizeEnd()
        }
      }, lengthSeconds * 1000)
    }

    if (!lengthSeconds && onResizeEnd) {
      onResizeEnd()
    }
  }

  const restoreSize = (lengthSeconds?: number) => {
    currentPos.value = originalBoxDimensions.value.pos.clone()
    currentSize.value = originalBoxDimensions.value.size.clone()

    if (lengthSeconds) {
      dragStyle.value.transition = `
      width ${lengthSeconds}s, 
      height ${lengthSeconds}s, 
      top ${lengthSeconds}s, 
      left ${lengthSeconds}s`

      setTimeout(() => {
        dragStyle.value.transition = undefined
        if (onResizeEnd) {
          onResizeEnd()
        }
      }, lengthSeconds * 1000)
    }

    isMaximized.value = false

    if (!lengthSeconds && onResizeEnd) {
      onResizeEnd()
    }
  }

  onMounted(() => {
    window.addEventListener('mouseup', onMouseUp)
  })

  onUnmounted(() => {
    window.removeEventListener('mouseup', onMouseUp)
  })

  return {
    currentSize,
    currentPos,
    dragStyle,
    isMaximized,
    maximizeSize,
    restoreSize,
    onDragStart,
    onResizeStart,
  }
}
