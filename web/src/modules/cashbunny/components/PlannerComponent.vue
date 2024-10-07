<template>
  <div class="cashbunny-planner">
    <PlannerWizardComponent />
  </div>
</template>
<script setup lang="ts">
import PlannerWizardComponent from './PlannerWizardComponent.vue'
</script>

<style scoped lang="scss">
@use '@/assets/colors';
.cashbunny-planner {
  height: 100%;
  background-color: colors.$high-opacity-white;
}
</style>
