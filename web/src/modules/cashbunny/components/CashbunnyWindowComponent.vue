<template>
  <WindowComponent
    :size="new RelativeSize(60, 70)"
    :title="t('cashbunny.title')"
    :controls="{
      minimize: true,
      maximize: true,
      close: true,
    }"
    :toolbar="toolbarOptions"
    :statusBarInfo="['Something goes here...', 'Something else here']"
    :isResizable="true"
    @click-close="emit('clickClose')"
  >
    <div class="container">
      <TabGroupComponent
        @tab-click="(e) => (currentTab = e.tabId)"
        :selected-tab-id="currentTab"
        :tabs="[
          {
            id: Tabs.overview,
          },
          {
            id: Tabs.planner,
          },
          {
            id: Tabs.accounts,
          },
          {
            id: Tabs.transactions,
          },
        ]"
      >
        <template #overview_label>
          <div class="tab-label">
            <OverviewTabIconComponent
              :fill="currentTab === Tabs.overview ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.overview') }}
          </div>
        </template>
        <template #overview>
          <OverviewComponent />
        </template>
        <template #planner_label>
          <div class="tab-label">
            <PlannerTabIconComponent
              width="30"
              height="30"
              :fill="currentTab === Tabs.planner ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.planner') }}
          </div>
        </template>
        <template #planner> </template>
        <template #accounts_label>
          <div class="tab-label">
            <AccountsTabIconComponent
              :fill="currentTab === Tabs.accounts ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.accounts') }}
          </div>
        </template>
        <template #accounts>
          <AccountDataTableComponent />
        </template>
        <template #transactions_label>
          <div class="tab-label">
            <TransactionsTabIconComponent
              :fill="currentTab === Tabs.transactions ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.transactions') }}
          </div>
        </template>
        <template #transactions>
          <TransactionDataTableComponent />
        </template>
      </TabGroupComponent>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import TabGroupComponent from '@/core/components/TabGroupComponent.vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import type { WindowToolbarRow } from '@/core/components/WindowToolbarComponent.vue'
import { RelativeSize } from '@/core/models/relative_size'
import AccountDataTableComponent from '@/modules/cashbunny/components/AccountDataTableComponent.vue'
import AccountsTabIconComponent from '@/modules/cashbunny/components/AccountsTabIconComponent.vue'
import OverviewTabIconComponent from '@/modules/cashbunny/components/OverviewTabIconComponent.vue'
import TransactionsTabIconComponent from '@/modules/cashbunny/components/TransactionsTabIconComponent.vue'
import OverviewComponent from './OverviewComponent.vue'
import PlannerTabIconComponent from './PlannerTabIconComponent.vue'
import TransactionDataTableComponent from './TransactionDataTableComponent.vue'

const Tabs = {
  overview: 'overview',
  planner: 'planner',
  accounts: 'accounts',
  transactions: 'transactions',
} as const

// type TabTypes = (typeof Tabs)[keyof typeof Tabs]

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const currentTab = ref<string>(Tabs.overview)
const toolbarOptions = computed<WindowToolbarRow[]>(() => [
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
                label: t('cashbunny.csv'),
                shortcutKey: 'Ctrl+A',
                isDisabled: false,
                onClick: () => {
                  console.log('Clicked Foo')
                },
              },
            ],
            [
              {
                label: 'Exit',
                shortcutKey: 'Alt+F4',
                isDisabled: false,
                onClick: () => {
                  emit('clickClose')
                },
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
])
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.container {
  width: 100%;
  height: 100%;
  padding: 5px;
}

.tab-label {
  padding: 5px 0;
  display: flex;
  align-items: center;
  gap: 5px;
  font-weight: bold;
}
</style>
