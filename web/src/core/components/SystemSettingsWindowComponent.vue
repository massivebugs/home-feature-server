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
    :statusBarInfo="[t('app.name') + ' v' + appVersion]"
    :resizable="true"
    @click-close="emit('clickClose')"
    v-slot="{ windowSizeQuery }"
  >
    <div
      class="system-settings__container"
      :class="{
        'system-settings__container-md': windowSizeQuery.md,
        'system-settings__container-lg': windowSizeQuery.lg,
      }"
    >
      <SystemSettingsSectionComponent :title="t('systemSettings.preferences.title')">
        <form @change="onChangePreferences">
          <SystemSettingsItemComponent :name="t('systemSettings.preferences.language.title')">
            <template #icon>
              <LanguageIconComponent width="100%" height="100%" />
            </template>
            <SelectInputComponent
              class="system-settings__input"
              :options="localeOptions"
              v-model="store.systemPreference.language"
            />
          </SystemSettingsItemComponent>
        </form>
      </SystemSettingsSectionComponent>
    </div>
    <SystemAboutDialogComponent
      v-if="showAboutDialog"
      pos="center"
      @click-close="onClickCloseShowAboutDialog"
    />
  </WindowComponent>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import type { WindowToolbarRow } from '@/core/components/WindowToolbarComponent.vue'
import { RelativeSize } from '@/core/models/relativeSize'
import { Locales } from '@/i18n'
import type { API } from '../composables/useAPI'
import { useCoreStore } from '../stores'
import LanguageIconComponent from './LanguageIconComponent.vue'
import SystemAboutDialogComponent from './SystemAboutDialogComponent.vue'
import SystemSettingsItemComponent from './SystemSettingsItemComponent.vue'
import SystemSettingsSectionComponent from './SystemSettingsSectionComponent.vue'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const props = defineProps<{
  api: API
}>()

const { t, locale } = useI18n()
const store = useCoreStore()
const appVersion = APP_VERSION
const showAboutDialog = ref<boolean>(false)
const localeOptions: { label: string; value: any }[] = [
  { label: t('systemSettings.locale.default'), value: null },
  ...Object.values(Locales).map((v) => ({ label: t(`systemSettings.locale.${v}`), value: v })),
]
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
                label: t('systemSettings.about.linkTitle'),
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

const onChangePreferences = async () => {
  // TODO: API Update preferences
  locale.value = store.systemPreference.language ?? navigator.language
  await props.api.updateUserSystemPreference(store.systemPreference)
}

const onClickCloseShowAboutDialog = () => {
  showAboutDialog.value = false
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.system-settings__container {
  width: 100%;
  min-height: 100%;
  padding: 0.5em 0.5em;
  background-color: colors.$light-grey;
  display: flex;
  flex-direction: column;
}

.system-settings__container-md {
  padding: 1em 7em;
}

.system-settings__container-lg {
  padding: 1em 10em;
}

.system-settings__input {
  > * {
    padding: 0 1em;
  }
}
</style>
