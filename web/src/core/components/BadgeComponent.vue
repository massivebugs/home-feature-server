<template>
  <span class="hfs-badge" :class="[colorClass]">{{ props.name }}</span>
</template>

<script lang="ts">
export const BadgeTypes: { [key: string]: string } = {
  primary: 'primary',
  secondary: 'secondary',
  success: 'success',
  info: 'info',
  warning: 'warning',
  danger: 'danger',
  dark: 'dark',
} as const
export type BadgeType = (typeof BadgeTypes)[keyof typeof BadgeTypes]
</script>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  name: string
  color?: BadgeType | string
}>()

const colorClass = computed<string>(() => {
  if (props.color) {
    if (props.color in BadgeTypes) {
      return `hfs-badge-${BadgeTypes[props.color]}`
    } else {
      return props.color
    }
  }
  return 'hfs-badge-secondary'
})
</script>

<style scoped lang="scss">
@use 'sass:color';
@use '@/assets/colors';

.hfs-badge {
  border-radius: 50px;
  padding: 0.3em 0.5em;
  font-size: 0.9em;
  user-select: none;
}

@mixin badge-styles($bg-color) {
  background-color: $bg-color;
}

.hfs-badge-primary {
  @include badge-styles(colors.$deep-sea-blue);
  color: colors.$white;
}

.hfs-badge-secondary {
  @include badge-styles(colors.$light-grey);
}

.hfs-badge-success {
  @include badge-styles(colors.$skobeloff);
  color: colors.$white;
}

.hfs-badge-info {
  @include badge-styles(colors.$viridian);
  color: colors.$white;
}

.hfs-badge-warning {
  @include badge-styles(colors.$peach);
}

.hfs-badge-danger {
  @include badge-styles(color.scale(colors.$red-cmyk, $saturation: -10%));
  color: colors.$white;
}

.hfs-badge-dark {
  @include badge-styles(colors.$rich-black);
  color: colors.$white;
}
</style>
