import type { I18nOptions } from 'vue-i18n'

export const Locales = {
  en: 'en',
  ja: 'ja',
}
export type Locale = (typeof Locales)[keyof typeof Locales]

export const i18nOptions: I18nOptions = {
  legacy: false,
  fallbackLocale: Locales.en,
  messages: {
    en: {
      common: {
        back: 'Back',
        continue: 'Continue',
        save: 'Save',
        close: 'Close',
        help: 'Help',
        action: 'Action',
        file: 'File',
        exit: 'Exit',
        about: 'About {v}',
      },
      app: {
        name: 'Home Feature Server',
      },
      login: {
        title: 'Log into @:app.name',
        submit: 'Login',
      },
      createAccount: {
        title: 'Create an account',
        message: "Don't have an account?",
        linkTitle: 'Create one here',
        disclaimer:
          "Submit the following form to create an account.\nYour account won't be available until administrator approval.",
        submit: 'Submit',
        pendingApproval:
          'Your account creation request has been\nsuccessfully submitted.\nPlease wait until the administrator enables your account.\nNo emails will be sent when approved (under development).\nPlease check regularly by logging in to view your approval status.',
      },
      desktop: {
        logOutDialogTitle: 'Log out',
        logOutConfirmMessage: 'Are you sure you want to log out?',
      },
      home: 'Home',
      ui: {
        success: 'Ok',
        cancel: 'Cancel',
      },
      systemSettings: {
        name: 'System Settings',
        locale: {
          default: 'Use system defaults',
          en: 'English',
          ja: '日本語',
        },
        preferences: {
          title: 'Preferences',
          language: {
            title: 'Language',
          },
        },
      },
      cashbunny: {
        name: 'Cashbunny',
        title: 'Cashbunny Budget Planner',
        processAlreadyExistsTitle: 'Error',
        processAlreadyExistsMessage: 'This program can have only one running process',
        errorLoadingDataTitle: 'Error',
        errorLoadingDataMessage: 'An error occured while loading data: {e}',
        assets: {
          name: 'Assets',
        },
        liabilities: {
          name: 'Liabilities',
        },
        recurrenceRule: {
          freq: 'Frequency',
          dtstart: 'Start from',
          count: 'Max recurrences',
          interval: 'Interval',
          until: 'End on',
          frequencies: {
            YEARLY: 'Yearly',
            MONTHLY: 'Monthly',
            WEEKLY: 'Weekly',
            DAILY: 'Daily',
            HOURLY: 'Hourly',
            MINUTELY: 'Minutely',
            SECONDLY: 'Secondly',
          },
        },
        featuresOverview: {
          name: 'Overview of features',
        },
        overview: 'Overview',
        planner: {
          name: 'Planner',
          wizard: {
            welcome: {
              name: 'Welcome',
              question:
                "Welcome! Let's set up your budget plan.\nWe will ask you a few questions about your current budget status.",
              info: 'The questions may take a few minutes to complete.\nWould you like to get started now?',
              yes: "Yes, let's get started",
              no: 'No, remind me later',
            },
            assets: {
              name: 'Assets',
              question:
                'What are your current assets?\nPlease list your accounts and their balances.',
              info: 'Assets are things you own that can be converted to cash or have value.',
            },
            revenues: {
              name: 'Revenues',
              question:
                'What are your sources of income?\nPlease provide details for your revenue.',
              info: 'Include any income you receive regularly.',
            },
            liabilities: {
              name: 'Liabilities',
              question:
                'Do you have any current liabilities?\nPlease list your debts and the amounts.',
              info: 'Liabilities are debts or obligations that you owe.',
            },
            expenses: {
              name: 'Expenses',
              question:
                'What are your regular monthly expenses?\nPlease categorize and enter amounts.',
              info: 'Include any expenses you spend regularly.',
            },
            complete: {
              name: 'Complete',
              question: 'Your budget planner is set up!',
              info: 'You may now exit the wizard.',
            },
          },
          asset: {
            name: 'Name',
            namePlaceholder: 'Checking account, savings in Some Bank',
            description: 'Description',
            descriptionPlaceholder: 'e.g. What this account is for',
            amount: 'Amount',
            currency: 'Currency',
            presets: {
              checkingAccount: {
                name: 'Checking account',
                description: 'I use this for everyday transactions',
              },
              savingsAccount: {
                name: 'Savings account',
                description: 'Store all of my savings here',
              },
              sinkingFunds: {
                name: 'Sinking funds',
                description: 'For discretionary payments not planned for',
              },
            },
          },
          revenue: {
            description: 'Description',
            descriptionPlaceholder: 'e.g. Salary wired from my company',
            amount: 'Amount',
            currency: 'Currency',
            from: 'From',
            to: 'To',
            recurrence: 'Payment Schedule',
            presets: {
              salary: {
                description: 'Salary (biweekly)',
                from: 'Some Company',
              },
              pension: {
                description: 'Pension (monthly)',
                from: 'Government',
              },
              dividends: {
                description: 'Dividends (quarterly)',
                from: 'Some Company',
              },
            },
          },
          liability: {
            description: 'Description',
            descriptionPlaceholder: 'e.g. Car loan',
            amount: 'Amount',
            currency: 'Currency',
            from: 'From',
            to: 'To',
            recurrence: 'Payment Schedule',
            presets: {
              studentLoans: {
                description: 'Student Loans (fixed interest rate)',
                from: 'Some School',
              },
              carLoans: {
                description: 'Car Loans (fixed interest rate)',
                from: 'Some Credit Union',
              },
              mortgage: {
                description: 'Mortgage (adjustable interest rate)',
                from: 'Some Bank',
              },
            },
          },
        },
        schedules: 'Schedules',
        accounts: 'Accounts',
        transactions: 'Transactions',
        scheduledTransactions: 'Scheduled Transactions',
        balance: 'Balance',
        edit: 'Edit',
        delete: 'Delete',
        csv: 'Export to CSV',
        save: 'Save',
        revenue: 'Revenue',
        expense: 'Expense',
        profit: 'Profit',
        netWorth: 'Net Worth',
        overviewForDate: 'Overview for {v}',
        overviewProfitLossSummary: 'Profit/Loss',
        overviewProfit: 'Profit',
        addAccount: 'Add Account',
        editAccount: 'Edit Account',
        accountName: 'Name',
        accountNamePlaceholder: 'e.g. ABC Bank account for savings',
        accountCategory: 'Category',
        accountDescription: 'Description',
        accountDescriptionPlaceholder:
          "e.g. Dump all savings to this account, can't withdraw until 2030/05/01",
        accountBalance: 'Balance',
        accountCurrency: 'Currency',
        accountType: 'Type',
        accountCreatedAt: 'Created at',
        accountUpdatedAt: 'Updated at',
        accountDeleteConfirmTitle: 'Confirm deletion',
        accountDeleteConfirmMessage:
          'Deleting this account will affect your transactions and balances. \nWould you really like to delete this account? | Deleting this account will affect your transactions and balances. Would you really like to delete these ({count}) accounts?',
        createTransaction: 'Create Transaction',
        editTransaction: 'Edit Transaction',
        transactionDescription: 'Description',
        transactionDescriptionPlaceholder: 'Shopping for groceries',
        transactionAmount: 'Amount',
        transactionCurrency: 'Currency',
        transactionSourceAccount: 'From',
        transactionDestinationAccount: 'To',
        transactionTransactedAt: 'Transacted at',
        transactionTransactedAtPlaceholder: '2006-01-02 15:04:05',
        transactionCreatedAt: 'Created at',
        transactionUpdatedAt: 'Updated at',
        transactionDeleteConfirmTitle: 'Confirm deletion',
        transactionDeleteConfirmMessage:
          'Would you really like to delete this transaction? | Would you really like to delete these ({count}) transactions?',
      },
      portfolio: {
        name: 'Portfolio',
        title: "Portfolio - {'@'}massivebugs",
        exit: 'Exit',
        help: 'Help',
        contact: 'Contact me',
        about: 'About Portfolio',
        githubLinkTitle: 'Check out my GitHub',
        sendContactMessage: 'Send message',
        aboutDialogTitle: 'About',
      },
    },
    ja: {
      common: {
        back: '前へ戻る',
        continue: '次に進む',
        save: '保存',
        close: '閉じる',
        help: 'ヘルプ',
        action: '操作',
        file: 'ファイル',
        exit: '閉じる',
        about: '{v}について',
      },
      app: {
        name: 'Home Feature Server',
      },
      login: {
        title: '@:app.name にログインする',
        submit: 'ログイン',
      },
      createAccount: {
        title: 'アカウントを作成する',
        message: 'アカウントをお持ちでない方は',
        linkTitle: 'こちらから作成してください。',
        disclaimer:
          '以下のフォームに入力してアカウントを作成してください。\n管理者の承認が完了するまで、アカウントは利用できません。',
        submit: '送信',
        pendingApproval:
          'アカウント作成リクエストが正常に送信されました。\n管理者がアカウントを有効にするまでお待ちください。\n承認時のメール通知は現在開発中のため送信されません。\n承認状況を確認するには、定期的にログインしてください。',
      },
      desktop: {
        logOutDialogTitle: 'ログアウト',
        logOutConfirmMessage: '本当にログアウトしますか？',
      },
      home: 'ホーム',
      ui: {
        success: 'Ok',
        cancel: 'キャンセル',
      },
      systemSettings: {
        name: 'システム設定',
        locale: {
          default: 'システムデフォルト',
          en: 'English',
          ja: '日本語',
        },
        preferences: {
          title: 'ユーザ設定',
          language: {
            title: '言語',
          },
        },
      },
      cashbunny: {
        name: 'Cashbunny',
        title: 'Cashbunny家計簿',
        processAlreadyExistsTitle: 'エラー',
        processAlreadyExistsMessage: 'すでに起動しているプロセスがあります',
        errorLoadingDataTitle: 'エラー',
        errorLoadingDataMessage: 'データ取得中エラーが発生しました: {e}',
        assets: {
          name: '資産',
        },
        liabilities: {
          name: '負債',
        },
        recurrenceRule: {
          freq: '頻度',
          dtstart: '開始日',
          count: '最大数',
          interval: '間隔',
          until: '終了日',
          frequencies: {
            YEARLY: '毎年',
            MONTHLY: '毎月',
            WEEKLY: '毎週',
            DAILY: '毎日',
            HOURLY: '毎時間',
            MINUTELY: '毎分',
            SECONDLY: '毎秒',
          },
        },
        featuresOverview: {
          name: '機能の概要',
        },
        overview: '概要',
        planner: {
          name: 'プランナー',
          wizard: {
            welcome: {
              name: 'ようこそ',
              question:
                'ようこそ！予算計画を設定しましょう。\n現在の予算状況に関するいくつかの質問をします。',
              info: '質問には数分かかる場合があります。\n今すぐ始めますか？',
              yes: '始めめる',
              no: 'いいえ、後でリマインドする',
            },
            assets: {
              name: '資産',
              question: '現在の資産は何ですか？\nアカウントとその残高をリストしてください。',
              info: '資産とは、現金に換えられるものや価値のあるものです。',
            },
            revenues: {
              name: '経常収入',
              question: '収入源は何ですか？\n収入の詳細を教えてください。',
              info: '定期的に受け取る収入をすべて含めてください。',
            },
            liabilities: {
              name: '負債',
              question: '現在の負債はありますか？\n負債とその金額をリストしてください。',
              info: '負債とは、支払わなければならない債務や義務です。',
            },
            expenses: {
              name: '経常経費',
              question: '毎月の通常の支出は何ですか？\nカテゴリーごとに金額を入力してください。',
              info: 'Include any expenses you spend regularly.',
            },
            complete: {
              name: '完了',
              question: '予算プランナーが設定されました！',
              info: 'プランナーを終了しても大丈夫です。',
            },
          },
          asset: {
            name: '資産名',
            namePlaceholder: '普通預金口座、定期口座',
            description: '説明',
            descriptionPlaceholder: '例）貯金するときはここだけに入れる',
            amount: '金額',
            currency: '通貨',
            presets: {
              checkingAccount: {
                name: '普通預金口座',
                description: '普段使い用の口座',
              },
              savingsAccount: {
                name: '定期口座',
                description: '貯金するお金はここにためていく',
              },
              sinkingFunds: {
                name: '緊急予備資金',
                description: '予想外の請求が発生したとき使える生活防衛資金',
              },
            },
          },
          revenue: {
            description: '説明',
            descriptionPlaceholder: '例）給料日に振り込まれる額',
            amount: '金額',
            currency: '通貨',
            from: '~から',
            to: '〜に',
            recurrence: '支払いスケジュール',
            presets: {
              salary: {
                description: '給料（隔週）',
                from: 'なんとか会社',
              },
              pension: {
                description: '年金（毎月）',
                from: '政府',
              },
              dividends: {
                description: '配当（四半期ごと）',
                from: 'なんとか会社',
              },
            },
          },
          liability: {
            description: '説明',
            descriptionPlaceholder: '例）給料日に振り込まれる額',
            amount: '金額',
            currency: '通貨',
            from: '~から',
            to: '〜に',
            recurrence: '支払いスケジュール',
            presets: {
              studentLoans: {
                description: '奨学金（固定金利）',
                from: 'なんとか学校',
              },
              carLoans: {
                description: '自動車ローン（固定金利）',
                from: 'なんとか信用組合',
              },
              mortgage: {
                description: '住宅ローン（変動金利）',
                from: 'なんとか銀行',
              },
            },
          },
        },
        schedules: 'スケジュール',
        accounts: 'アカウント',
        transactions: '取引',
        scheduledTransactions: '予定されている取引',
        balance: '合計',
        edit: '編集',
        delete: '削除',
        csv: 'CSVダウンロード',
        save: '保存',
        revenue: '収益',
        expense: '消費',
        profit: '利益',
        netWorth: '純資産',
        overviewForDate: '{v} の概要',
        overviewProfitLossSummary: '利益・損失',
        overviewProfit: '利益',
        addAccount: 'アカウントを追加',
        editAccount: 'アカウントを修正',
        accountName: 'アカウント名',
        accountNamePlaceholder: 'e.g. ABC 銀行 定期口座',
        accountCategory: 'カテゴリ',
        accountDescription: '説明',
        accountDescriptionPlaceholder: 'e.g. ここに頑張ってお金貯める！',
        accountBalance: '残高',
        accountCurrency: '通貨',
        accountType: '種類',
        accountCreatedAt: '作成日',
        accountUpdatedAt: '更新日',
        accountDeleteConfirmTitle: '本当にこのアカウントを削除しますか？',
        accountDeleteConfirmMessage:
          'アカウントを削除すると金額情報や\n取引記録も影響されます。\n本当に削除しますか？',
        createTransaction: '取引を作成',
        editTransaction: '取引を修正',
        transactionDescription: '説明',
        transactionDescriptionPlaceholder: 'お買い物（食材）',
        transactionAmount: '金額',
        transactionCurrency: '通貨',
        transactionSourceAccount: '〜から',
        transactionDestinationAccount: '〜に',
        transactionTransactedAt: '取引日',
        transactionTransactedAtPlaceholder: '2006-01-02 15:04:05',
        transactionCreatedAt: '作成日',
        transactionUpdatedAt: '更新日',
        transactionDeleteConfirmTitle: '本当にこの取引を削除しますか？',
        transactionDeleteConfirmMessage:
          '本当にこの取引を削除しますか？ | 本当に選択された取引　{count}件　を削除しますか?',
      },
      portfolio: {
        name: 'ポートフォリオ',
        title: "ポートフォリオ - {'@'}massivebugs",
        exit: '閉じる',
        help: 'ヘルプ',
        contact: '開発者に連絡する',
        about: 'ポートフォリオについて',
        githubLinkTitle: 'GitHubを見てみる',
        sendContactMessage: 'メッセージを送信する',
        aboutDialogTitle: "ポートフォリオ - {'@'}massivebugs について",
      },
    },
  },
}
