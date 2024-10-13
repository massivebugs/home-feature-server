<template>
  <div class="input-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <select
      v-if="options"
      v-model="value"
      :required="required"
      class="hfs-input-group__input"
      :class="{
        'hfs-input-group_has-error': !!errorMessage,
      }"
    >
      <option v-for="{ label, value } in options" :key="value" :value="value">
        {{ label }}
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
  options?: { label: string; value: any }[]
  errorMessage?: string
  required?: boolean
}>()

const value = defineModel()
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
