<template>
  <div
    ref="contextMenuEl"
    class="context-menu context-menu"
    :style="{
      left: pos.x + '%',
      top: pos.y + '%',
    }"
  >
    <template v-for="(itemGroup, idx) in options.itemGroups" :key="idx">
      <ul>
        <li v-for="(item, idx) in itemGroup" :key="idx" @click="item.onClick">
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
  pos: RelativePosition
  options: ContextMenuOptions
}>()

const contextMenuEl = ref()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.context-menu {
  z-index: 999;
  position: absolute;
  padding: 2px;
  border: 1px solid colors.$black;
  background-color: colors.$white;
  // box-shadow: 5px 5px 2px -3px rgba(0, 0, 0, 0.6);
  // -webkit-box-shadow: 5px 5px 2px -3px rgba(0, 0, 0, 0.6);
  // -moz-box-shadow: 5px 5px 2px -3px rgba(0, 0, 0, 0.6);
  min-width: 100px;
  user-select: none;

  hr {
    margin: 4px 2px;
    border-color: colors.$black;
    border-style: solid;
    border-bottom: 0;
  }

  > ul {
    margin: 0;
    display: flex;
    flex-direction: column;
    padding-left: 0;
    list-style-type: none;
  }
  > ul > li {
    padding: 2px 0;
    display: flex;
    align-items: center;
    justify-content: space-between;
    &:hover {
      background-color: colors.$black;
      color: colors.$white;
    }
    > span.context-menu__item-icon {
      width: 17px;
    }
    > span.context-menu__item-label {
      flex-grow: 1;
      margin-right: 15px;
    }
    > span.context-menu__item-suffix {
      text-align: end;
      margin-right: 15px;
    }
  }
}
</style>
