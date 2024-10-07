<template>
  <section class="system-settings__section">
    <h2>{{ props.title }}</h2>
    <ul>
      <slot />
    </ul>
  </section>
</template>
<script setup lang="ts">
const props = defineProps<{
  title: string
}>()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.system-settings__section {
  background-color: colors.$high-opacity-white;
  border-radius: 5px;
  padding: 0.7em;
  display: flex;
  flex-direction: column;
  gap: 0.7em;

  h2 {
    margin: 0;
    font-size: 1.2em;
    font-weight: 500;
    user-select: none;
  }

  ul {
    margin: 0;
    padding: 0;
    list-style-type: none;
    display: flex;
    flex-direction: column;
    gap: 0.3em;
  }
}
</style>
