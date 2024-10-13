<template>
  <div class="hfs-range-input-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <div>
      <input
        :disabled="disabled"
        :min="min"
        :max="max"
        :step="step"
        type="range"
        :name="name"
        v-model="value"
      />
      {{ valueSuffix }}{{ value }}{{ valuePrefix }}
    </div>
    <p v-if="errorMessage" class="hfs-range-input-group__error-message">
      {{ (name ? name + ' ' : 'value') + errorMessage }}
    </p>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  name?: string
  label?: string
  errorMessage?: string
  disabled?: boolean
  max?: number
  min?: number
  step?: number
  valueSuffix?: string
  valuePrefix?: string
}>()

const value = defineModel()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-range-input-group {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 0.3em;
  > div {
    display: flex;
    gap: 0.7em;
    > input {
      flex: 1;
    }
  }
}

.hfs-range-input-group__error-message {
  margin: 0;
  font-size: 0.8em;
  color: colors.$red-cmyk;
}
</style>
