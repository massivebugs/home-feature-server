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
      <p class="system-about-dialog__container__links">
        <a href="https://github.com/massivebugs/home-feature-server" target="_blank">
          <GitHubIconComponent color="black" height="3em" />
        </a>
      </p>
      <p>
        <a href="https://github.com/massivebugs/home-feature-server" target="_blank">
          {{ t('systemSettings.about.gitHubRepoDescription') }}
        </a>
      </p>
      <p>
        {{ t('systemSettings.about.disclaimer') }}
      </p>
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'
import GitHubIconComponent from './GitHubIconComponent.vue'

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
    white-space: pre-line;
  }
}

.system-about-dialog__container__title {
  font-size: 1.2em;
}

.system-about-dialog__container__links {
  margin-top: 20px !important;
}
</style>
