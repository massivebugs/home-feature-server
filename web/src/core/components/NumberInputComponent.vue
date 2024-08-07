<template>
  <div class="input-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <div class="number-unit-group">
      <input
        type="number"
        :name="name"
        :placeholder="placeholder"
        :min="min"
        :max="max"
        :step="step"
        v-model="value"
      />
      <select v-if="units" v-model="unit">
        <option v-for="unit in units" :key="unit" :value="unit">
          {{ unit }}
        </option>
      </select>
    </div>
    <p v-if="errorMessage" class="error-message">
      {{ (name ? name + ' ' : 'value') + errorMessage }}
    </p>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  name?: string
  label?: string
  placeholder?: string
  min?: number
  max?: number
  step?: number
  units?: string[]
  errorMessage?: string
}>()

const value = defineModel('value')
const unit = defineModel('unit')
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.input-group {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 0.3em;
  .error-message {
    margin: 0;
    font-size: 0.8em;
    color: colors.$red-cmyk;
  }
}

.number-unit-group {
  display: flex;
  gap: 0.3em;
  > input {
    flex: 1;
  }
}
</style>
