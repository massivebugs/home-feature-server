<template>
  <div
    ref="contextMenuEl"
    class="twc-context-menu context-menu"
    :style="{
      left: relPos.x + '%',
      top: relPos.y + '%',
    }"
  >
    <template v-for="(itemGroup, idx) in options.itemGroups" :key="idx">
      <ul>
        <li v-for="(item, idx) in itemGroup" :key="idx">
          <span class="context-menu__item-icon"> </span>
          <span class="context-menu__item-label">
            {{ item.label }}
          </span>
          <span class="context-menu__item-suffix">
            {{ item.shortcutKey }}
          </span>
        </li>
      </ul>
      <hr v-if="options.itemGroups.length > 1 && idx < options.itemGroups.length - 1" />
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { RelativePosition } from '../models/relative_position'

export type ContextMenuItem = {
  icon?: string
  label: string
  shortcutKey?: string
  isDisabled?: boolean
  onClick?: () => void
  children?: ContextMenuItem[][]
}

export type ContextMenuOptions = {
  itemGroups: ContextMenuItem[][] // Groups are separated with a spacer
}

defineProps<{
  relPos: RelativePosition
  options: ContextMenuOptions
}>()

const contextMenuEl = ref()
</script>

<style scoped lang="scss">
@use 'xp.css/dist/XP.css';
@use '@/assets/xp.custom';

.twc-context-menu {
  z-index: 999;
  position: absolute;
}
</style>
