package cashbunny

type Ledger struct {
	accounts     []*Account
	transactions []*Transaction
}

func NewLedger(accounts []*Account, transactions []*Transaction) *Ledger {
	l := &Ledger{
		accounts:     accounts,
		transactions: transactions,
	}

	l.mapAccountsAndTransactions()

	return l
}

func (l *Ledger) mapAccountsAndTransactions() {
	for _, a := range l.accounts {
		for _, tr := range l.transactions {
			if tr.IsSourceAccount(a) {
				a.AddOutgoingTransaction(tr)
			} else if tr.IsDestinationAccount(a) {
				a.AddIncomingTransaction(tr)
			}
		}
	}
}

func (l *Ledger) GetTransactions() []*Transaction {
	return l.transactions
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
