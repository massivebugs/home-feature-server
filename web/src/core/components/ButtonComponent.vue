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
  primary: 'primary',
  success: 'success',
  info: 'info',
  warning: 'warning',
  danger: 'danger',
  dark: 'dark',
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
@use 'sass:color';
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

  &:disabled {
    background-color: colors.$light-grey !important;
    color: colors.$dark-grey !important;
  }

  &:hover:not(:disabled) {
    background-color: darken(colors.$white, 10);
  }

  &:active:not(:disabled) {
    background-color: darken(colors.$white, 20);
  }

  @mixin button-styles($bg-color) {
    background-color: $bg-color;

    &:hover:not(:disabled) {
      background-color: darken($bg-color, 9);
    }

    &:active:not(:disabled) {
      background-color: darken($bg-color, 19);
    }
  }

  &.hfs-button_primary {
    @include button-styles(colors.$deep-sea-blue);
    color: colors.$white;
  }

  &.hfs-button_success {
    @include button-styles(colors.$skobeloff);
    color: colors.$white;
  }

  &.hfs-button_info {
    @include button-styles(colors.$viridian);
    color: colors.$white;
  }

  &.hfs-button_warning {
    @include button-styles(colors.$peach);
  }

  &.hfs-button_danger {
    @include button-styles(color.scale(colors.$red-cmyk, $saturation: -10%));
    color: colors.$white;
  }

  &.hfs-button_dark {
    @include button-styles(colors.$rich-black);
    color: colors.$white;
  }
}
</style>
