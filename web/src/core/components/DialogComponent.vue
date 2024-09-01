<template>
  <WindowComponent
    class="hfs-dialog"
    :pos="pos"
    :size="size"
    :title="title"
    :controls="
      controls ?? {
        close: true,
        minimize: false,
        maximize: false,
      }
    "
    :resizable="resizable"
    :fit-content="fitContent"
    @click-close="emit('clickClose')"
  >
    <template #title>
      <slot name="title" />
    </template>
    <div class="hfs-dialog__container">
      <slot />
      <div class="hfs-dialog__buttons">
        <ButtonComponent
          v-if="buttons.success"
          type="success"
          :disabled="disabled"
          :loading-spinner="loadingSpinner"
          @click="emit('clickSuccess')"
        >
          {{ typeof buttons.success === 'string' ? buttons.success : t('ui.success') }}
        </ButtonComponent>
        <ButtonComponent
          v-if="buttons.cancel"
          :disabled="disabled"
          :loading-spinner="loadingSpinner"
          @click="emit('clickCancel')"
        >
          {{ typeof buttons.cancel === 'string' ? buttons.cancel : t('ui.cancel') }}
        </ButtonComponent>
      </div>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { inject, onMounted, onUnmounted } from 'vue'
import { useI18n } from 'vue-i18n'
import type { RelativePosition } from '../models/relativePosition'
import { RelativeSize } from '../models/relativeSize'
import ButtonComponent from './ButtonComponent.vue'
import type { SpinnerType } from './SpinnerIconComponent.vue'
import WindowComponent, {
  type BlockWindowFunc,
  type WindowTitleBarControls,
} from './WindowComponent.vue'

const emit = defineEmits<{
  (e: 'clickSuccess'): void
  (e: 'clickCancel'): void
  (e: 'clickClose'): void
}>()

const props = defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
  fitContent?: boolean
  title?: string
  blocking?: boolean
  buttons: {
    success: boolean | string
    cancel: boolean | string
  }
  disabled?: boolean
  controls?: WindowTitleBarControls
  resizable?: boolean
  loadingSpinner?: SpinnerType
}>()

const { t } = useI18n()
const blockParentWindow = props.blocking
  ? (inject('blockParentWindow') as BlockWindowFunc)
  : undefined

onMounted(() => {
  if (props.blocking && blockParentWindow) {
    blockParentWindow(true)
  }
})

onUnmounted(() => {
  if (props.blocking && blockParentWindow) {
    blockParentWindow(false)
  }
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-dialog__container {
  height: 100%;
  padding: 1em;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  user-select: none;
  background-color: colors.$white;
}

.hfs-dialog__buttons {
  width: 80%;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 5px;
  > button {
    flex: 1;
  }
}
</style>
