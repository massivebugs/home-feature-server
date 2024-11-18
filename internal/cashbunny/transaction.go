package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type Transaction struct {
	ID            uint32
	SrcAccountID  uint32
	DestAccountID uint32
	Description   string
	Amount        *money.Money
	TransactedAt  time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time

	SourceAccount        *Account
	DestinationAccount   *Account
	ScheduledTransaction *ScheduledTransaction
}

func NewTransactionFromQueries(data *queries.CashbunnyTransaction) *Transaction {
	return &Transaction{
		ID:            data.ID,
		SrcAccountID:  data.SrcAccountID,
		DestAccountID: data.DestAccountID,
		Description:   data.Description,
		Amount:        money.NewFromFloat(data.Amount, data.Currency),
		TransactedAt:  data.TransactedAt,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}
}

func (tr *Transaction) IsSourceAccount(a *Account) bool {
	return tr.SrcAccountID == a.ID
}

func (tr *Transaction) IsDestinationAccount(a *Account) bool {
	return tr.DestAccountID == a.ID
}
