<template>
  <main ref="desktopViewEl" class="hfs-desktop">
    <h1 class="hfs-desktop__logo">{{ t('app.name') }} v.0.0.1</h1>
    <FileListComponent class="hfs-desktop__file-list" :files="fileOptions" />
    <div class="hfs-desktop__taskbar-container">
      <TaskbarComponent
        class="hfs-desktop__taskbar"
        :running-processes="store.processesByInsertOrder"
        :selected-process-id="store.topLevelProcessId"
        @click-log-out="onClickLogOut"
        @select-process="onSelectProcess"
      />
    </div>
    <template v-for="process in store.processes.values()" :key="process.id">
      <component
        v-show="!hiddenWindowPIDs.has(process.id)"
        :is="process.program.component"
        v-bind="process.program.componentProps"
        @mousedown="store.setTopLevelProcess(process.id)"
        @click-close="onClickWindowClose(process.id)"
        @click-cancel="onClickWindowCancel(process.id)"
        @click-minimize="onClickWindowMinimize(process.id)"
      />
    </template>
    <ConfirmDialogComponent
      v-if="showLogOutConfirmDialog"
      @click-success="onSuccessConfirmLogOutDialog"
      @click-close="onCloseConfirmLogOutDialog"
      @click-cancel="onCloseConfirmLogOutDialog"
      pos="center"
      :title="t('desktop.logOutDialogTitle')"
      :message="t('desktop.logOutConfirmMessage')"
    />
    <ContextMenuComponent
      ref="contextMenuEl"
      v-if="contextMenuOptions"
      :options="contextMenuOptions"
      :pos="contextMenuPos"
    />
  </main>
</template>

<script setup lang="ts">
import { AxiosError, HttpStatusCode } from 'axios'
import { uniqueId } from 'lodash'
import { onMounted, onUnmounted, provide, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute, useRouter } from 'vue-router'
import ConfirmDialogComponent from '../components/ConfirmDialogComponent.vue'
import ContextMenuComponent, {
  type ContextMenuOptions,
} from '../components/ContextMenuComponent.vue'
import FileListComponent from '../components/FileListComponent.vue'
import type { FileShortcutIconOption } from '../components/FileShortcutIconComponent.vue'
import TaskbarComponent, {
  type TaskbarSelectProcessEvent,
} from '../components/TaskbarComponent.vue'
import type { AbsolutePosition } from '../models/absolutePosition'
import { Process } from '../models/process'
import { RelativePosition } from '../models/relativePosition'
import { useCoreStore } from '../stores'
import { getRelativeParentPosition } from '../utils/element'

export type SetContextMenu = (
  newContextMenu: ContextMenuOptions | null,
  pos?: AbsolutePosition,
) => void

const { t, locale } = useI18n()
const route = useRoute()
const router = useRouter()
const store = useCoreStore()
const desktopViewEl = ref()
const contextMenuEl = ref<HTMLElement>()
const contextMenuOptions = ref<ContextMenuOptions | null>(null)
const contextMenuPos = ref<RelativePosition>(new RelativePosition(0, 0))
const fileOptions = ref<FileShortcutIconOption[]>([])
const showLogOutConfirmDialog = ref<boolean>(false)
const hiddenWindowPIDs = ref<Set<string>>(new Set())

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

const onClickWindowMinimize = (processId: string) => {
  hiddenWindowPIDs.value.add(processId)
  store.topLevelProcessId = null
}

const onClickLogOut = () => {
  showLogOutConfirmDialog.value = true
}

const onSuccessConfirmLogOutDialog = async () => {
  showLogOutConfirmDialog.value = false

  // Remove API token and reload login page
  // TODO: Call server endpoint to remove cookie for us

  await router.push({ name: 'login' })
  router.go(0)
}

const onCloseConfirmLogOutDialog = () => {
  showLogOutConfirmDialog.value = false
}

const onSelectProcess = (payload: TaskbarSelectProcessEvent) => {
  hiddenWindowPIDs.value.delete(payload.processId)
  store.setTopLevelProcess(payload.processId)
}

onMounted(async () => {
  window.addEventListener('click', clearContextMenu)

  // Load user data
  try {
    const res = await store.getUserSystemPreference()
    store.systemPreference = res.data.user_system_preference

    locale.value = store.systemPreference.language ?? navigator.language
  } catch (error) {
    if (error instanceof AxiosError && error.status === HttpStatusCode.NotFound) {
      try {
        const res = await store.createUserSystemPreference()
        store.systemPreference = res.data.user_system_preference
      } catch (error) {
        // TODO: Error handling
      }
    } else {
      // TODO: Error handling
    }
  }

  store.programs.forEach((program) => {
    fileOptions.value.push({
      name: t(`${program.id}.name`),
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
@use '@/assets/media-query';
@use '@/assets/colors';
$taskbarHeight: 4em;
$taskbarBottom: 0.7em;
$taskbarHeightMd: 5em;
$taskbarBottomMd: 1em;

.hfs-desktop {
  position: relative;
  width: 100%;
  height: calc(100% - $taskbarHeight - $taskbarBottom - 0.2em);
  @include media-query.md {
    height: calc(100% - $taskbarHeightMd - $taskbarBottomMd - 0.2em);
  }

  > .hfs-window:not(.hfs-dialog) {
    position: absolute !important;
  }
}

.hfs-desktop__logo {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -150%);
  color: colors.$light-grey;
  font-family: monospace;
  font-size: calc(1vw + 1em);
  font-style: italic;
  user-select: none;
  pointer-events: none;
}

.hfs-desktop__file-list {
  position: relative;
  height: 100%;
  width: 100%;
}

// Make height shorter on smaller screens!
.hfs-desktop__taskbar-container {
  position: fixed;
  left: 5%;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 90%;
  bottom: $taskbarBottom;
  @include media-query.md {
    bottom: $taskbarBottomMd;
  }

  > .hfs-desktop__taskbar {
    height: $taskbarHeight;
    @include media-query.md {
      height: $taskbarHeightMd;
    }
  }
}
</style>
