<template>
  <main ref="desktopViewEl">
    <template v-for="(process, idx) in store.processes.values()" :key="process.id">
      <WindowComponent
        :pos="
          new RelativePosition(
            idx + 50 - process.program.windowOptions.size.width / 2,
            idx + 50 - process.program.windowOptions.size.height / 2,
          )
        "
        :options="process.program.windowOptions"
        @click-close="() => onClickWindowClose(process.id)"
      >
        <component
          :is="process.program.component"
          v-bind="process.program.componentProps"
        ></component>
      </WindowComponent>
    </template>
    <ContextMenuComponent
      ref="contextMenuEl"
      v-if="contextMenuOptions"
      :options="contextMenuOptions"
      :pos="contextMenuPos"
    />
  </main>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, provide, ref } from 'vue'
import ContextMenuComponent, {
  type ContextMenuOptions,
} from '../components/ContextMenuComponent.vue'
import WindowComponent from '../components/WindowComponent.vue'
import type { AbsolutePosition } from '../models/absolute_position'
import { RelativePosition } from '../models/relative_position'
import { useStore } from '../stores'
import { getRelativeParentPosition } from '../utils/element'

export type SetContextMenu = (
  newContextMenu: ContextMenuOptions | null,
  pos?: AbsolutePosition,
) => void

const store = useStore()
const desktopViewEl = ref()
const contextMenuEl = ref<HTMLElement>()
const contextMenuOptions = ref<ContextMenuOptions | null>(null)
const contextMenuPos = ref<RelativePosition>(new RelativePosition(0, 0))

const setContextMenu: SetContextMenu = (
  options: ContextMenuOptions | null,
  pos?: AbsolutePosition,
) => {
  contextMenuOptions.value = options
  if (desktopViewEl.value && pos) {
    contextMenuPos.value = getRelativeParentPosition(pos, desktopViewEl.value)
  }
}

const clearContextMenu = () => {
  setContextMenu(null)
}

provide('contextMenu', contextMenuOptions)
provide('setContextMenu', setContextMenu)

const onClickWindowClose = (processId: string) => {
  store.removeProcess(processId)
}

onMounted(() => {
  window.addEventListener('click', clearContextMenu)
})

onUnmounted(() => {
  window.removeEventListener('click', clearContextMenu)
})
</script>

<style scoped lang="scss">
main {
  width: 100%;
  height: 100%;
}
</style>
