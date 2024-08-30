<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :title="t('portfolio.contact')"
    :buttons="{
      success: t('portfolio.sendContactMessage'),
      cancel: false,
    }"
    :blocking="true"
    :resizable="true"
    @click-success="emit('submit', { subject: 'TODO', email: 'TODO', message: 'TODO' })"
  >
    <div class="container"></div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'

export type ContactFormSuccessEvent = {
  subject: string
  email: string
  message: string
}

const emit = defineEmits<{
  (e: 'submit', payload: ContactFormSuccessEvent): void
}>()

defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
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

.title {
  text-align: center;
  font-weight: 500;
}
</style>
