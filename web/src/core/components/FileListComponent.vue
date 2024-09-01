<template>
  <div class="file-list__container" @click.self="onContainerClick">
    <FileShortcutIconComponent
      v-for="file in files"
      class="file-list__shortcut-icon"
      :key="file.name"
      :selected="selectedFiles.includes(file.name)"
      :option="file"
      @click="onFileClick(file)"
      @dblclick="onFileDblClick(file)"
      @touchend="onFileDblClick(file)"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import FileShortcutIconComponent, {
  type FileShortcutIconOption,
} from './FileShortcutIconComponent.vue'

export type FileListProps = {
  files: FileShortcutIconOption[]
}

defineProps<FileListProps>()

const selectedFiles = ref<string[]>([])

const onContainerClick = () => {
  selectedFiles.value = []
}

const onFileClick = (file: FileShortcutIconOption) => {
  selectedFiles.value = [file.name]
}

const onFileDblClick = (file: FileShortcutIconOption) => {
  selectedFiles.value = []
  if (file.onDblClick) {
    file.onDblClick()
  }
}
</script>

<style scoped lang="scss">
.file-list__container {
  padding: 1em;
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  gap: 1em;
}

.file-list__shortcut-icon {
  width: 5em;
}
</style>
