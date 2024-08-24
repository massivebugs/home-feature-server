<template>
  <WindowComponent
    class="cashbunny-window"
    :size="new RelativeSize(60, 70)"
    :title="t('cashbunny.title')"
    :controls="{
      minimize: true,
      maximize: true,
      close: true,
    }"
    :toolbar="toolbarOptions"
    :statusBarInfo="['Something goes here...', 'Something else here']"
    :resizable="true"
    @click-close="emit('clickClose')"
    v-slot="slotProps"
  >
    <div class="cashbunny-window__container">
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
            id: Tabs.schedule,
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
          <div class="cashbunny__tab-label">
            <OverviewTabIconComponent
              :fill="currentTab === Tabs.overview ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.overview') }}
          </div>
        </template>
        <template #overview>
          <div class="cashbunny__tab-body">
            <OverviewComponent />
          </div>
        </template>
        <template #planner_label>
          <div class="cashbunny__tab-label">
            <PlannerTabIconComponent
              width="30"
              height="30"
              :fill="currentTab === Tabs.planner ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.planner') }}
          </div>
        </template>
        <template #planner>
          <div class="cashbunny__tab-body"></div>
        </template>
        <template #schedule_label>
          <div class="cashbunny__tab-label">
            <ScheduleTabIconComponent
              :color="currentTab === Tabs.schedule ? '#ebebeb' : undefined"
              :day="today.getDate()"
            />
            {{ t('cashbunny.schedule') }}
          </div>
        </template>
        <template #schedule>
          <div class="cashbunny__tab-body">
            <ScheduleComponent />
          </div>
        </template>
        <template #accounts_label>
          <div class="cashbunny__tab-label">
            <AccountsTabIconComponent
              :fill="currentTab === Tabs.accounts ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.accounts') }}
          </div>
        </template>
        <template #accounts>
          <div class="cashbunny__tab-body">
            <AccountDataTableComponent v-if="slotProps.windowEl" :window-el="slotProps.windowEl" />
          </div>
        </template>
        <template #transactions_label>
          <div class="cashbunny__tab-label">
            <TransactionsTabIconComponent
              :fill="currentTab === Tabs.transactions ? '#ebebeb' : undefined"
            />
            {{ t('cashbunny.transactions') }}
          </div>
        </template>
        <template #transactions>
          <div class="cashbunny__tab-body">
            <TransactionDataTableComponent
              v-if="slotProps.windowEl"
              :window-el="slotProps.windowEl"
            />
          </div>
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
import { RelativeSize } from '@/core/models/relativeSize'
import AccountDataTableComponent from '@/modules/cashbunny/components/AccountDataTableComponent.vue'
import AccountsTabIconComponent from '@/modules/cashbunny/components/AccountsTabIconComponent.vue'
import OverviewTabIconComponent from '@/modules/cashbunny/components/OverviewTabIconComponent.vue'
import TransactionsTabIconComponent from '@/modules/cashbunny/components/TransactionsTabIconComponent.vue'
import OverviewComponent from './OverviewComponent.vue'
import PlannerTabIconComponent from './PlannerTabIconComponent.vue'
import ScheduleComponent from './ScheduleComponent.vue'
import ScheduleTabIconComponent from './ScheduleTabIconComponent.vue'
import TransactionDataTableComponent from './TransactionDataTableComponent.vue'

const Tabs = {
  overview: 'overview',
  planner: 'planner',
  schedule: 'schedule',
  accounts: 'accounts',
  transactions: 'transactions',
} as const

// type TabTypes = (typeof Tabs)[keyof typeof Tabs]

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const currentTab = ref<string>(Tabs.overview)
const today = new Date()
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

.cashbunny-window__container {
  width: 100%;
  height: 100%;
  padding: 5px;
  background: colors.$light-grey;
}

.cashbunny__tab-label {
  padding: 5px 0;
  display: flex;
  align-items: center;
  gap: 5px;
  font-weight: bold;
}

.cashbunny__tab-body {
  background-color: colors.$white;
  padding: 0.5em;
}
</style>
