<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :title="t('systemSettings.about.linkTitle')"
    :buttons="{
      success: false,
      cancel: false,
    }"
    :blocking="true"
    :resizable="true"
  >
    <div class="system-about-dialog__container">
      <p class="system-about-dialog__container__title">{{ t('app.name') }}</p>
      <p>v{{ appVersion }}</p>
      <p>https://github.com/massivebugs/home-feature-server</p>
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'

defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
}>()

const appVersion = APP_VERSION

const { t } = useI18n()
</script>

<style scoped lang="scss">
.system-about-dialog__container {
  width: 100%;
  height: 100%;

  > p {
    text-align: center;
    margin: 10px 0;
  }
}

.system-about-dialog__container__title {
  font-size: 1.2em;
}
</style>
