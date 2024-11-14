<template>
  <RouterView />
</template>

<script setup lang="ts">
import { useCoreStore } from './core/stores'
import { programList } from './programList'

const coreStore = useCoreStore()

// Register all programs to system
programList.forEach((program) => {
  coreStore.addProgram(program)
})
</script>

<style scoped lang="scss"></style>
