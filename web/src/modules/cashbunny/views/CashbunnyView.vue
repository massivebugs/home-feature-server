<template>
  <WindowComponent
    :size="new RelativeSize(60, 70)"
    :title="t('cashbunny.name')"
    :controls="{
      minimize: true,
      maximize: true,
      close: true,
    }"
    :toolbar="[
      {
        isMenu: true,
        items: [
          {
            label: 'File',
            contextMenuOptions: {
              itemGroups: [
                [
                  {
                    icon: 'check',
                    label: 'Foo',
                    shortcutKey: 'Ctrl+A',
                    isDisabled: false,
                    onClick: () => {
                      console.log('Clicked Foo')
                    },
                  },
                  {
                    label: 'Bar',
                    shortcutKey: 'Ctrl+B',
                    isDisabled: true,
                    onClick: () => {
                      console.log('Clicked Bar')
                    },
                    children: [
                      [
                        {
                          label: 'Child of Bar',
                        },
                      ],
                    ],
                  },
                  {
                    label: 'Baz',
                    shortcutKey: 'Ctrl+C',
                    isDisabled: false,
                    onClick: () => {
                      console.log('Clicked Baz')
                    },
                    children: [
                      [
                        {
                          label: 'Child of Baz',
                        },
                      ],
                    ],
                  },
                ],
              ],
            },
          },
          {
            label: 'Edit',
          },
          {
            label: 'View',
          },
          {
            label: 'Favorites',
            contextMenuOptions: {
              itemGroups: [
                [
                  {
                    icon: 'check',
                    label: 'Foo',
                    shortcutKey: 'Ctrl+A',
                    isDisabled: false,
                    onClick: () => {
                      console.log('Clicked Foo')
                    },
                  },
                  {
                    label: 'Scan with TeamViewer_setup.exe',
                    shortcutKey: 'Ctrl+B',
                    isDisabled: true,
                    onClick: () => {
                      console.log('Clicked Bar')
                    },
                    children: [
                      [
                        {
                          label: 'Child of Bar',
                        },
                      ],
                    ],
                  },
                  {
                    label: 'Baz',
                    shortcutKey: 'Ctrl+C',
                    isDisabled: false,
                    onClick: () => {
                      console.log('Clicked Baz')
                    },
                    children: [
                      [
                        {
                          label: 'Child of Baz',
                        },
                      ],
                    ],
                  },
                ],
                [
                  {
                    icon: 'check',
                    label: 'Foo',
                    shortcutKey: 'Ctrl+A',
                    isDisabled: false,
                    onClick: () => {
                      console.log('Clicked Foo')
                    },
                  },
                  {
                    label: 'Scan with TeamViewer_setup.exe',
                    shortcutKey: 'Ctrl+B',
                    isDisabled: true,
                    onClick: () => {
                      console.log('Clicked Bar')
                    },
                    children: [
                      [
                        {
                          label: 'Child of Bar',
                        },
                      ],
                    ],
                  },
                  {
                    label: 'Baz',
                    shortcutKey: 'Ctrl+C',
                    isDisabled: false,
                    onClick: () => {
                      console.log('Clicked Baz')
                    },
                    children: [
                      [
                        {
                          label: 'Child of Baz',
                        },
                      ],
                    ],
                  },
                ],
              ],
            },
          },
          {
            label: 'Tools',
          },
          {
            label: 'Help',
          },
        ],
      },
    ]"
    :statusBarInfo="['Something goes here...', 'Something else here']"
    :isResizable="true"
    @click-close="emit('clickClose')"
  >
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
  </WindowComponent>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import TabGroupCompmonent from '@/core/components/TabGroupComponent.vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import { RelativeSize } from '@/core/models/relative_size'
import AccountingBalanceComponent from '@/modules/cashbunny/components/AccountingBalanceComponent.vue'
import AccountsDataTableComponent from '@/modules/cashbunny/components/AccountsDataTableComponent.vue'
import AccountsTabIconComponent from '@/modules/cashbunny/components/AccountsTabIconComponent.vue'
import OverviewTabIconComponent from '@/modules/cashbunny/components/OverviewTabIconComponent.vue'
import TransactionsTabIconComponent from '@/modules/cashbunny/components/TransactionsTabIconComponent.vue'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const currentTabId = ref<string>('overview')
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
