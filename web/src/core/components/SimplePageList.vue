<template>
  <ul class="hfs-simple-pagelist" :class="{ 'hfs-simple-pagelist_tabs': props.type === 'tabs' }">
    <li
      v-for="{ key, name } in props.pages"
      :key="key"
      @click="emit('clickPage', { key })"
      :class="{ 'hfs-simple-pagelist__current': key === props.currentKey }"
    >
      {{ name }}
    </li>
  </ul>
</template>

<script setup lang="ts">
export type SimplePageListClickPageEvent<T> = {
  key: T
}

const emit = defineEmits<{
  (e: 'clickPage', payload: SimplePageListClickPageEvent<any>): void
}>()

const props = defineProps<{
  pages: { key: string; name: string }[]
  type?: 'list' | 'tabs'
  currentKey?: string
}>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-simple-pagelist {
  border: 1px solid colors.$light-grey;
  border-radius: 5px;
  list-style-type: none;
  background-color: white;
  margin-top: 0;
  margin-bottom: 0;
  padding: 0.5em 0;
  display: flex;
  flex-direction: column;
  user-select: none;

  > li {
    &:first-child {
      border-top: 1px solid colors.$light-grey;
    }
    border-bottom: 1px solid colors.$light-grey;
    padding: 0.1em 1em;
  }

  > li:active {
    background-color: colors.$dark-grey;
    color: colors.$white;
  }
}

.hfs-simple-pagelist_tabs {
  flex-direction: row;
  justify-content: center;
  background: none;
  border: 0;

  > li {
    border: 1px solid colors.$light-grey;
    background-color: white;
  }

  > li:first-child {
    border-radius: 10px 0 0 10px;
  }

  > li:last-child {
    border-radius: 0 10px 10px 0;
  }
}

.hfs-simple-pagelist__current {
  background-color: colors.$skobeloff !important;
  color: colors.$white;
}
</style>
