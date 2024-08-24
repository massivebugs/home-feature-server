<template>
  <svg
    width="10"
    height="10"
    viewBox="0 0 24 24"
    fill="none"
    stroke="black"
    xmlns="http://www.w3.org/2000/svg"
  >
    <g stroke-width="4" v-if="type === 'minimize'">
      <path d="M8,3V6A2,2,0,0,1,6,8H3" />

      <path d="M16,21V18a2,2,0,0,1,2-2h3" />

      <path d="M8,21V18a2,2,0,0,0-2-2H3" />

      <path d="M16,3V6a2,2,0,0,0,2,2h3" />
    </g>
    <g stroke-width="5" v-else>
      <path d="M3,8V5A2,2,0,0,1,5,3H8" />

      <path d="M21,16v3a2,2,0,0,1-2,2H16" />

      <path d="M3,16v3a2,2,0,0,0,2,2H8" />

      <path d="M21,8V5a2,2,0,0,0-2-2H16" />
    </g>
  </svg>
</template>

<script setup lang="ts">
defineProps<{
  type: 'minimize' | 'maximize'
}>()
</script>
