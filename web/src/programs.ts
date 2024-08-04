import { Program } from './core/models/program'
import CashbunnyView from './modules/cashbunny/views/CashbunnyView.vue'
import { RelativePosition } from './core/models/relative_position'
import { RelativeSize } from './core/models/relative_size'
import type { PiniaStore } from './utils/pinia'
import type { useStore } from './core/stores'
import { type ComposerTranslation } from 'vue-i18n'

// Returns all of the programs to be registered to the system,
// and options for window display
export const getPrograms = (
  coreStore: PiniaStore<typeof useStore>,
  t: ComposerTranslation,
): Program[] => [
  new Program(
    t('cashbunny.name'),
    CashbunnyView,
    {},
    {
      size: new RelativeSize(60, 70),
      title: t('cashbunny.name'),
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
      isResizable: true,
    },
  ),
]
