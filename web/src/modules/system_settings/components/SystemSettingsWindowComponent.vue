<template>
  <WindowComponent
    :size="new RelativeSize(70, 80)"
    :title="t('systemSettings.name')"
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
    <div class="hfs-system-settings__container"></div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import WindowComponent from '@/core/components/WindowComponent.vue'
import type { WindowToolbarRow } from '@/core/components/WindowToolbarComponent.vue'
import { RelativeSize } from '@/core/models/relativeSize'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const showContactFormDialog = ref<boolean>(false)
const showAboutDialog = ref<boolean>(false)
const toolbarOptions = computed<WindowToolbarRow[]>(() => [
  {
    isMenu: true,
    items: [
      {
        label: t('systemSettings.file'),
        contextMenuOptions: {
          itemGroups: [
            [
              {
                label: t('systemSettings.exit'),
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
        label: t('systemSettings.help'),
        contextMenuOptions: {
          itemGroups: [
            [
              {
                label: t('systemSettings.about'),
                isDisabled: false,
                onClick: () => {
                  showAboutDialog.value = true
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

.hfs-system-settings__container {
  width: 100%;
  height: 100%;
  padding: 5px;
  background-color: colors.$light-grey;
}
</style>
