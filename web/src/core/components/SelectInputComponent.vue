<template>
  <div class="input-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <select v-if="options" v-model="value">
      <option v-for="option in options" :key="option" :value="option">
        {{ option }}
      </option>
    </select>
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
  options?: string[]
  errorMessage?: string
}>()

const value = defineModel('value')
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
