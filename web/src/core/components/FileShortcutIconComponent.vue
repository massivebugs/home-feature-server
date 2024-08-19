<template>
  <div class="file-shortcut-icon" :class="{ 'file-shortcut-icon_selected': selected }">
    <img class="file-shortcut-icon__img" :src="option.icon" :alt="option.name + ' icon'" />
    <div class="file-shortcut-icon__name">
      {{ option.name }}
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ContextMenuOptions } from './ContextMenuComponent.vue'

export type FileShortcutIconOption = {
  icon: string
  name: string
  onDblClick: () => void
  contextMenuOptions?: ContextMenuOptions
}

defineProps<{ option: FileShortcutIconOption; selected: boolean }>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.file-shortcut-icon {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  gap: 2px;
  user-select: none;
  width: 6em;
  text-align: center;
}

.file-shortcut-icon_selected .file-shortcut-icon__name {
  background-color: black;
  color: white;
  outline: 1px dotted white;
  outline-offset: -1px;
}

.file-shortcut-icon__name {
  padding: 0.2em;
}

.file-shortcut-icon__img {
  width: 5em;
}
</style>
