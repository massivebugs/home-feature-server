<template>
  <RouterView />
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { useCoreStore } from './core/stores'
import { getPrograms } from './programs'

const { t } = useI18n()
const store = useCoreStore()

// Register all programs to system
const programs = getPrograms(t)
programs.forEach((program) => {
  store.addProgram(program)
})
</script>

<style scoped lang="scss"></style>
