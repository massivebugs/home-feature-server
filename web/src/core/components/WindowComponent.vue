<template>
  <div
    v-bind="$attrs"
    class="window"
    :class="{
      focused: store.focusedWindowUid === windowUid,
      blocked: isBlocked,
    }"
    ref="windowEl"
    @mousedown.stop="onWindowMouseDown"
    @touchstart.stop="onWindowMouseDown"
    :style="{
      width: boxWidth + '%',
      height: boxHeight + '%',
      top: boxTop + '%',
      left: boxLeft + '%',
      ...dragStyle,
    }"
  >
    <div
      class="title-bar"
      @mousedown="!isMaximized ? onDragStart($event) : undefined"
      @touchstart="!isMaximized ? onDragStart($event) : undefined"
      @dblclick.self="onToolbarDblClick"
      :style="{
        borderRadius: isMaximized ? 0 : undefined,
      }"
    >
      <div class="title-bar-title">{{ title }}</div>
      <div class="title-bar-controls">
        <button v-if="controls?.minimize" aria-label="Minimize">_</button>
        <button
          v-if="controls?.maximize"
          @click="isMaximized ? restoreSize() : maximizeSize()"
          :aria-label="isMaximized ? 'Restore' : 'Maximize'"
        >
          {{ isMaximized ? '■' : '□' }}
        </button>
        <button v-if="controls?.close" aria-label="Close" @click="emit('clickClose')">X</button>
      </div>
    </div>
    <WindowToolbarComponent v-if="toolbar" :rows="toolbar" />
    <div class="window-body">
      <slot />
    </div>
    <div v-if="statusBarInfo" class="status-bar">
      <p v-for="info in statusBarInfo || []" :key="info">
        {{ info }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { uniqueId } from 'lodash'
import { onMounted, onUpdated, provide, ref } from 'vue'
import { useDraggableResizable } from '../composables/useDragResize'
import { RelativePosition } from '../models/relative_position'
import type { RelativeSize } from '../models/relative_size'
import { useStore } from '../stores'
import WindowToolbarComponent, { type WindowToolbarRow } from './WindowToolbarComponent.vue'

export type IBlockWindowFunc = (block: boolean) => void

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const props = defineProps<{
  pos?: RelativePosition
  size: RelativeSize
  title?: string
  controls?: {
    minimize: boolean
    maximize: boolean
    close: boolean
  }
  toolbar?: WindowToolbarRow[]
  statusBarInfo?: string[]
  isResizable?: boolean
}>()

const store = useStore()
const windowUid = uniqueId('window_')
const windowEl = ref<HTMLElement>()
const isBlocked = ref<boolean>(false)
const isMaximized = ref<boolean>(false)
const originalBoxValues = ref({
  width: 0,
  height: 0,
  top: 0,
  left: 0,
})
const windowIdx = store.processes.size
const { boxWidth, boxHeight, boxTop, boxLeft, dragStyle, onDragStart, onResizeStart } =
  useDraggableResizable(
    props.pos ??
      new RelativePosition(
        windowIdx + 50 - props.size.width / 2,
        windowIdx + 50 - props.size.height / 2,
      ),
    props.size,
    windowEl,
  )

const blockWindow: IBlockWindowFunc = (blocked: boolean) => {
  isBlocked.value = blocked
}

provide('blockParentWindow', blockWindow)

const checkIsMaximized = () => {
  const windowRect = windowEl.value?.getBoundingClientRect()
  const parentRect = windowEl.value?.parentElement?.getBoundingClientRect()

  return windowRect?.width === parentRect?.width && windowRect?.height === parentRect?.height
}

const maximizeSize = () => {
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

const restoreSize = () => {
  boxLeft.value = originalBoxValues.value.left
  boxTop.value = originalBoxValues.value.top
  boxWidth.value = originalBoxValues.value.width
  boxHeight.value = originalBoxValues.value.height
}

const onWindowMouseDown = (e: MouseEvent | TouchEvent) => {
  store.focusedWindowUid = windowUid
  if (props.isResizable) {
    isMaximized.value ? undefined : onResizeStart(e)
  }
}

const onToolbarDblClick = () => {
  if (props.isResizable) {
    isMaximized.value ? restoreSize() : maximizeSize()
  }
}

onMounted(() => {
  store.focusedWindowUid = windowUid
})

onUpdated(() => {
  isMaximized.value = checkIsMaximized()
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.window {
  position: fixed;
  display: flex;
  flex-direction: column;
  // Just enough base size to display the title bar control buttons
  min-width: 100px;
  min-height: 30px;

  background: colors.$white;
  border: 1px solid colors.$black;

  &.focused {
    z-index: 999;
  }

  &.blocked {
    pointer-events: none;
  }

  .window {
    pointer-events: all;
  }
}

.title-bar {
  background-color: colors.$black;
  color: colors.$white;
  display: flex;
  justify-content: space-between;
  padding: 3px;
  font-weight: 500;
  user-select: none;
}

.title-bar-title {
  margin: 0;
  min-width: 0;
  text-wrap: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  pointer-events: none;
  user-select: none;
  margin-top: 0.2em;
}

.window-body {
  overflow: auto;
  flex-grow: 1;
}

.status-bar {
  overflow: hidden;
  display: flex;
  > p {
    margin: 0;
    padding: 3px;
    flex-grow: 1;
    border-top: 1px solid colors.$black;
    &:not(:last-child) {
      border-right: 1px solid colors.$black;
    }
  }
}
</style>
