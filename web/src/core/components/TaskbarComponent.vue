<template>
  <div class="hfs-taskbar">
    <LogOutIconComponent class="hfs-taskbar__log-out-icon" @click="emit('clickLogOut')" />
    <FileShortcutIconComponent
      v-for="process in runningProcesses.values()"
      class="hfs-taskbar__file-icon"
      :class="{
        'hfs-taskbar__selected': selectedProcessId === process.id,
      }"
      :key="process.id"
      :option="{ programId: process.program.id, icon: process.program.icon, name: '' }"
      :hide-name="true"
      @click="emit('selectProcess', { processId: process.id })"
    />
  </div>
</template>
<script setup lang="ts">
import type { Process } from '../models/process'
import FileShortcutIconComponent from './FileShortcutIconComponent.vue'
import LogOutIconComponent from './LogOutIconComponent.vue'

export type TaskbarSelectProcessEvent = {
  processId: string
}

const emit = defineEmits<{
  (e: 'clickLogOut'): void
  (e: 'selectProcess', payload: TaskbarSelectProcessEvent): void
}>()

defineProps<{
  runningProcesses: Map<string, Process>
  selectedProcessId?: string | null
}>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-taskbar {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 0.7em;
  padding: 0.5em 1em;
  background-color: colors.$mid-opacity-white;
  border-radius: 10px;
  height: 6em;
  overflow: hidden;
}

.hfs-taskbar__log-out-icon {
  cursor: pointer;
  height: 4em;
  width: 4em;
  transition: transform 1s;
  stroke: colors.$black;

  &:hover {
    stroke: desaturate(colors.$red-cmyk, 20);
  }
}

.hfs-taskbar__log-out-icon,
.hfs-taskbar__file-icon {
  position: relative;
  flex: 1;
  cursor: pointer;
  height: calc(100% - 1em);
  width: 4em;
  transition:
    padding-bottom 0.1s,
    height 0.1s;
  transition-delay: 0.1s;

  &.hfs-taskbar__selected,
  &:hover {
    height: calc(100%);
    padding-bottom: 1em;
    transition-delay: 0s;
  }

  &:after {
    transition:
      width 0.1s,
      left 0.1s;
    transition-delay: 0s;
    content: '';
    position: absolute;
    left: 50%;
    bottom: 0;
    height: 5px;
    width: 0;
    background-color: colors.$light-grey;
    border-radius: 5px;
  }

  &.hfs-taskbar__selected:after {
    background-color: colors.$skobeloff !important;
  }

  &.hfs-taskbar__selected:after,
  &:hover:after {
    width: 100%;
    left: 0;
    transition-delay: 0.1s;
  }
}
</style>
