import type { I18nOptions } from 'vue-i18n'

export const i18nOptions: I18nOptions = {
  legacy: false,
  locale: 'en',
  fallbackLocale: 'ja',
  messages: {
    en: {
      app: {
        name: 'Massivebugs Systems',
      },
      login: {
        title: 'Log into @:app.name',
        login: 'Login',
      },
      home: 'Home',
      ui: {
        success: 'Ok',
        cancel: 'Cancel',
      },
      cashbunny: {
        name: 'Cashbunny',
        title: 'Cashbunny Budget Planner',
        processAlreadyExistsTitle: 'Error',
        processAlreadyExistsMessage: 'This program can have only one running process',
        overview: 'Overview',
        accounts: 'Accounts',
        transactions: 'Transactions',
        balance: 'Balance',
        edit: 'Edit',
        delete: 'Delete',
        csv: 'Export to CSV',
        addAccount: 'Add Account',
        accountName: 'Account name',
        save: 'Save',
        accountDeleteConfirmTitle: 'Really delete this account?',
        accountDeleteConfirmMessage:
          'Deleting this account will affect your transactions and balances. Would you really like to delete this account?',
      },
      portfolio: {
        name: 'Portfolio',
        title: "Portfolio - {'@'}massivebugs",
        help: 'Help',
        about: 'About Portfolio',
        githubLinkTitle: 'Check out my GitHub',
      },
    },
    ja: {
      app: {
        name: 'Massivebugs Systems',
      },
      login: {
        title: '@:app.name にログインする',
        login: 'ログイン',
      },
      home: 'ホーム',
      ui: {
        success: 'Ok',
        cancel: 'キャンセル',
      },
      cashbunny: {
        name: 'Cashbunny',
        title: 'Cashbunny家計簿',
        processAlreadyExistsTitle: 'エラー',
        processAlreadyExistsMessage: 'すでに起動しているプロセスがあります',
        overview: 'Overview',
        accounts: 'Accounts',
        transactions: 'Transactions',
        balance: '合計',
        edit: '編集',
        delete: '削除',
        csv: 'CSVダウンロード',
        save: '保存',
        addAccount: 'アカウントを追加',
        accountName: 'アカウント名',
        accountDeleteConfirmTitle: '本当にこのアカウントを削除しますか？',
        accountDeleteConfirmMessage:
          'アカウントを削除すると金額情報やトランザクション記録も影響されます。本当に削除しますか？',
      },
      portfolio: {
        name: 'ポートフォリオ',
        title: "ポートフォリオ - {'@'}massivebugs",
        help: 'ヘルプ',
        about: 'ポートフォリオについて',
        githubLinkTitle: 'GitHubを見てみる',
      },
    },
  },
}
