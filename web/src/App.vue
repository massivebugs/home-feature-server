<template>
  <RouterView />
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { useStore } from './core/stores'
import { getPrograms } from './programs'
import { Process } from './core/models/process'
import { uniqueId } from 'lodash'

const { t } = useI18n()
const store = useStore()

// Register all programs to system
const programs = getPrograms(store, t)
programs.forEach((program) => {
  store.addProgram(program)
})

// TODO: For now, run cashbunny by default
const cashbunnyProgram = programs[0]
const cashbunnyProcess = new Process(uniqueId('pid_'), cashbunnyProgram)
store.addProcess(cashbunnyProcess)
</script>

<style scoped lang="scss"></style>
