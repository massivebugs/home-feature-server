package cashbunny

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/service/cashbunny_entry"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_transaction"
)

type TransactionType string

type Transaction struct {
	ID           uint32
	UserID       uint32
	Description  string
	TransactedAt time.Time
	CreatedAt    time.Time

	Entries []Entry
}

func NewTransaction(t *cashbunny_transaction.CashbunnyTransaction, es []*cashbunny_entry.CashbunnyEntry) *Transaction {
	entries := make([]Entry, len(es))
	for idx, e := range es {
		entries[idx] = NewEntry(e)
	}

	return &Transaction{
		ID:           t.ID,
		UserID:       t.UserID,
		Description:  t.Description,
		TransactedAt: t.TransactedAt,
		CreatedAt:    t.CreatedAt,
		Entries:      entries,
	}
}
