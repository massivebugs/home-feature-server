package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	validation "github.com/go-ozzo/ozzo-validation"
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

func NewTransactionFromDBGateway(transaction *queries.CashbunnyTransaction) *Transaction {
	return &Transaction{
		ID:            transaction.ID,
		SrcAccountID:  transaction.SrcAccountID,
		DestAccountID: transaction.DestAccountID,
		Description:   transaction.Description,
		Amount:        money.NewFromFloat(transaction.Amount, transaction.Currency),
		TransactedAt:  transaction.TransactedAt,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func (tr *Transaction) Validate() error {
	return validation.ValidateStruct(
		tr,
		validation.Field(
			&tr.ID,
			validation.Required,
		),
		validation.Field(
			&tr.Description,
			validation.Required,
		),
		validation.Field(
			&tr.Amount,
			validation.Required,
			validation.By(IsMoneyNotNegative(tr.Amount)),
		),
		validation.Field(
			&tr.TransactedAt,
			validation.Required,
		),
	)
}

func (tr *Transaction) IsSourceAccount(a *Account) bool {
	return tr.SrcAccountID == a.ID
}

func (tr *Transaction) IsDestinationAccount(a *Account) bool {
	return tr.DestAccountID == a.ID
}
