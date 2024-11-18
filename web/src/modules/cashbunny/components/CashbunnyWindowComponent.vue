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
    :resizable="true"
    @click-close="emit('clickClose')"
    v-slot="windowProps"
  >
    <div
      class="cashbunny-window__body"
      :class="{
        'cashbunny-window__tab-body_fix-height': !windowProps.windowSizeQuery.md,
      }"
    >
      <TabGroupComponent
        class="cashbunny-window__tab-group"
        :class="{
          'cashbunny-window__tab-group_stretch': !windowProps.windowSizeQuery.lg,
        }"
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
            id: Tabs.schedules,
          },
          {
            id: Tabs.accounts,
          },
          {
            id: Tabs.transactions,
          },
        ]"
      >
        <template #overview>
          <OverviewTabIconComponent :fill="currentTab === Tabs.overview ? '#ebebeb' : undefined" />
          {{ t('cashbunny.overview') }}
        </template>
        <template #planner>
          <PlannerTabIconComponent
            width="30"
            height="30"
            :fill="currentTab === Tabs.planner ? '#ebebeb' : undefined"
          />
          {{ t('cashbunny.planner.name') }}
        </template>
        <template #schedules>
          <SchedulesTabIconComponent
            :color="currentTab === Tabs.schedules ? '#ebebeb' : undefined"
            :day="today.getDate()"
          />
          {{ t('cashbunny.schedules') }}
        </template>
        <template #accounts>
          <AccountsTabIconComponent :fill="currentTab === Tabs.accounts ? '#ebebeb' : undefined" />
          {{ t('cashbunny.accounts') }}
        </template>
        <template #transactions>
          <TransactionsTabIconComponent
            :fill="currentTab === Tabs.transactions ? '#ebebeb' : undefined"
          />
          {{ t('cashbunny.transactions') }}
        </template>
      </TabGroupComponent>
      <div class="cashbunny-window__tab-body">
        <OverviewComponent v-if="currentTab === Tabs.overview" :api="props.api" />
        <PlannerComponent v-else-if="currentTab === Tabs.planner" />
        <SchedulesComponent v-else-if="currentTab === Tabs.schedules" />
        <AccountDataTableComponent v-else-if="currentTab === Tabs.accounts" :api="props.api" />
        <TransactionDataTableComponent
          v-else-if="currentTab === Tabs.transactions"
          :api="props.api"
        />
      </div>
    </div>
    <FeaturesOverviewDialogComponent
      v-if="showFeaturesOverviewDialog"
      :size="new RelativeSize(50, 60)"
      @click-close="onClickCloseFeaturesOverviewDialog"
    />
  </WindowComponent>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import TabGroupComponent from '@/core/components/TabGroupComponent.vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import type { WindowToolbarRow } from '@/core/components/WindowToolbarComponent.vue'
import type { API } from '@/core/composables/useAPI'
import { RelativeSize } from '@/core/models/relativeSize'
import AccountDataTableComponent from '@/modules/cashbunny/components/AccountDataTableComponent.vue'
import AccountsTabIconComponent from '@/modules/cashbunny/components/AccountsTabIconComponent.vue'
import OverviewTabIconComponent from '@/modules/cashbunny/components/OverviewTabIconComponent.vue'
import TransactionsTabIconComponent from '@/modules/cashbunny/components/TransactionsTabIconComponent.vue'
import FeaturesOverviewDialogComponent from './FeaturesOverviewDialogComponent.vue'
import OverviewComponent from './OverviewComponent.vue'
import PlannerComponent from './PlannerComponent.vue'
import PlannerTabIconComponent from './PlannerTabIconComponent.vue'
import SchedulesComponent from './SchedulesComponent.vue'
import SchedulesTabIconComponent from './SchedulesTabIconComponent.vue'
import TransactionDataTableComponent from './TransactionDataTableComponent.vue'

const Tabs = {
  overview: 'overview',
  planner: 'planner',
  schedules: 'schedules',
  accounts: 'accounts',
  transactions: 'transactions',
} as const

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const props = defineProps<{
  api: API
}>()

const { t } = useI18n()
const currentTab = ref<string>(Tabs.overview)
const showFeaturesOverviewDialog = ref<boolean>(false)
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
                  console.log('Clicked CSV')
                },
              },
            ],
            [
              {
                label: t('common.close'),
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
        label: 'Help',
        contextMenuOptions: {
          itemGroups: [
            [
              {
                label: t('cashbunny.featuresOverview.name'),
                onClick: () => {
                  showFeaturesOverviewDialog.value = true
                },
              },
            ],
          ],
        },
      },
    ],
  },
])

const onClickCloseFeaturesOverviewDialog = () => {
  showFeaturesOverviewDialog.value = false
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';
.cashbunny-window__body {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.cashbunny-window__tab-group {
  max-width: 100%;
}

.cashbunny-window__tab-group_stretch {
  > * {
    flex: 1;
  }
}

.cashbunny-window__tab-body {
  flex: 1;
  overflow: auto;
}

.cashbunny-window__tab-body_fix-height {
  overflow: hidden;
}
</style>
