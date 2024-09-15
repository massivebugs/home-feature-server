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
  background-color: colors.$white;
  height: 100%;
}
</style>
