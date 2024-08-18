<template>
  <WindowComponent
    :pos="new RelativePosition(35, 35)"
    :size="new RelativeSize(30, 30)"
    :hide-titlebar="true"
  >
    <div class="cashbunny-splash__container">
      <div class="cashbunny-splash__background">
        <CashbunnyIconComponent class="cashbunny-splash__image-back" />
        <div>
          <CashbunnyIconComponent class="cashbunny-splash__image-front" />
        </div>
      </div>
      <div class="cashbunny-splash__contents">
        <div class="cashbunny-splash__title">
          <h1>Cashbunny</h1>
          <p>Best home budget planner ever</p>
        </div>
        <div class="cashbunny-splash__footer">
          <p>2024 @massivebugs</p>
          <p class="cashbunny-splash__loading-message">
            <span>{{ loadMessage }}</span
            ><span>{{ loadPercent }}%</span>
          </p>
        </div>
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
import CashbunnyIconComponent from './CashbunnyIconComponent.vue'

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
  // emit('loaded')
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.cashbunny-splash__container {
  position: relative;
  height: 100%;
  user-select: none;
  overflow: hidden;

  .cashbunny-splash__background {
    position: absolute;
    width: 100%;
    height: 100%;
    border-radius: 50px;
    background-color: colors.$viridian;

    .cashbunny-splash__image-back {
      position: absolute;
      width: 100%;
      height: 100%;
      transform: scale(1.7) translate(9%, 3%);
      fill: colors.$white;
    }

    :last-child {
      overflow: hidden;
      background-color: colors.$rich-black;
      border-radius: 50px 100% 50px 50px;

      position: absolute;
      width: 100%;
      height: 100%;

      .cashbunny-splash__image-front {
        width: 100%;
        height: 100%;
        transform: scale(1.7) translate(9%, 3%);
        fill: colors.$dark-grey;
      }
    }
  }

  .cashbunny-splash__contents {
    width: 100%;
    height: 100%;
    color: colors.$white;
    padding: 2em;
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    > * {
      z-index: 1;
    }
  }

  .cashbunny-splash__title {
    h1 {
      margin: 0;
    }
  }

  .cashbunny-splash__loading-message {
    margin: 0;
    display: flex;
    justify-content: space-between;
  }
}
</style>
