<template>
  <div class="form-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <input type="text" :name="name" :placeholder="placeholder" v-model="value" />
  </div>
</template>

<script setup lang="ts">
defineProps<{
  name?: string
  label?: string
  placeholder?: string
}>()

const value = defineModel()
</script>

<style scoped lang="scss">
.form-group {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 0.3em;
}
</style>
