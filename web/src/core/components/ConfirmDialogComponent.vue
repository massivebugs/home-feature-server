<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :title="title"
    :buttons="{
      success: true,
      cancel: true,
    }"
    :blocking="blocking"
  >
    <p class="message">{{ message }}</p>
  </DialogComponent>
</template>

<script setup lang="ts">
import type { RelativePosition } from '../models/relativePosition'
import type { RelativeSize } from '../models/relativeSize'
import DialogComponent from './DialogComponent.vue'

defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
  title: string
  message: string
  blocking?: boolean
}>()
</script>

<style scoped lang="scss">
.message {
  text-align: center;
  margin: 0;
  padding-bottom: 2em;
  white-space: pre-line;
}
</style>
