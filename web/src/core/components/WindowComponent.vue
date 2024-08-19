<template>
  <div
    v-bind="$attrs"
    class="hfs-window"
    :class="{
      focused: store.focusedWindowUid === windowUid,
      blocked: isBlocked,
      'hfs-window__static': static,
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
    <div v-if="!hideTitlebar && toolbar" class="hfs-window__header">
      <div
        v-if="!hideTitlebar"
        class="hfs-window__title-bar"
        @mousedown="!isMaximized ? onDragStart($event) : undefined"
        @touchstart="!isMaximized ? onDragStart($event) : undefined"
        @dblclick.self="onTitlebarDblClick"
        :style="{
          borderRadius: isMaximized ? 0 : undefined,
        }"
      >
        <div class="hfs-window__title-bar__title">{{ title }}</div>
        <div class="hfs-window__title-bar__controls">
          <button v-if="controls?.minimize" aria-label="Minimize">
            <CollapseIconComponent />
          </button>
          <button
            v-if="controls?.maximize"
            @click="isMaximized ? restoreSize() : maximizeSize()"
            :aria-label="isMaximized ? 'Restore' : 'Maximize'"
          >
            <MinimizeMaximizeIconComponent :type="isMaximized ? 'minimize' : 'maximize'" />
          </button>
          <button v-if="controls?.close" aria-label="Close" @click="emit('clickClose')">
            <CloseIconComponent />
          </button>
        </div>
      </div>
      <WindowToolbarComponent v-if="toolbar" :rows="toolbar" />
    </div>
    <div class="hfs-window__contents">
      <slot :window-el="windowEl" />
    </div>
    <div v-if="statusBarInfo" class="hfs-window__status-bar">
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
import { useCoreStore } from '../stores'
import CloseIconComponent from './CloseIconComponent.vue'
import CollapseIconComponent from './CollapseIconComponent.vue'
import MinimizeMaximizeIconComponent from './MinimizeMaximizeIconComponent.vue'
import WindowToolbarComponent, { type WindowToolbarRow } from './WindowToolbarComponent.vue'

export type IBlockWindowFunc = (block: boolean) => void

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const props = defineProps<{
  pos?: RelativePosition
  size: RelativeSize
  title?: string
  hideTitlebar?: boolean
  controls?: {
    minimize: boolean
    maximize: boolean
    close: boolean
  }
  toolbar?: WindowToolbarRow[]
  statusBarInfo?: string[]
  resizable?: boolean
  static?: boolean
}>()

const store = useCoreStore()
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
    undefined,
    undefined,
    () => {
      const event = new CustomEvent('resize')
      windowEl.value?.dispatchEvent(event)
    },
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
  if (props.resizable) {
    isMaximized.value ? undefined : onResizeStart(e)
  }
}

const onTitlebarDblClick = () => {
  if (props.resizable) {
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

.hfs-window {
  position: fixed;
  display: flex;
  flex-direction: column;
  // Just enough base size to display the title bar control buttons
  min-width: 100px;
  min-height: 30px;
  border-radius: 10px;

  &.focused {
    z-index: 999;
  }

  &.blocked {
    pointer-events: none;
  }

  .hfs-window {
    pointer-events: all;
  }
}

.hfs-window:not(.hfs-window__static) {
  background-color: colors.$high-opacity-white;
  @supports (backdrop-filter: blur()) {
    backdrop-filter: blur(10px);
    background-color: colors.$low-opacity-white;
  }

  box-shadow: 1px 1px 6px 2px rgba(0, 0, 0, 0.4);
  -webkit-box-shadow: 1px 1px 6px 2px rgba(0, 0, 0, 0.4);
  -moz-box-shadow: 1px 1px 6px 2px rgba(0, 0, 0, 0.4);
}

.hfs-window__title-bar {
  display: flex;
  justify-content: space-between;
  padding: 3px 3px 3px 10px;
  font-weight: 500;
  user-select: none;
}

.hfs-window__title-bar__title {
  margin: 0;
  min-width: 0;
  text-wrap: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  pointer-events: none;
  user-select: none;
  margin-top: 0.2em;
}

.hfs-window__title-bar__controls {
  & > button {
    border: 0;
    background: none;

    &:hover {
      border-radius: 5px;
      background-color: colors.$light-grey;
    }
  }
}

.hfs-window__contents {
  overflow: auto;
  flex-grow: 1;
}

.hfs-window__status-bar {
  overflow: hidden;
  display: flex;
  font-size: 0.9em;

  > p {
    margin: 0;
    padding: 3px 3px 3px 10px;
    flex-grow: 1;
    &:not(:last-child) {
      box-shadow: 0px 1px 10px 0px rgba(0, 0, 0, 0.4);
      -webkit-box-shadow: 0px 1px 10px 0px rgba(0, 0, 0, 0.4);
      -moz-box-shadow: 0px 1px 10px 0px rgba(0, 0, 0, 0.4);
    }
  }
}
</style>
