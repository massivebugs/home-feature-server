<template>
  <ErrorDialogComponent
    v-if="errorTitle && errorMessage"
    pos="center"
    :title="errorTitle"
    :message="errorMessage"
  />
  <CashbunnySplashComponent
    v-else-if="!isLoadedInitialData"
    :api="api"
    @loaded="onLoadedData"
    @error="onErrorLoadingData"
  />
  <CashbunnyWindowComponent v-else />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ErrorDialogComponent from '@/core/components/ErrorDialogComponent.vue'
import { useAPI } from '@/core/composables/useAPI'
import { API_URL } from '@/core/constants'
import { useCoreStore } from '@/core/stores'
import CashbunnySplashComponent from '../components/CashbunnySplashComponent.vue'
import CashbunnyWindowComponent from '../components/CashbunnyWindowComponent.vue'
import { CASHBUNNY_PROGRAM_ID } from '../constants'

const { t } = useI18n()
const api = useAPI(API_URL)
const coreStore = useCoreStore()
const errorTitle = ref<string | null>(null)
const errorMessage = ref<string | null>(null)
const isLoadedInitialData = ref<boolean>(false)

onMounted(() => {
  if (coreStore.findProgramProcesses(CASHBUNNY_PROGRAM_ID).length > 1) {
    errorTitle.value = t('cashbunny.processAlreadyExistsTitle')
    errorMessage.value = t('cashbunny.processAlreadyExistsMessage')
  }
})

const onLoadedData = () => {
  isLoadedInitialData.value = true
}

const onErrorLoadingData = (e: any) => {
  errorTitle.value = t('cashbunny.errorLoadingDataTitle')
  errorMessage.value = t('cashbunny.errorLoadingDataMessage', { e })
}
</script>

<style scoped lang="scss"></style>
