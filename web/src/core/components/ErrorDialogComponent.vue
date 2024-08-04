<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :title="title"
    :buttons="{
      success: false,
      cancel: true,
    }"
    :blocking="blocking"
  >
    <p class="message">{{ message }}</p>
  </DialogComponent>
</template>

<script setup lang="ts">
import type { RelativePosition } from '../models/relative_position'
import type { RelativeSize } from '../models/relative_size'
import DialogComponent from './DialogComponent.vue'

defineProps<{
  pos: RelativePosition
  size: RelativeSize
  title: string
  message: string
  blocking?: boolean
}>()
</script>

<style scoped lang="scss">
.message {
  text-align: center;
  margin: 0;
}
</style>
