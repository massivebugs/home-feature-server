package cashbunny

import (
	"time"
)

type Ledger struct {
	accounts     []*Account
	transactions []*Transaction
}

// Transactions should be ordered ASC by transactedAt
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
				tr.sourceAccount = a
				a.addOutgoingTransaction(tr)
			} else if tr.IsDestinationAccount(a) {
				tr.destinationAccount = a
				a.addIncomingTransaction(tr)
			}
		}
	}
}

func (l *Ledger) getTransactions(from *time.Time, to *time.Time) []*Transaction {
	var result []*Transaction
	for _, tr := range l.transactions {
		if from != nil && tr.transactedAt.Before(*from) {
			continue
		}
		if to != nil && tr.transactedAt.After(*to) {
			continue
		}
		result = append(result, tr)
	}
	return result
}

func (l *Ledger) getNetWorth(to *time.Time) CurrencySums {
	sums := NewCurrencySums(nil)

	for _, a := range l.accounts {
		sum := a.sumTransactions(nil, to)

		if a.category == AccountCategoryAssets {
			sums.addSums(sum)
		} else if a.category == AccountCategoryLiabilities {
			sums.subtractSums(sum)
		}
	}

	return sums
}

func (l *Ledger) getProfitLoss(from *time.Time, to *time.Time) (revenues CurrencySums, expenses CurrencySums, sums CurrencySums) {
	revenues = NewCurrencySums(nil)
	expenses = NewCurrencySums(nil)
	sums = NewCurrencySums(nil)

	for _, a := range l.accounts {
		sum := a.sumTransactions(from, to)

		if a.category == AccountCategoryRevenues {
			revenues.addSums(sum)
		} else if a.category == AccountCategoryExpenses {
			expenses.addSums(sum)
		}
	}

	sums.addSums(revenues)
	sums.subtractSums(expenses)

	return
}

func (l *Ledger) getAccountsByCategory(c AccountCategory) []*Account {
	var res []*Account
	for _, a := range l.accounts {
		if a.category == c {
			res = append(res, a)
		}
	}

	return res
}
