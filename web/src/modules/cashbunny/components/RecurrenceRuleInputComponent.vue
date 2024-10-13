<template>
  <div class="hfs-recurrence-rule-input-group">
    <label v-if="label" :for="name">{{ label }}</label>
    <DateInputComponent
      :required="true"
      name="dtstart"
      :label="t('cashbunny.recurrenceRule.dtstart')"
      v-model="recurrenceRule.dtstart"
    />
    <DateInputComponent
      :required="true"
      name="until"
      :label="t('cashbunny.recurrenceRule.until')"
      v-model="recurrenceRule.until"
    />
    <SelectInputComponent
      :required="true"
      name="freq"
      :label="t('cashbunny.recurrenceRule.freq')"
      :options="
        Object.keys(FrequencyStrs).map((v) => ({
          label: t(`cashbunny.recurrenceRule.frequencies.${v}`),
          value: v,
        }))
      "
      v-model="recurrenceRule.freq"
    />
    <NumberInputComponent
      :required="true"
      name="interval"
      :label="t('cashbunny.recurrenceRule.interval')"
      placeholder="1"
      :min="1"
      v-model:value="recurrenceRule.interval"
    />
    <NumberInputComponent
      :required="true"
      name="count"
      :label="t('cashbunny.recurrenceRule.count')"
      placeholder="0"
      :min="0"
      v-model:value="recurrenceRule.count"
    />
    <p>
      {{ stringRepresentation }}
    </p>
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import DateInputComponent from '@/core/components/DateInputComponent.vue'
import NumberInputComponent from '@/core/components/NumberInputComponent.vue'
import SelectInputComponent from '@/core/components/SelectInputComponent.vue'
import { FrequencyStrs, RecurrenceRule } from '@/modules/cashbunny/models/recurrence_rule'
import type { RecurrenceRuleDto } from '../models/dto'

defineProps<{
  name?: string
  label?: string
}>()

const { t } = useI18n()
const recurrenceRule = defineModel<RecurrenceRuleDto>({
  default: {
    freq: FrequencyStrs.MONTHLY,
    dtstart: dayjs().toISOString(),
    count: 0,
    interval: 1,
    until: dayjs().add(1, 'year').toISOString(),
  },
})

const stringRepresentation = computed<string>(() => {
  return new RecurrenceRule(recurrenceRule.value).toHumanFriendlyString()
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-recurrence-rule-input-group {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  gap: 0.3em;
}
</style>
