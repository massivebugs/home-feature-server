<template>
  <RouterView />
</template>

<script setup lang="ts">
import { useCoreStore } from './core/stores'
import { getPrograms } from './programs'

const store = useCoreStore()

// Register all programs to system
const programs = getPrograms()
programs.forEach((program) => {
  store.addProgram(program)
})
</script>

<style scoped lang="scss"></style>
