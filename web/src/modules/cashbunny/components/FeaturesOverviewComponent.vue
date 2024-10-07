<template>
  <div
    class="cashbunny-features-overview"
    :class="{ 'cashbunny-features-overview_column': isColumnLayout }"
  >
    <SimplePageList
      :pages="Object.values(FeaturesOverviewPages).map((v) => ({ key: v, name: startCase(v) }))"
      :type="isColumnLayout ? 'tabs' : 'list'"
      @click-page="onClickSimplePageListPage"
    />
    <div>
      <template v-if="currentPage === FeaturesOverviewPages.overview"> </template>
      <template v-else-if="currentPage === FeaturesOverviewPages.planner">
        <h2>Planner</h2>
        <p></p>
      </template>
      <template v-else-if="currentPage === FeaturesOverviewPages.schedules"> </template>
      <template v-else-if="currentPage === FeaturesOverviewPages.accounts"> </template>
      <template v-else-if="currentPage === FeaturesOverviewPages.transactions"> </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { startCase } from 'lodash'
import { computed, inject, ref } from 'vue'
import SimplePageList, {
  type SimplePageListClickPageEvent,
} from '@/core/components/SimplePageList.vue'
import type { WindowSizeQuery } from '@/core/components/WindowComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import type { RelativeSize } from '@/core/models/relativeSize'

const FeaturesOverviewPages = {
  overview: 'overview',
  planner: 'planner',
  schedules: 'schedules',
  accounts: 'accounts',
  transactions: 'transactions',
} as const
type FeaturesOverviewPage = (typeof FeaturesOverviewPages)[keyof typeof FeaturesOverviewPages]

defineProps<{
  pos?: RelativePosition | 'center'
  size?: RelativeSize
}>()

const windowSizeQuery = inject<WindowSizeQuery>('windowSizeQuery')
const isColumnLayout = computed(() => {
  return !windowSizeQuery?.lg
})
const currentPage = ref<FeaturesOverviewPage>(FeaturesOverviewPages.overview)

const onClickSimplePageListPage = ({ key }: SimplePageListClickPageEvent<FeaturesOverviewPage>) => {
  currentPage.value = key
}
</script>

<style scoped lang="scss">
.cashbunny-features-overview {
  width: 100%;
  height: 100%;
  display: flex;
  gap: 0.5em;
}

.cashbunny-features-overview_column {
  flex-direction: column;
}

.cashbunny-features-overview > pre {
  min-width: 300px;
}
</style>
