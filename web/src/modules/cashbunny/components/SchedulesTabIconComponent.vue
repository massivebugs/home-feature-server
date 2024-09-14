<template>
  <svg
    width="30"
    height="30"
    viewBox="0 0 48 48"
    stroke-width="2"
    xmlns="http://www.w3.org/2000/svg"
  >
    <text :x="textX" y="32" font-size="20" font-weight="500" :fill="color">{{ day }}</text>

    <path
      :stroke="color ?? '#000'"
      fill="none"
      d="m8.0988,7.8588c-1.0601,0-1.9194.8593-1.9194,1.9194v31.8025c0,1.06.8593,1.9193,1.9193,1.9193h31.8025c1.06,0,1.9193-.8592,1.9194-1.9192h0V9.7782c0-1.0601-.8593-1.9194-1.9194-1.9194H8.0988Z"
    />

    <line :stroke="color ?? '#000'" fill="none" x1="12.8166" y1="10.258" x2="12.8166" y2="4.5" />

    <line :stroke="color ?? '#000'" fill="none" x1="35.1834" y1="10.258" x2="35.1834" y2="4.5" />
  </svg>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{ color?: string; day: number }>()

// Use pre-defined x position based on day's number
const textX = computed(() => {
  if (props.day > 9) {
    return 14
  } else {
    return 19
  }
})
</script>
