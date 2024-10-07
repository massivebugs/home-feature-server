<template>
  <div class="hfs-input-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <input
      :required="required"
      :disabled="disabled"
      :max="max"
      class="hfs-input-group__input"
      :class="{
        'hfs-input-group_has-error': !!errorMessage,
      }"
      :type="type ?? 'text'"
      :name="name"
      :placeholder="placeholder"
      v-model="value"
      :list="`${name}_suggestions`"
      :autocomplete="autocomplete"
    />
    <datalist v-if="list" :id="`${name}_suggestions`">
      <option v-for="value in list" :key="value" :value="value"></option>
    </datalist>
    <p v-if="errorMessage" class="hfs-input-group__error-message">
      {{ (name ? name + ' ' : 'value') + errorMessage }}
    </p>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  name?: string
  label?: string
  placeholder?: string
  list?: string[]
  errorMessage?: string
  type?: string
  disabled?: boolean
  autocomplete?: string
  required?: boolean
  max?: number
}>()

const value = defineModel()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-input-group {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 0.3em;
}

.hfs-input-group__error-message {
  margin: 0;
  font-size: 0.8em;
  color: colors.$red-cmyk;
}

.hfs-input-group__input {
  border-radius: 5px;
  padding: 0.5em 0.3em;
  border: 1px solid colors.$light-grey;
  transition: box-shadow 0.2s;
  background-color: white;

  &:focus:not(:disabled):not(.hfs-input-group_has-error) {
    outline: none;
    box-shadow: 0px 0px 5px 1px colors.$high-opacity-viridian;
    -webkit-box-shadow: 0px 0px 5px 1px colors.$high-opacity-viridian;
    -moz-box-shadow: 0px 0px 5px 1px colors.$high-opacity-viridian;
  }
}

.hfs-input-group_has-error {
  box-shadow: 0px 0px 5px 1px colors.$high-opacity-red-cmyk;
  -webkit-box-shadow: 0px 0px 5px 1px colors.$high-opacity-red-cmyk;
  -moz-box-shadow: 0px 0px 5px 1px colors.$high-opacity-red-cmyk;
}
</style>
