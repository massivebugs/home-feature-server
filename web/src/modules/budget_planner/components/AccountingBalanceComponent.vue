<template>
  <div class="twc-accounting-balance-component">
    <div
      class="twc-accounting-balance-component__balance"
      v-for="(money, idx) in balance"
      :key="idx"
    >
      {{ money.currency }}:{{ money.amount }}
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Balance } from '@/modules/budget_planner/models/balance'

defineProps<{
  balance: Balance[]
}>()
</script>

<style scoped lang="scss"></style>
