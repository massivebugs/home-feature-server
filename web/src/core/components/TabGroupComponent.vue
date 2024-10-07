<template>
  <menu class="hfs-tab-group">
    <button
      role="tab"
      v-for="tab in tabs"
      :key="tab.id"
      :aria-selected="isCurrentTab(tab.id)"
      :aria-controls="tab.id"
      class="hfs-tab-group__tab"
      :class="{
        'hfs-tab-group__tab-selected': isCurrentTab(tab.id),
      }"
      @click="emit('tabClick', { tabId: tab.id })"
    >
      <slot :name="tab.id">
        {{ tab.label }}
      </slot>
    </button>
  </menu>
</template>

<script setup lang="ts">
export type TabClickEvent = {
  tabId: string
}

export type TabInfo = {
  id: string // An identifier for the tab. If this is passed to 'selectedTabId' in props, that tab will be shown
  label?: string // This is shown on the 'select' buttons of the tab group
}

const emit = defineEmits<{
  (e: 'tabClick', value: TabClickEvent): void
}>()

const props = defineProps<{
  tabs: TabInfo[]
  selectedTabId: string
}>()

const isCurrentTab = (tabId: string) => {
  return tabId === props.selectedTabId
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';
.hfs-tab-group {
  display: flex;
  padding: 0;
  margin: 0;
  user-select: none;
  flex-wrap: wrap;
}

.hfs-tab-group__tab {
  padding: 0.3em 0.5em;
  display: flex;
  align-items: center;
  gap: 5px;
  font-weight: bold;
}

.hfs-tab-group__tab-selected {
  background-color: colors.$dark-grey;
  color: colors.$white;
  border-top: 3px solid colors.$white;
}
</style>
