<template>
  <DialogComponent
    :pos="props.pos"
    :size="props.size"
    :title="props.title"
    :buttons="{
      success: props.buttons?.success ?? true,
      cancel: props.buttons?.cancel ?? true,
    }"
    :blocking="props.blocking"
  >
    <p class="hfs-confirm-dialog__message">{{ props.message }}</p>
  </DialogComponent>
</template>

<script setup lang="ts">
import type { RelativePosition } from '../models/relativePosition'
import type { RelativeSize } from '../models/relativeSize'
import DialogComponent from './DialogComponent.vue'

const props = defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
  title: string
  message: string
  blocking?: boolean
  buttons?: {
    success: string
    cancel: string
  }
}>()
</script>

<style scoped lang="scss">
.hfs-confirm-dialog__message {
  text-align: center;
  margin: 0;
  padding-bottom: 2em;
  white-space: pre-line;
}
</style>
