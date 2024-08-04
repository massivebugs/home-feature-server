<template>
  <ErrorDialogComponent
    v-if="isProcessAlreadyExists"
    :pos="new RelativePosition(40, 40)"
    :size="new RelativeSize(20, 20)"
    :title="t('cashbunny.processAlreadyExistsTitle')"
    :message="t('cashbunny.processAlreadyExistsMessage')"
  />
  <CashbunnySplashComponent v-else-if="!isLoadedInitialData" @loaded="onLoadedData" />
  <CashbunnyWindowComponent v-else />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ErrorDialogComponent from '@/core/components/ErrorDialogComponent.vue'
import { RelativePosition } from '@/core/models/relative_position'
import { RelativeSize } from '@/core/models/relative_size'
import { useCoreStore } from '@/core/stores'
import CashbunnySplashComponent from '../components/CashbunnySplashComponent.vue'
import CashbunnyWindowComponent from '../components/CashbunnyWindowComponent.vue'
import { CASHBUNNY_PROGRAM_ID } from '../constants'

const { t } = useI18n()
const coreStore = useCoreStore()
const isProcessAlreadyExists = ref<boolean>(false)
const isLoadedInitialData = ref<boolean>(false)

onMounted(() => {
  if (coreStore.findProgramProcesses(CASHBUNNY_PROGRAM_ID).length > 1) {
    isProcessAlreadyExists.value = true
  }
})

const onLoadedData = () => {
  isLoadedInitialData.value = true
}
</script>

<style scoped lang="scss"></style>
