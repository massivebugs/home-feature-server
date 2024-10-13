<template>
  <div
    v-bind="$attrs"
    class="hfs-window"
    :class="{
      'hfs-window_blocked': isBlocked,
      'hfs-window_static': static,
    }"
    ref="windowEl"
    @mousedown="onWindowMouseDown"
    @touchstart="onWindowMouseDown"
    :style="
      isDragResizeReady && {
        width: fitContent ? undefined : currentSize.w + '%',
        height: fitContent ? undefined : currentSize.h + '%',
        top: currentPos.y + '%',
        left: currentPos.x + '%',
        ...dragStyle,
      }
    "
  >
    <div class="hfs-window__header">
      <div
        v-if="!hideTitlebar"
        class="hfs-window__title-bar"
        @mousedown="onDragStart($event)"
        @touchstart="onDragStart($event)"
        @dblclick.self="onClickToggleSize"
      >
        <div class="hfs-window__title-bar__title">
          <img v-if="titleBarIcon" class="hfs-window__title-bar__icon" :src="titleBarIcon" />
          {{ title }}
          <slot name="title" />
        </div>
        <div class="hfs-window__title-bar__controls">
          <button v-if="controls?.minimize" @click="emit('clickMinimize')" aria-label="Minimize">
            <CollapseIconComponent />
          </button>
          <button
            v-if="controls?.maximize"
            @click="onClickToggleSize"
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
      <slot :window-el="windowEl" :window-size-query="windowSizeQuery" />
    </div>
    <div v-if="statusBarInfo" class="hfs-window__status-bar">
      <p v-for="info in statusBarInfo || []" :key="info">
        {{ info }}
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onBeforeUnmount, onMounted, provide, reactive, ref } from 'vue'
import { ResizeDirection, useDragResize } from '../composables/useDragResize'
import { RelativePosition } from '../models/relativePosition'
import type { RelativeSize } from '../models/relativeSize'
import { useCoreStore } from '../stores'
import CloseIconComponent from './CloseIconComponent.vue'
import CollapseIconComponent from './CollapseIconComponent.vue'
import MinimizeMaximizeIconComponent from './MinimizeMaximizeIconComponent.vue'
import WindowToolbarComponent, { type WindowToolbarRow } from './WindowToolbarComponent.vue'

export type WindowUID = string

export type WindowSizeQuery = {
  md: boolean
  lg: boolean
  xl: boolean
}

export type BlockWindowFunc = (block: boolean) => void

export type ToggleWindowResizeHandlerFunc = (handlerFunc: () => void) => void

export type WindowTitleBarControls = {
  minimize: boolean
  maximize: boolean
  close: boolean
}

const BREAKPOINT_MD = 768
const BREAKPOINT_LG = 992
const BREAKPOINT_XL = 1200
const TOGGLE_SIZE_SECONDS = 0.3

const emit = defineEmits<{
  (e: 'clickClose'): void
  (e: 'clickMinimize'): void
}>()

const props = defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
  fitContent?: boolean
  title?: string
  hideTitlebar?: boolean
  controls?: WindowTitleBarControls
  toolbar?: WindowToolbarRow[]
  statusBarInfo?: string[]
  resizable?: boolean
  static?: boolean
  titleBarIcon?: string
}>()

const store = useCoreStore()
const windowEl = ref<HTMLElement>()
const isBlocked = ref<boolean>(false)
const windowIdx = store.processes.size
const {
  isDragResizeReady,
  currentSize,
  currentPos,
  dragStyle,
  isMaximized,
  maximizeSize,
  restoreSize,
  onDragStart,
  onResizeStart,
} = useDragResize(
  windowEl,
  props.pos
    ? props.pos
    : !props.pos && props.size
      ? // TODO: This doesn't make sense as it's relative size...
        new RelativePosition(windowIdx + 50 - props.size.w / 2, windowIdx + 50 - props.size.h / 2)
      : undefined,
  props.size,
  undefined,
  undefined,
  () => {
    const event = new CustomEvent('resize')
    windowEl.value?.dispatchEvent(event)
  },
)
const windowSizeQuery = reactive<WindowSizeQuery>({
  md: false,
  lg: false,
  xl: false,
})
let resizeObserver: ResizeObserver | null = null

// TODO: Remove resize listener and use resize observer only
const addWindowResizeListener: ToggleWindowResizeHandlerFunc = (handlerFunc: () => void) => {
  windowEl.value?.addEventListener('resize', handlerFunc)
}

const removeWindowResizeListener: ToggleWindowResizeHandlerFunc = (handlerFunc: () => void) => {
  windowEl.value?.removeEventListener('resize', handlerFunc)
}

const blockWindow: BlockWindowFunc = (blocked: boolean) => {
  isBlocked.value = blocked
}

provide('windowSizeQuery', windowSizeQuery)
provide('addWindowResizeListener', addWindowResizeListener)
provide('removeWindowResizeListener', removeWindowResizeListener)
provide('blockParentWindow', blockWindow)

const onWindowMouseDown = (e: MouseEvent | TouchEvent) => {
  if (props.resizable && !isBlocked.value) {
    onResizeStart(e)
  }
}

const onClickToggleSize = () => {
  if (props.resizable) {
    isMaximized.value
      ? restoreSize(TOGGLE_SIZE_SECONDS)
      : maximizeSize(ResizeDirection.All, TOGGLE_SIZE_SECONDS)
  }
}

onMounted(() => {
  if (windowEl.value) {
    resizeObserver = new ResizeObserver((entries) => {
      entries.forEach((entry) => {
        const width = entry.contentRect.width
        windowSizeQuery.md = width >= BREAKPOINT_MD
        windowSizeQuery.lg = width >= BREAKPOINT_LG
        windowSizeQuery.xl = width >= BREAKPOINT_XL
      })
    })

    resizeObserver.observe(windowEl.value)
  }
})

onBeforeUnmount(() => {
  if (resizeObserver && windowEl.value) {
    resizeObserver.unobserve(windowEl.value)
  }
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
  z-index: 1;
  overflow: hidden;

  &.hfs-window_blocked::after {
    border-radius: inherit;
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: colors.$low-opacity-dark-grey;
  }
}

.hfs-window:not(.hfs-window_static) {
  box-shadow: 1px 1px 6px 2px rgba(0, 0, 0, 0.4);
  -webkit-box-shadow: 1px 1px 6px 2px rgba(0, 0, 0, 0.4);
  -moz-box-shadow: 1px 1px 6px 2px rgba(0, 0, 0, 0.4);
}

// https://stackoverflow.com/questions/77143569/position-fixed-not-working-when-backdrop-filter-is-used
.hfs-window:not(.hfs-window_static)::before {
  border-radius: 10px;
  position: absolute;
  content: '';
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: colors.$high-opacity-white;
  @supports (backdrop-filter: blur()) {
    backdrop-filter: blur(10px);
    background-color: colors.$low-opacity-white;
  }
  z-index: -1;
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
  padding-top: 0.2em;
  display: flex;
  align-items: center;
  gap: 0.3em;
}

.hfs-window__title-bar__icon {
  height: 1em;
  max-width: 1em;
  margin-right: 0.3em;
}

.hfs-window__title-bar__controls {
  display: flex;
  & > button {
    border: 0;
    background: none;
    border-radius: 50%;
    width: 1.5em;
    height: 1.5em;
    display: flex;
    justify-content: center;
    align-items: center;
    stroke: colors.$black;
    transition: background-color 0.2s;

    &:hover {
      svg {
        stroke: colors.$white;
        fill: colors.$white;
      }
      &[aria-label='Minimize'] {
        background-color: colors.$dark-grey;
      }
      &[aria-label='Restore'],
      &[aria-label='Maximize'] {
        background-color: colors.$skobeloff;
      }
      &[aria-label='Close'] {
        background-color: colors.$red-cmyk;
      }
    }
  }
}

.hfs-window__contents {
  overflow: auto;
  flex: 1;
}

.hfs-window__status-bar {
  user-select: none;
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
