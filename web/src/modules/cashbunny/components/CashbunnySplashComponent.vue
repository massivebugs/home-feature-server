<template>
  <WindowComponent
    class="cashbunny-splash-window"
    pos="center"
    :hide-titlebar="true"
    :static="true"
  >
    <div class="cashbunny-splash__container">
      <CashbunnyIconComponent class="cashbunny-splash__image" />
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
import type { API } from '@/core/composables/useAPI'
import { sleep } from '@/core/utils/time'
import { useCashbunnyStore } from '../stores'
import CashbunnyIconComponent from './CashbunnyIconComponent.vue'

const emit = defineEmits<{
  (e: 'loaded'): void
  (e: 'error', message: any): void
}>()

const props = defineProps<{
  api: API
}>()

const cashbunnyStore = useCashbunnyStore()
const loadMessage = ref<string>('Initializing Cashbunny...')
const loadPercent = ref<number>(0)

const retrieveUserPreferences = async () => {
  loadMessage.value = 'Retrieving user preferences...'
  try {
    const res = await props.api.getCashbunnyUserPreference({ 404: () => {} })
    cashbunnyStore.userPreference = res.userPreference
  } catch (e) {
    if (!props.api.isError(e)) {
      throw e
    }
  }
}

const createUserPreferences = async () => {
  loadMessage.value = 'Setting user preferences...'
  if (!cashbunnyStore.userPreference) {
    const res = await props.api.createCashbunnyDefaultUserPreference()
    cashbunnyStore.userPreference = res.userPreference
  }
}

const retrieveAllCurrencies = async () => {
  loadMessage.value = 'Retrieving list of supported currencies...'
  const res = await props.api.getCashbunnySupportedCurrencies()
  cashbunnyStore.setCurrencies(res)
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
@use '@/assets/media-query';
@use '@/assets/colors';

.cashbunny-splash-window {
  // TODO: Ugly overrides, make a more streamlined way for window sizing
  min-width: 300px;
  height: auto !important;
  width: 90%;

  @include media-query.md {
    width: 70%;
  }

  @include media-query.lg {
    width: 500px;
  }
}

.cashbunny-splash__container {
  position: relative;
  height: 100%;
  user-select: none;
  overflow: hidden;
  border-radius: 50px;
  background-color: colors.$viridian;

  .cashbunny-splash__image {
    position: absolute;
    width: 100%;
    height: 100%;
    transform: scale(1.7) translate(9%, 3%);
    fill: colors.$white;
  }

  &::after {
    background-color: colors.$rich-black;
    opacity: 0.6;
    border-radius: 50px 100% 50px 50px;

    position: absolute;
    content: '';
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }

  .cashbunny-splash__contents {
    position: relative;
    width: 100%;
    height: 100%;
    color: colors.$white;
    padding: 2em;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    z-index: 1;
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
