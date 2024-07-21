<template>
  <div
    class="twc-window-component window"
    ref="windowEl"
    @mousedown="options.isResizable && !isMaximized ? onResizeStart($event) : undefined"
    @touchstart="options.isResizable && !isMaximized ? onResizeStart($event) : undefined"
    :style="{
      width: boxWidth + '%',
      height: boxHeight + '%',
      top: boxTop + '%',
      left: boxLeft + '%',
      cursor: resizeCursor || undefined,
    }"
  >
    <div
      class="title-bar"
      @mousedown="options.isDraggable && !isMaximized ? onDragStart($event) : undefined"
      @touchstart="options.isDraggable && !isMaximized ? onDragStart($event) : undefined"
      @dblclick.self="isMaximized ? restoreSize() : maximizeSize()"
      :style="{
        borderRadius: isMaximized ? 0 : undefined,
      }"
    >
      <div class="title-bar-text">{{ options.title }}</div>
      <div class="title-bar-controls">
        <button v-if="options.controls?.minimize" aria-label="Minimize" />
        <button
          v-if="options.controls?.maximize"
          :aria-label="isMaximized ? 'Restore' : 'Maximize'"
          @click="isMaximized ? restoreSize() : maximizeSize()"
        />
        <button v-if="options.controls?.close" aria-label="Close" @click="options.onClose" />
      </div>
    </div>
    <WindowToolbarComponent v-if="options.toolbar" :rows="options.toolbar" />
    <div class="window-body">
      <slot />
    </div>
    <div v-if="options.isShowStatusBar" class="status-bar">
      <p v-for="info in options.statusBarInfo || []" :key="info" class="status-bar-field">
        {{ info }}
      </p>
    </div>
    <CoreContextMenuComponent
      ref="contextMenuEl"
      v-if="contextMenu"
      :options="contextMenu"
      :relPos="contextMenuRelativePosition"
    />
  </div>
</template>

<script setup lang="ts">
import { useDraggableResizable } from '@/core/composables/useDragResize'
import CoreContextMenuComponent from './CoreContextMenuComponent.vue'
import WindowToolbarComponent, {
  type WindowToolbarRow,
} from './WindowToolbarComponent.vue'
import type { RelativePosition } from '../models/relative_position'
import type { RelativeSize } from '../models/relative_size'
import { onUpdated, ref } from 'vue'
import { useContext } from '../composables/useContext'

export type WindowTitleBarControls = {
  minimize: boolean
  maximize: boolean
  close: boolean
}

export type WindowOptions = {
  pos: RelativePosition
  size: RelativeSize
  title?: string
  controls?: WindowTitleBarControls
  toolbar?: WindowToolbarRow[]
  statusBarInfo?: string[]
  isDraggable?: boolean
  isResizable?: boolean
  isShowStatusBar?: boolean
  onClose?: () => void
}

const props = defineProps<{
  options: WindowOptions
}>()

const windowEl = ref<HTMLElement>()
const contextMenuEl = ref<HTMLElement>()
const isMaximized = ref<boolean>(false)
const originalBoxValues = ref({
  width: 0,
  height: 0,
  top: 0,
  left: 0,
})

const { contextMenu, contextMenuRelativePosition } = useContext(windowEl, {})
const { boxWidth, boxHeight, boxTop, boxLeft, resizeCursor, onDragStart, onResizeStart } =
  useDraggableResizable(props.options.pos, props.options.size, windowEl)

onUpdated(() => {
  isMaximized.value = checkIsMaximized()
})

function checkIsMaximized() {
  const windowRect = windowEl.value?.getBoundingClientRect()
  const parentRect = windowEl.value?.parentElement?.getBoundingClientRect()

  return windowRect?.width === parentRect?.width && windowRect?.height === parentRect?.height
}

function maximizeSize() {
  originalBoxValues.value = {
    width: boxWidth.value,
    height: boxHeight.value,
    top: boxTop.value,
    left: boxLeft.value,
  }

  boxLeft.value = 0
  boxTop.value = 0
  boxWidth.value = 100
  boxHeight.value = 100
}

function restoreSize() {
  boxLeft.value = originalBoxValues.value.left
  boxTop.value = originalBoxValues.value.top
  boxWidth.value = originalBoxValues.value.width
  boxHeight.value = originalBoxValues.value.height
}
</script>

<style scoped lang="scss">
@use 'xp.css/dist/XP.css';
@use '@/assets/xp.custom';

.twc-window-component {
  position: absolute;
  display: flex;
  flex-direction: column;
  // Just enough base size to display the title bar control buttons
  min-width: 100px;
  min-height: 30px;
}

.title-bar {
  height: auto;
}

.window-body {
  overflow: auto;
  margin: 0 3px;
  flex-grow: 1;
}

.status-bar {
  overflow: hidden;
}

.title-bar-text {
  margin: 0;
  min-width: 0;
  text-wrap: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  pointer-events: none;
  user-select: none;
}
</style>
