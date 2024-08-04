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
        overview: 'Overview',
        accounts: 'Accounts',
        transactions: 'Transactions',
        balance: 'Balance',
        edit: 'Edit',
        delete: 'Delete',
        copy: 'Copy',
        csv: 'Export to CSV',
        accountDeleteConfirmTitle: 'Really delete this account?',
        accountDeleteConfirmMessage:
          'Deleting this account will affect your transactions and balances. Would you really like to delete this account?',
      },
      portfolio: {
        name: 'Portfolio',
        toolbarTitle: "Portfolio - {'@'}massivebugs",
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
        overview: 'Overview',
        accounts: 'Accounts',
        transactions: 'Transactions',
        balance: '合計',
        edit: '編集',
        delete: '削除',
        copy: 'コピー',
        csv: 'CSVダウンロード',
        accountDeleteConfirmTitle: '本当にこのアカウントを削除しますか？',
        accountDeleteConfirmMessage:
          'アカウントを削除すると金額情報やトランザクション記録も影響されます。本当に削除しますか？',
      },
      portfolio: {
        name: 'ポートフォリオ',
        toolbarTitle: "ポートフォリオ - {'@'}massivebugs",
        help: 'ヘルプ',
        about: 'ポートフォリオについて',
        githubLinkTitle: 'GitHubを見てみる',
      },
    },
  },
}
