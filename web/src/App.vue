<template>
  <RouterView />
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { Process } from './core/models/process'
import { Program } from './core/models/program'
import { RelativePosition } from './core/models/relative_position'
import { RelativeSize } from './core/models/relative_size'
import { useStore } from './core/stores'
import BudgetPlannerView from './modules/budget_planner/views/BudgetPlannerView.vue'

const {t} = useI18n();

const store = useStore()
store.processes.set(
  '1',
  new Process('1', new Program(BudgetPlannerView, {}), {
    pos: new RelativePosition(10, 20),
    size: new RelativeSize(60, 70),
    title: t('budgetPlanner.name'),
    controls: {
      minimize: true,
      maximize: true,
      close: true,
    },
    toolbar: [
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
    ],
    statusBarInfo: ['Something goes here...', 'Something else here'],
    isDraggable: true,
    isResizable: true,
    isShowStatusBar: true,
    onClose: () => {
      store.removeProcess('1')
    },
  }),
)
</script>

<style scoped lang="scss"></style>
