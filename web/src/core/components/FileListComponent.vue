<template>
  <div class="container" @click.self="onContainerClick">
    <div
      class="file"
      :class="{ selected: selectedFiles.includes(file.name) }"
      v-for="file in files"
      :key="file.name"
      @click="onFileClick(file)"
      @dblclick="onFileDblClick(file)"
      @touchend="onFileDblClick(file)"
    >
      <img class="file-icon" :src="file.icon" :alt="file.name + ' icon'" />
      <div class="file-name">
        {{ file.name }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { ContextMenuOptions } from './ContextMenuComponent.vue'

export type FileOption = {
  icon: string
  name: string
  onDblClick: () => void
  contextMenuOptions?: ContextMenuOptions
}

export type FileListProps = {
  files: FileOption[]
}

defineProps<FileListProps>()

const selectedFiles = ref<string[]>([])

const onContainerClick = () => {
  selectedFiles.value = []
}

const onFileClick = (file: FileOption) => {
  selectedFiles.value = [file.name]
}

const onFileDblClick = (file: FileOption) => {
  selectedFiles.value = []
  file.onDblClick()
}
</script>

<style scoped lang="scss">
.container {
  padding: 1em;
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  gap: 1em;
}

.file {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  gap: 10px;
  user-select: none;
  width: 6em;
  text-align: center;

  &.selected .file-name {
    background-color: black;
    color: white;
  }
}

.file-name {
  padding: 0.2em;
}

.file-icon {
  width: 4em;
}
</style>
