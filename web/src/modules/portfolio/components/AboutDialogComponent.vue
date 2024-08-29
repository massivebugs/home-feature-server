<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :title="t('portfolio.aboutDialogTitle')"
    :buttons="{
      success: false,
      cancel: false,
    }"
    :blocking="true"
    :resizable="false"
  >
    <div class="container">Foobar</div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'

defineProps<{
  pos: RelativePosition
  size: RelativeSize
}>()

const { t } = useI18n()
</script>

<style scoped lang="scss">
.container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}
</style>
