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
          <span>{{ loadMessage }}</span
          ><span>{{ loadPercent }}%</span>
        </p>
      </div>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import type { APIResponse } from '@/core/models/dto'
import { RelativePosition } from '@/core/models/relative_position'
import { RelativeSize } from '@/core/models/relative_size'
import { sleep } from '@/core/utils/time'
import { isAPIError } from '@/utils/api'
import { useCashbunnyStore } from '../stores'

const emit = defineEmits<{
  (e: 'loaded'): void
  (e: 'error', message: any): void
}>()

const store = useCashbunnyStore()
const loadMessage = ref<string>('Initializing Cashbunny...')
const loadPercent = ref<number>(0)

const retrieveUserPreferences = async () => {
  loadMessage.value = 'Retrieving user preferences...'
  try {
    const res = await store.getUserPreferences()
    store.userPreferences = res.data.data
  } catch (e) {
    if (isAPIError(e)) {
      const res = e.response?.data as APIResponse<null>
      // TODO: Group API error codes
      if (res.error?.code !== 'not_found') {
        throw e
      }
    } else {
      throw e
    }
  }
}

const createUserPreferences = async () => {
  loadMessage.value = 'Setting user preferences...'
  if (!store.userPreferences) {
    const res = await store.createUserPreferences()
    store.userPreferences = res.data.data
  }
}

const retrieveAllCurrencies = async () => {
  loadMessage.value = 'Retrieving list of supported currencies...'
  const res = await store.getAllCurrencies()
  if (res.data.error === null) {
    store.setCurrencies(res.data.data)
  }
}

onMounted(async () => {
  const loadProcesses: (() => Promise<void>)[] = [
    retrieveUserPreferences,
    createUserPreferences,
    retrieveAllCurrencies,
  ]

  for (let i = 0; i < loadProcesses.length; i++) {
    try {
      await loadProcesses[i]()

      while (loadPercent.value < (100 / loadProcesses.length) * (i + 1)) {
        loadPercent.value = loadPercent.value + 1
        await sleep(10)
      }
    } catch (e) {
      emit('error', e)
      return
    }
  }
  emit('loaded')
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
