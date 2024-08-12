<template>
  <div ref="calendar"></div>
</template>

<script setup lang="ts">
import VC from 'vanilla-calendar-pro'
import type { IOptions } from 'vanilla-calendar-pro/types'
import { onMounted, ref } from 'vue'

const props = defineProps<{
  config?: IOptions
}>()

const calendar = ref()
const vanillaCalendar = ref<VC | null>(null)

onMounted(() => {
  vanillaCalendar.value = new VC(calendar.value, props.config)
  vanillaCalendar.value.init()
})
</script>

<style lang="scss">
@import 'vanilla-calendar-pro/build/vanilla-calendar.min.css';
</style>
