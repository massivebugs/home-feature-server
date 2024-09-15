<template>
  <button class="hfs-button" :class="[type ? `hfs-button_${type}` : '']">
    <SpinnerIconComponent
      v-if="loadingSpinner"
      :type="loadingSpinner === true ? SpinnerTypes.ring : loadingSpinner"
    />
    <slot v-else />
  </button>
</template>

<script lang="ts">
export const ButtonTypes = {
  success: 'success',
} as const
export type ButtonType = (typeof ButtonTypes)[keyof typeof ButtonTypes]
</script>

<script setup lang="ts">
import type { SpinnerType } from './SpinnerIconComponent.vue'
import SpinnerIconComponent from './SpinnerIconComponent.vue'
import { SpinnerTypes } from './SpinnerIconComponent.vue'

defineProps<{
  type?: ButtonType
  loadingSpinner?: boolean | SpinnerType
}>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-button {
  border-radius: 5px;
  padding: 0.5em 0.3em;
  cursor: pointer;
  border: 1px solid colors.$light-grey;
  transition: background-color 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
  user-select: none;

  // Button types
  background-color: colors.$white;

  &:hover:not(:disabled) {
    background-color: darken(colors.$white, 10);
  }

  &:active:not(:disabled) {
    background-color: darken(colors.$white, 20);
  }

  // Success
  &.hfs-button_success {
    background-color: colors.$skobeloff;
    color: colors.$white;

    &:hover:not(:disabled) {
      background-color: darken(colors.$skobeloff, 10);
    }

    &:active:not(:disabled) {
      background-color: darken(colors.$skobeloff, 20);
    }
  }
}
</style>
