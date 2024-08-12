package cashbunny

type Ledger struct {
	accounts []*Account
}

func NewLedger(accounts []*Account) *Ledger {
	return &Ledger{
		accounts: accounts,
	}
}

func (l *Ledger) GetProfitLoss() (revenues CurrencySums, expenses CurrencySums, sums CurrencySums) {
	revenues = NewCurrencySums(nil)
	expenses = NewCurrencySums(nil)
	sums = NewCurrencySums(nil)

	for _, a := range l.accounts {
		sum := a.SumTransactions()

		if a.Category == AccountCategoryRevenues {
			revenues.AddSums(sum)
		} else if a.Category == AccountCategoryExpenses {
			expenses.AddSums(sum)
		}
	}

	sums.AddSums(revenues)
	sums.SubtractSums(expenses)

	return
}
