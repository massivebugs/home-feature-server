<template>
  <nav class="window-toolbar">
    <template v-for="row in rows" :key="row">
      <ul v-if="row.isMenu" class="window-menu">
        <li
          v-for="(item, idx) in row.items"
          :key="idx"
          @click.stop="onMenuItemClick($event, item)"
          @mouseover="contextMenu && onMenuItemClick($event, item)"
        >
          {{ item.label }}
        </li>
      </ul>
      <div v-else>
        <div v-for="(item, idx) in row.items" :key="idx">{{ item.label }}</div>
      </div>
    </template>
  </nav>
</template>

<script setup lang="ts">
import { type GetContextualMetrics, type SetContextMenu } from '../composables/useContext'
import type { ContextMenuOptions } from './CoreContextMenuComponent.vue'
import { inject, type Ref } from 'vue'

export type WindowToolbarItem = {
  label?: string
  contextMenuOptions?: ContextMenuOptions
}

export type WindowToolbarRow = {
  items: WindowToolbarItem[]
  isMenu?: boolean // Items in a Menu have a fixed UI and does not show dropdown arrow icon
}

defineProps<{
  rows: WindowToolbarRow[]
}>()

const contextMenu = inject<Ref<ContextMenuOptions | null>>('contextMenu')
const setContextMenu = inject('setContextMenu') as SetContextMenu
const getContextualMetrics = inject('getContextualMetrics') as GetContextualMetrics

function onMenuItemClick(e: Event, item: WindowToolbarItem) {
  if (item.contextMenuOptions) {
    const { relPos, relSize } = getContextualMetrics(e.target as HTMLElement)
    setContextMenu(item.contextMenuOptions ?? null, relPos.add(0, relSize.height))
  }
}
</script>

<style scoped lang="scss">
@use '@/assets/theme.default';
</style>
