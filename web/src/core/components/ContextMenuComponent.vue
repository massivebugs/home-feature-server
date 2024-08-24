<template>
  <div
    ref="contextMenuEl"
    class="hfs-context-menu"
    :style="{
      left: pos.x + '%',
      top: pos.y + '%',
    }"
  >
    <template v-for="(itemGroup, idx) in options.itemGroups" :key="idx">
      <ul>
        <template v-for="(item, idx) in itemGroup" :key="idx">
          <li v-if="item" @click="item.onClick">
            <span class="hfs-context-menu__item-icon"> </span>
            <span class="hfs-context-menu__item-label">
              {{ item.label }}
            </span>
            <span class="hfs-context-menu__item-suffix">
              {{ item.shortcutKey }}
            </span>
          </li>
        </template>
      </ul>
      <hr v-if="options.itemGroups.length > 1 && idx < options.itemGroups.length - 1" />
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { RelativePosition } from '../models/relativePosition'

export type ContextMenuItem = {
  icon?: string
  label: string
  shortcutKey?: string
  isDisabled?: boolean
  onClick?: () => void
  children?: ContextMenuItem[][]
}

export type ContextMenuOptions = {
  itemGroups: (ContextMenuItem | undefined)[][] // Groups are separated with a spacer
}

defineProps<{
  pos: RelativePosition
  options: ContextMenuOptions
}>()

const contextMenuEl = ref()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-context-menu {
  z-index: 999;
  position: absolute;
  padding: 2px;
  min-width: 100px;
  user-select: none;
  border-radius: 0 5px 5px 5px;
  background-color: colors.$high-opacity-white;
  @supports (backdrop-filter: blur()) {
    backdrop-filter: blur(10px);
    background-color: colors.$high-opacity-white;
  }
  box-shadow: 1px 1px 4px 1px rgba(0, 0, 0, 0.4);
  -webkit-box-shadow: 1px 1px 4px 1px rgba(0, 0, 0, 0.4);
  -moz-box-shadow: 1px 1px 4px 1px rgba(0, 0, 0, 0.4);

  hr {
    margin: 4px 2px;
    border: 1px solid colors.$light-grey;
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
    font-size: 0.9em;
    padding: 2px 0;
    display: flex;
    align-items: center;
    justify-content: space-between;
    &:hover {
      background-color: colors.$dark-grey;
      color: colors.$white;
    }
    > span.hfs-context-menu__item-icon {
      width: 17px;
    }
    > span.hfs-context-menu__item-label {
      flex-grow: 1;
      margin-right: 15px;
    }
    > span.hfs-context-menu__item-suffix {
      text-align: end;
      margin-right: 15px;
    }
  }
}
</style>
