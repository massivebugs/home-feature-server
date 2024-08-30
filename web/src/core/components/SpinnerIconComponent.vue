<!-- Spinners from https://github.com/n3r4zzurr0/svg-spinners -->
<template>
  <svg
    v-if="type === SpinnerTypes.ring"
    width="16"
    height="16"
    stroke="#000"
    viewBox="0 0 24 24"
    xmlns="http://www.w3.org/2000/svg"
  >
    <g class="spinner_V8m1">
      <circle cx="12" cy="12" r="9.5" fill="none" stroke-width="3"></circle>
    </g>
  </svg>
  <svg
    v-else-if="type === SpinnerTypes.dots"
    width="19"
    height="19"
    viewBox="0 0 24 24"
    xmlns="http://www.w3.org/2000/svg"
  >
    <circle class="spinner_qM83" cx="4" cy="12" r="3" />
    <circle class="spinner_qM83 spinner_oXPr" cx="12" cy="12" r="3" />
    <circle class="spinner_qM83 spinner_ZTLf" cx="20" cy="12" r="3" />
  </svg>
</template>

<script lang="ts">
export const SpinnerTypes = {
  ring: 'ring',
  dots: 'dots',
} as const
export type SpinnerType = (typeof SpinnerTypes)[keyof typeof SpinnerTypes]
</script>

<script setup lang="ts">
defineProps<{
  type: SpinnerType
}>()
</script>

<style scoped lang="scss">
// Ring
.spinner_V8m1 {
  transform-origin: center;
  animation: spinner_zKoa 2s linear infinite;
}
.spinner_V8m1 circle {
  stroke-linecap: round;
  animation: spinner_YpZS 1.5s ease-in-out infinite;
}
@keyframes spinner_zKoa {
  100% {
    transform: rotate(360deg);
  }
}
@keyframes spinner_YpZS {
  0% {
    stroke-dasharray: 0 150;
    stroke-dashoffset: 0;
  }
  47.5% {
    stroke-dasharray: 42 150;
    stroke-dashoffset: -16;
  }
  95%,
  100% {
    stroke-dasharray: 42 150;
    stroke-dashoffset: -59;
  }
}

// Dots
.spinner_qM83 {
  animation: spinner_8HQG 1.05s infinite;
}
.spinner_oXPr {
  animation-delay: 0.1s;
}
.spinner_ZTLf {
  animation-delay: 0.2s;
}
@keyframes spinner_8HQG {
  0%,
  57.14% {
    animation-timing-function: cubic-bezier(0.33, 0.66, 0.66, 1);
    transform: translate(0);
  }
  28.57% {
    animation-timing-function: cubic-bezier(0.33, 0, 0.66, 0.33);
    transform: translateY(-6px);
  }
  100% {
    transform: translate(0);
  }
}
</style>
