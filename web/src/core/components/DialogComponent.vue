<template>
  <WindowComponent
    :pos="pos"
    :size="size"
    :title="title"
    :controls="{
      close: true,
      minimize: false,
      maximize: false,
    }"
    :is-resizable="isResizable"
    @click-close="emit('clickClose')"
  >
    <div class="container">
      <slot />
      <div class="button-container">
        <button v-if="buttons.success" @click="emit('clickSuccess')">
          {{ typeof buttons.success === 'string' ? buttons.success : t('ui.success') }}
        </button>
        <button v-if="buttons.cancel" @click="emit('clickCancel')">
          {{ typeof buttons.cancel === 'string' ? buttons.cancel : t('ui.cancel') }}
        </button>
      </div>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { inject, onMounted, onUnmounted } from 'vue'
import { useI18n } from 'vue-i18n'
import type { RelativePosition } from '../models/relative_position'
import type { RelativeSize } from '../models/relative_size'
import WindowComponent, { type IBlockWindowFunc } from './WindowComponent.vue'

const emit = defineEmits<{
  (e: 'clickSuccess'): void
  (e: 'clickCancel'): void
  (e: 'clickClose'): void
}>()

const props = defineProps<{
  pos: RelativePosition
  size: RelativeSize
  title: string
  blocking?: boolean
  buttons: {
    success: boolean | string
    cancel: boolean | string
  }
  isResizable?: boolean
}>()

const { t } = useI18n()
const blockParentWindow = props.blocking
  ? (inject('blockParentWindow') as IBlockWindowFunc)
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
.container {
  height: 100%;
  padding: 1em;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  user-select: none;
}

.button-container {
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
