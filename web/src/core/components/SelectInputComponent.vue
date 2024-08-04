<template>
  <div class="form-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <select v-if="options" v-model="value">
      <option v-for="option in options" :key="option" :value="option">
        {{ option }}
      </option>
    </select>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  name?: string
  label?: string
  placeholder?: string
  options?: string[]
}>()

const value = defineModel('value')
</script>

<style scoped lang="scss">
.form-group {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 0.3em;
}

.number-unit-group {
  display: flex;
  gap: 0.3em;
  > input {
    flex: 1;
  }
}
</style>
