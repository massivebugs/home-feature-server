<template>
  <div class="accounting-balance-component">
    <div class="accounting-balance-component__balance" v-for="(money, idx) in balance" :key="idx">
      {{ money.currency }}:{{ money.amount }}
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Balance } from '@/modules/cashbunny/models/balance'

defineProps<{
  balance: Balance[]
}>()
</script>

<style scoped lang="scss"></style>
