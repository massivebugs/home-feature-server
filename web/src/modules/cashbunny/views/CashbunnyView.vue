<template>
  <div class="budget-planner-view">
    <TabGroupCompmonent
      @tab-click="(e) => (currentTabId = e.tabId)"
      :selected-tab-id="currentTabId"
      :tabs="[
        {
          id: 'overview',
        },
        {
          id: 'accounts',
        },
        {
          id: 'transactions',
        },
      ]"
    >
      <template #overview_label>
        <div class="budget-planner-tab-label">
          <OverviewTabIconComponent />
          {{ t('cashbunny.overview') }}
        </div>
      </template>
      <template #overview>
        <!-- <AccountingBalanceComponent :balance="store.summary?.balances || []" /> -->
      </template>
      <template #accounts_label>
        <div class="budget-planner-tab-label">
          <AccountsTabIconComponent />
          {{ t('cashbunny.accounts') }}
        </div>
      </template>
      <template #accounts>
        <AccountsDataTableComponent />
      </template>
      <template #transactions_label>
        <div class="budget-planner-tab-label">
          <TransactionsTabIconComponent />
          {{ t('cashbunny.transactions') }}
        </div>
      </template>
    </TabGroupCompmonent>
  </div>
</template>

<script setup lang="ts">
import TabGroupCompmonent from '@/core/components/TabGroupComponent.vue'
import AccountsDataTableComponent from '@/modules/cashbunny/components/AccountsDataTableComponent.vue'
import AccountingBalanceComponent from '@/modules/cashbunny/components/AccountingBalanceComponent.vue'
import OverviewTabIconComponent from '@/modules/cashbunny/components/OverviewTabIconComponent.vue'
import AccountsTabIconComponent from '@/modules/cashbunny/components/AccountsTabIconComponent.vue'
import TransactionsTabIconComponent from '@/modules/cashbunny/components/TransactionsTabIconComponent.vue'
import { useStore } from '@/modules/cashbunny/stores'
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const store = useStore()
const currentTabId = ref<string>('overview')

onMounted(() => {
  // TODO
})
</script>

<style scoped lang="scss">
.budget-planner-view {
  width: 100%;
  height: 100%;
  padding: 5px;
}

.budget-planner-tab-label {
  padding: 5px 0;
  display: flex;
  align-items: center;
  gap: 5px;
  font-weight: bold;
}
</style>
