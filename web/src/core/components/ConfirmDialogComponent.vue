<template>
  <WindowComponent
    :pos="pos"
    :size="props.size"
    :title="props.title"
    :controls="{
      close: true,
      minimize: false,
      maximize: false,
    }"
    @click-close="emit('clickClose')"
  >
    <div class="container">
      <p class="message">{{ message }}</p>
      <div class="button-container">
        <button @click="emit('clickSuccess')">
          {{ t('ui.success') }}
        </button>
        <button @click="emit('clickCancel')">
          {{ t('ui.cancel') }}
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
  message: string
}>()

const { t } = useI18n()
const blockParentWindow = inject('blockParentWindow') as IBlockWindowFunc

onMounted(() => {
  blockParentWindow(true)
})

onUnmounted(() => {
  blockParentWindow(false)
})
</script>

<style scoped lang="scss">
.container {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  user-select: none;
}

.message {
  text-align: center;
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
