<template>
  <div class="cashbunny__schedule-container">
    <div class="cashbunny__schedule__controls">
      <div></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import dayjs, { Dayjs } from 'dayjs'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCashbunnyStore } from '../stores'

const { t } = useI18n()
const store = useCashbunnyStore()
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.cashbunny__schedule__controls {
  display: flex;
  justify-content: space-between;
}

.cashbunny__schedule__calendar {
  border-radius: 0;
  width: 100%;
}
</style>
