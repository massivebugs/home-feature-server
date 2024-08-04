<template>
  <div>
    {{ icon }}
    <br />
    {{ name }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  icon: string
  name: string
  onDblClick: () => void
}>()
</script>
