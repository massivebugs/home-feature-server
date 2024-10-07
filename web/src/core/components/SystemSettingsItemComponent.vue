<template>
  <li class="system-settings__item">
    <div class="sytem-settings__item-icon">
      <slot name="icon" />
    </div>
    <div class="sytem-settings__item-name">{{ props.name }}</div>
    <div>
      <slot />
    </div>
  </li>
</template>
<script setup lang="ts">
const props = defineProps<{
  name: string
}>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.system-settings__item {
  background-color: colors.$low-opacity-light-grey;
  border-radius: 5px;
  padding: 0.3em 1.3em;
  display: flex;
  gap: 0.4em;
  user-select: none;

  > div {
    display: flex;
    align-items: center;
  }
}

.sytem-settings__item-icon {
  width: 1em;
}

.sytem-settings__item-name {
  flex: 1;
}
</style>
