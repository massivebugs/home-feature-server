<template>
  <main ref="desktopViewEl">
    <FileListComponent class="file-list" :files="fileOptions" />
    <template v-for="process in store.processes.values()" :key="process.id">
      <component
        :is="process.program.component"
        v-bind="process.program.componentProps"
        @click-close="onClickWindowClose(process.id)"
      />
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
import { uniqueId } from 'lodash'
import { onMounted, onUnmounted, provide, ref } from 'vue'
import ContextMenuComponent, {
  type ContextMenuOptions,
} from '../components/ContextMenuComponent.vue'
import FileListComponent, { type FileOption } from '../components/FileListComponent.vue'
import type { AbsolutePosition } from '../models/absolute_position'
import { Process } from '../models/process'
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
const fileOptions = ref<FileOption[]>([])

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

  store.programs.forEach((program) => {
    fileOptions.value.push({
      name: program.name,
      icon: program.icon,
      onDblClick: () => {
        const process = new Process(uniqueId('pid_'), program)
        store.addProcess(process)
      },
    })
  })
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

.file-list {
  height: 100%;
  width: 100%;
}
</style>
