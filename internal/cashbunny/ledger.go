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
				tr.SourceAccount = a
				a.AddOutgoingTransaction(tr)
			} else if tr.IsDestinationAccount(a) {
				tr.DestinationAccount = a
				a.AddIncomingTransaction(tr)
			}
		}
	}
}

func (l *Ledger) GetTransactions(from *time.Time, to *time.Time) []*Transaction {
	var result []*Transaction
	for _, tr := range l.transactions {
		if from != nil && tr.TransactedAt.Before(*from) {
			continue
		}
		if to != nil && tr.TransactedAt.After(*to) {
			continue
		}
		result = append(result, tr)
	}
	return result
}

func (l *Ledger) GetNetWorth(to *time.Time) CurrencySums {
	sums := NewCurrencySums(nil)

	for _, a := range l.accounts {
		sum := a.SumTransactions(nil, to)

		if a.Category == AccountCategoryAssets {
			sums.AddSums(sum)
		} else if a.Category == AccountCategoryLiabilities {
			sums.SubtractSums(sum)
		}
	}

	return sums
}

func (l *Ledger) GetProfitLoss(from *time.Time, to *time.Time) (revenues CurrencySums, expenses CurrencySums, sums CurrencySums) {
	revenues = NewCurrencySums(nil)
	expenses = NewCurrencySums(nil)
	sums = NewCurrencySums(nil)

	for _, a := range l.accounts {
		sum := a.SumTransactions(from, to)

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

func (l *Ledger) GetAccountsByCategory(c AccountCategory) []*Account {
	var res []*Account
	for _, a := range l.accounts {
		if a.Category == c {
			res = append(res, a)
		}
	}

	return res
}
