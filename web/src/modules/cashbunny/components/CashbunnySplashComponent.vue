<template>
  <WindowComponent
    :pos="new RelativePosition(35, 35)"
    :size="new RelativeSize(30, 30)"
    :hide-titlebar="true"
  >
    <div class="container">
      <div class="title">
        <h1>Cashbunny</h1>
        <p>Best home budget planner ever</p>
      </div>
      <div class="footer">
        <p>2024 @massivebugs</p>
        <p class="loading">
          <span>Loading...</span><span>{{ loadPercent }}%</span>
        </p>
      </div>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import { RelativePosition } from '@/core/models/relative_position'
import { RelativeSize } from '@/core/models/relative_size'

const emit = defineEmits<{
  (e: 'loaded'): void
}>()

const loadPercent = ref<number>(0)

onMounted(() => {
  const intervalId = setInterval(() => {
    loadPercent.value = loadPercent.value + 1
    if (loadPercent.value === 100) {
      clearInterval(intervalId)
      setTimeout(() => {
        emit('loaded')
      }, 1500)
    }
  }, 10)
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.container {
  height: 100%;
  background-color: colors.$black;
  color: colors.$white;
  padding: 2em;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  user-select: none;

  .title {
    h1 {
      margin: 0;
    }
  }

  .loading {
    margin: 0;
    display: flex;
    justify-content: space-between;
  }
}
</style>
