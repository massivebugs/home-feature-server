<template>
  <main ref="desktopViewEl">
    <h1 class="hfs-desktop__logo">massivebugs.github.io</h1>
    <FileListComponent class="file-list" :files="fileOptions" />
    <template v-for="process in store.processes.values()" :key="process.id">
      <component
        :is="process.program.component"
        v-bind="process.program.componentProps"
        @mousedown="setTopLevelProcess(process.id)"
        @click-close="onClickWindowClose(process.id)"
        @click-cancel="onClickWindowCancel(process.id)"
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
import { useRoute } from 'vue-router'
import ContextMenuComponent, {
  type ContextMenuOptions,
} from '../components/ContextMenuComponent.vue'
import FileListComponent from '../components/FileListComponent.vue'
import type { FileShortcutIconOption } from '../components/FileShortcutIconComponent.vue'
import type { AbsolutePosition } from '../models/absolute_position'
import { Process } from '../models/process'
import { RelativePosition } from '../models/relative_position'
import { useCoreStore } from '../stores'
import { getRelativeParentPosition } from '../utils/element'

export type SetContextMenu = (
  newContextMenu: ContextMenuOptions | null,
  pos?: AbsolutePosition,
) => void

const route = useRoute()
const store = useCoreStore()
const desktopViewEl = ref()
const contextMenuEl = ref<HTMLElement>()
const contextMenuOptions = ref<ContextMenuOptions | null>(null)
const contextMenuPos = ref<RelativePosition>(new RelativePosition(0, 0))
const fileOptions = ref<FileShortcutIconOption[]>([])

const setTopLevelProcess = (processId: string) => {
  const process = store.processes.get(processId)
  if (process) {
    store.removeProcess(processId)
    store.addProcess(process)
  }
}

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

const onClickWindowCancel = (processId: string) => {
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

  if (route.params.programId) {
    const process = store.findProgramProcesses(route.params.programId as string)
    const program = store.programs.get(route.params.programId as string)
    if (process.length) {
      console.log(process[0])
    } else if (program) {
      const process = new Process(uniqueId('pid_'), program)
      store.addProcess(process)
    } else {
      console.log('No programs found!')
      // Show an error dialog
    }
  }
})

onUnmounted(() => {
  window.removeEventListener('click', clearContextMenu)
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

main {
  position: relative;
  width: 100%;
  height: 100%;
}

.hfs-desktop__logo {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -150%);
  color: colors.$light-grey;
  font-family: monospace;
  font-size: 4em;
  font-style: italic;
  user-select: none;
  pointer-events: none;
}

.file-list {
  height: 100%;
  width: 100%;
}
</style>
