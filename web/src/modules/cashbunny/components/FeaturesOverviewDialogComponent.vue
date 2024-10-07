<template>
  <DialogComponent
    :pos="pos"
    :size="size"
    :title="t('cashbunny.featuresOverview.name')"
    :buttons="{
      success: false,
      cancel: false,
    }"
    :controls="{
      minimize: false,
      maximize: true,
      close: true,
    }"
    :blocking="true"
    :resizable="true"
  >
    <FeaturesOverviewComponent />
  </DialogComponent>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'
import FeaturesOverviewComponent from './FeaturesOverviewComponent.vue'

defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
}>()

const { t } = useI18n()
</script>
