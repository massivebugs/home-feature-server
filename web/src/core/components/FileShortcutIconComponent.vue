<template>
  <div class="file-shortcut-icon" :class="{ 'file-shortcut-icon_selected': selected }">
    <slot name="icon">
      <div
        class="file-shortcut-icon__img"
        :style="{
          backgroundImage: `url(${option.icon})`,
        }"
      />
    </slot>
    <div v-if="option.name && !hideName" class="file-shortcut-icon__name">
      {{ option.name }}
    </div>
  </div>
</template>

<script setup lang="ts">
import type { ContextMenuOptions } from './ContextMenuComponent.vue'

export type FileShortcutIconOption = {
  icon: string
  name: string
  onDblClick?: () => void
  contextMenuOptions?: ContextMenuOptions
}

defineProps<{
  option: FileShortcutIconOption
  selected?: boolean
  hideName?: boolean
}>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.file-shortcut-icon {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  justify-content: center;
  gap: 2px;
  user-select: none;
  height: 6em;
  text-align: center;
}

.file-shortcut-icon_selected .file-shortcut-icon__name {
  background-color: colors.$high-opacity-dark-grey;
  color: colors.$white;
  outline: 1px dotted colors.$white;
  outline-offset: -1px;
}

.file-shortcut-icon__name {
  padding: 0.2em;
  max-width: 100%;
}

.file-shortcut-icon__img {
  flex: 1;
  width: 100%;
  background-position: center;
  background-repeat: no-repeat;
  background-size: contain;
}
</style>
