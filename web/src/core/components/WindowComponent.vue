<template>
  <div
    class="twc-window-component twc-window"
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
      class="twc-title-bar"
      @mousedown="options.isDraggable && !isMaximized ? onDragStart($event) : undefined"
      @touchstart="options.isDraggable && !isMaximized ? onDragStart($event) : undefined"
      @dblclick.self="isMaximized ? restoreSize() : maximizeSize()"
      :style="{
        borderRadius: isMaximized ? 0 : undefined,
      }"
    >
      <div class="twc-title-bar-title">{{ options.title }}</div>
      <div class="twc-title-bar-controls">
        <button
          class="twc-title-bar-close-btn"
          v-if="options.controls?.minimize"
          aria-label="Minimize"
        />
        <button
          :class="[isMaximized ? 'twc-title-bar-restore-btn' : 'twc-title-bar-maximize-btn']"
          v-if="options.controls?.maximize"
          @click="isMaximized ? restoreSize() : maximizeSize()"
          :aria-label="isMaximized ? 'Restore' : 'Maximize'"
        />
        <button
          class="twc-title-bar-close-btn"
          v-if="options.controls?.close"
          aria-label="Close"
          @click="options.onClose"
        />
      </div>
    </div>
    <WindowToolbarComponent v-if="options.toolbar" :rows="options.toolbar" />
    <div class="twc-window-body">
      <slot />
    </div>
    <div v-if="options.isShowStatusBar" class="twc-status-bar">
      <p v-for="info in options.statusBarInfo || []" :key="info">
        {{ info }}
      </p>
    </div>
    <ContextMenuComponent
      ref="contextMenuEl"
      v-if="contextMenu"
      :options="contextMenu"
      :relPos="contextMenuRelativePosition"
    />
  </div>
</template>

<script setup lang="ts">
import { useDraggableResizable } from '@/core/composables/useDragResize'
import ContextMenuComponent from './ContextMenuComponent.vue'
import WindowToolbarComponent, { type WindowToolbarRow } from './WindowToolbarComponent.vue'
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
@use '@/assets/theme.default';

.twc-window-component {
  position: absolute;
  display: flex;
  flex-direction: column;
  // Just enough base size to display the title bar control buttons
  min-width: 100px;
  min-height: 30px;
}
</style>
