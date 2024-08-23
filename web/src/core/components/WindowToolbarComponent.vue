<template>
  <nav class="window-toolbar">
    <template v-for="row in rows" :key="row">
      <ul v-if="row.isMenu" class="window-menu">
        <li
          v-for="(item, idx) in row.items"
          :key="idx"
          @click.stop="onMenuItemClick($event, item)"
          @mouseover="contextMenuOptions && onMenuItemClick($event, item)"
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
import { type Ref, inject } from 'vue'
import { AbsolutePosition } from '../models/absolute_position'
import type { SetContextMenu } from '../views/DesktopView.vue'
import type { ContextMenuOptions } from './ContextMenuComponent.vue'

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

const contextMenuOptions = inject<Ref<ContextMenuOptions | null>>('contextMenu')
const setContextMenu = inject('setContextMenu') as SetContextMenu

function onMenuItemClick(e: Event, item: WindowToolbarItem) {
  if (item.contextMenuOptions) {
    const { left, top, height } = (e.target as HTMLElement).getBoundingClientRect()
    const contextMenuPos = new AbsolutePosition(left, top + height)
    setContextMenu(item.contextMenuOptions ?? null, contextMenuPos)
  }
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';

nav.window-toolbar {
  color: colors.$black;
  display: flex;
  flex-direction: column;
  user-select: none;
  overflow: hidden;
  font-size: 0.9em;

  > ul.window-menu {
    margin: 0;
    padding-left: 0;
    display: flex;
    list-style-type: none;
  }
  > ul > li {
    padding: 0 7px;
    line-height: 24px;
    transition:
      background-color 0.1s,
      color 0.1s;
    &:hover {
      background-color: colors.$dark-grey;
      color: colors.$white;
    }
  }
}
</style>
