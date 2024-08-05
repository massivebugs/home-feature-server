<template>
  <section class="tabs">
    <menu>
      <button
        role="tab"
        v-for="tab in tabs"
        :key="tab.id"
        :aria-selected="isCurrentTab(tab.id)"
        :aria-controls="tab.id"
        @click="emit('tabClick', { tabId: tab.id })"
      >
        <slot :name="tab.id + '_label'">
          {{ tab.label }}
        </slot>
      </button>
    </menu>
    <article>
      <slot :name="selectedTabId" />
    </article>
  </section>
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
menu {
  padding: 0;
  margin: 0;
}
</style>
