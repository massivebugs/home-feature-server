<template>
  <div class="twc-budget-planner-view">
    <WindowsXPTabGroupCompmonent
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
        {
          id: 'incomes',
        },
        {
          id: 'expenses',
        },
      ]"
    >
      <template #overview_label>
        <div class="twc-budget-planner-tab-label">
          <OverviewTabIconComponent />
          {{ $t('budgetPlanner.overview') }}
        </div>
      </template>
      <template #overview>
        <div>
          <AccountingBalanceComponent :balance="store.summary?.balances || []" />
        </div>
      </template>
      <template #accounts_label>
        <div class="twc-budget-planner-tab-label">
          <AccountsTabIconComponent />
          {{ $t('budgetPlanner.accounts') }}
        </div>
      </template>
      <template #accounts>
        <div>
          <DataTableComponent />
        </div>
      </template>
      <template #transactions_label>
        <div class="twc-budget-planner-tab-label">
          <TransactionsTabIconComponent />
          {{ $t('budgetPlanner.transactions') }}
        </div>
      </template>
      <template #transactions>
        <div>Transactions</div>
      </template>
      <template #incomes_label>
        <div class="twc-budget-planner-tab-label">
          <IncomesTabIconComponent />
          {{ $t('budgetPlanner.incomes') }}
        </div>
      </template>
      <template #incomes>
        <div>Incomes</div>
      </template>
      <template #expenses_label>
        <div class="twc-budget-planner-tab-label">
          <ExpensesTabIconComponent />
          {{ $t('budgetPlanner.expenses') }}
        </div>
      </template>
      <template #expenses>
        <div>Expenses</div>
      </template>
    </WindowsXPTabGroupCompmonent>
  </div>
</template>

<script setup lang="ts">
import WindowsXPTabGroupCompmonent from '@/core/components/WindowsXPTabGroupComponent.vue'
import DataTableComponent from '@/modules/budget_planner/components/DataTableComponent.vue'
import AccountingBalanceComponent from '@/modules/budget_planner/components/AccountingBalanceComponent.vue'
import OverviewTabIconComponent from '@/modules/budget_planner/components/OverviewTabIconComponent.vue'
import AccountsTabIconComponent from '@/modules/budget_planner/components/AccountsTabIconComponent.vue'
import TransactionsTabIconComponent from '@/modules/budget_planner/components/TransactionsTabIconComponent.vue'
import IncomesTabIconComponent from '@/modules/budget_planner/components/IncomesTabIconComponent.vue'
import ExpensesTabIconComponent from '@/modules/budget_planner/components/ExpensesTabIconComponent.vue'
import { useStore } from '@/modules/budget_planner/stores'
import { onMounted, ref } from 'vue'

const store = useStore()
const currentTabId = ref<string>('overview')

onMounted(() => {
  store.fetchSummary()
})
</script>

<style scoped lang="scss">
.twc-budget-planner-view {
  width: 100%;
  height: 98%;
  padding: 5px 2px 0 2px;
}

.twc-budget-planner-tab-label {
  padding: 5px 0;
  display: flex;
  align-items: center;
  gap: 5px;
  font-weight: bold;
}
</style>
