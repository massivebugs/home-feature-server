package cashbunny

import "time"

type summary map[string]struct {
	Expense string `json:"expense"`
	Profit  string `json:"profit"`
	Revenue string `json:"revenue"`
}

type overview struct {
	From                      time.Time         `json:"from"`
	To                        time.Time         `json:"to"`
	NetWorth                  map[string]string `json:"net_worth"`
	ProfitLossSummary         summary           `json:"profit_loss_summary"`
	AssetAccounts             []*Account        `json:"asset_accounts"`
	LiabilityAccounts         []*Account        `json:"liability_accounts"`
	Transactions              []*Transaction    `json:"transactions"`
	TransactionsFromScheduled []*Transaction    `json:"transactions_from_scheduled"`
}

func newOverview(from time.Time, to time.Time, ledger *Ledger, transactionsFromScheduled []*Transaction) overview {
	netWorth := map[string]string{}
	for k, money := range ledger.getNetWorth(&to) {
		netWorth[k] = money.Display()
	}

	revenues, expenses, sums := ledger.getProfitLoss(&from, &to)
	result := summary{}
	for k, money := range revenues {
		values := result[k]
		values.Revenue = money.Display()
		result[k] = values
	}

	for k, money := range expenses {
		values := result[k]
		values.Expense = money.Display()
		result[k] = values
	}

	for k, money := range sums {
		values := result[k]
		values.Profit = money.Display()
		result[k] = values
	}

	return overview{
		From:                      from,
		To:                        to,
		NetWorth:                  netWorth,
		ProfitLossSummary:         result,
		AssetAccounts:             ledger.getAccountsByCategory(AccountCategoryAssets),
		LiabilityAccounts:         ledger.getAccountsByCategory(AccountCategoryLiabilities),
		Transactions:              ledger.getTransactions(&from, &to),
		TransactionsFromScheduled: transactionsFromScheduled,
	}
}
