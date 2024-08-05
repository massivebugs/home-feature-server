<template>
  <svg
    :width="widthPx ?? '30'"
    :height="heightPx ?? '30'"
    viewBox="0 0 24 24"
    role="img"
    focusable="false"
    aria-hidden="true"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M2 22h21v1H1V1h1zm8.613-3.193L14.61 5.793l3.87 8.956 3.726-6.35-.862-.506-2.728 4.649-4.165-9.637-4.097 13.338-3.877-7.108L2.9 16.552l.9.435 2.732-5.664z"
    />
  </svg>
</template>

<script setup lang="ts">
defineProps<{
  widthPx?: number
  heightPx?: number
}>()
</script>
