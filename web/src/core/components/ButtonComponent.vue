<template>
  <button class="hfs-button">
    <SpinnerIconComponent
      v-if="loadingSpinner"
      :type="loadingSpinner === true ? SpinnerTypes.ring : loadingSpinner"
    />
    <slot v-else />
  </button>
</template>
<script setup lang="ts">
import type { SpinnerType } from './SpinnerIconComponent.vue'
import SpinnerIconComponent from './SpinnerIconComponent.vue'
import { SpinnerTypes } from './SpinnerIconComponent.vue'

defineProps<{
  loadingSpinner?: boolean | SpinnerType
}>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-button {
  border-radius: 5px;
  padding: 0.5em 0.3em;
  cursor: pointer;
  background-color: colors.$white;
  border: 1px solid colors.$light-grey;
  transition: background-color 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover:not(:disabled) {
    background-color: darken(colors.$white, 10);
  }

  &:active:not(:disabled) {
    background-color: darken(colors.$white, 20);
  }
}
</style>
