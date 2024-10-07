<template>
  <WindowComponent
    :size="new RelativeSize(70, 80)"
    :title="t('telebyte.title')"
    :controls="{
      minimize: true,
      maximize: true,
      close: true,
    }"
    :toolbar="toolbarOptions"
    :statusBarInfo="['Something goes here...', 'Something else here']"
    :resizable="true"
    @click-close="emit('clickClose')"
  >
    <div class="telebyte-window__container">Telebyte!</div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import WindowComponent from '@/core/components/WindowComponent.vue'
import type { WindowToolbarRow } from '@/core/components/WindowToolbarComponent.vue'
import { RelativeSize } from '@/core/models/relativeSize'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const toolbarOptions = computed<WindowToolbarRow[]>(() => [
  {
    isMenu: true,
    items: [
      {
        label: t('common.file'),
        contextMenuOptions: {
          itemGroups: [
            [
              {
                label: t('common.exit'),
                shortcutKey: 'Alt+F4',
                isDisabled: false,
                onClick: () => {
                  emit('clickClose')
                },
              },
            ],
          ],
        },
      },
      {
        label: t('common.help'),
        contextMenuOptions: {
          itemGroups: [
            [
              {
                label: t('common.about', { v: 'Telebyte' }),
                isDisabled: false,
                onClick: () => {
                  //
                },
              },
            ],
          ],
        },
      },
    ],
  },
])
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.telebyte-window__container {
  width: 100%;
  height: 100%;
  padding: 5px;
}
</style>
