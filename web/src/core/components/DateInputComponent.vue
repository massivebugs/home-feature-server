<template>
  <div class="hfs-input-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <input
      :name="name"
      type="date"
      class="hfs-input-group__input"
      :class="{
        'hfs-input-group_has-error': !!errorMessage,
      }"
      v-model="model"
      :list="`${name}_suggestions`"
    />
    <p v-if="errorMessage" class="hfs-input-group__error-message">
      {{ (name ? name + ' ' : 'value') + errorMessage }}
    </p>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { DateFormat } from '../utils/time'

defineProps<{
  name?: string
  label?: string
  list?: string[]
  errorMessage?: string
}>()

const model = defineModel({
  get(value: string) {
    return dayjs(value).format(DateFormat)
  },
})
</script>
