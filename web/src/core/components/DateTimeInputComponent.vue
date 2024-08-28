<template>
  <div class="input-group">
    <label v-if="label" :for="name">{{ label }}{{ model }}</label>
    <input type="datetime-local" :name="name" v-model="model" :list="`${name}_suggestions`" />
    <datalist v-if="list" :id="`${name}_suggestions`">
      <option v-for="value in list" :key="value" :value="value"></option>
    </datalist>
    <p v-if="errorMessage" class="error-message">
      {{ (name ? name + ' ' : 'value') + errorMessage }}
    </p>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { DateTimeFormat } from '../utils/time'

defineProps<{
  name?: string
  label?: string
  list?: string[]
  errorMessage?: string
}>()

const model = defineModel({
  get(value: string) {
    return dayjs(value).format(DateTimeFormat)
  },
})
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
</style>
