package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type Transaction struct {
	id            uint32
	srcAccountID  uint32
	destAccountID uint32
	description   string
	amount        *money.Money
	transactedAt  time.Time
	createdAt     time.Time
	updatedAt     time.Time

	sourceAccount        *Account
	destinationAccount   *Account
	scheduledTransaction *ScheduledTransaction
}

func NewTransactionFromQueries(data *queries.CashbunnyTransaction) *Transaction {
	return &Transaction{
		id:            data.ID,
		srcAccountID:  data.SrcAccountID,
		destAccountID: data.DestAccountID,
		description:   data.Description,
		amount:        money.NewFromFloat(data.Amount, data.Currency),
		transactedAt:  data.TransactedAt,
		createdAt:     data.CreatedAt,
		updatedAt:     data.UpdatedAt,
	}
}

func (tr *Transaction) Validate() error {
	return validation.ValidateStruct(
		tr,
		validation.Field(
			&tr.id,
			validation.Required,
		),
		validation.Field(
			&tr.description,
			validation.Required,
		),
		validation.Field(
			&tr.amount,
			validation.Required,
			validation.By(isMoneyNotNegative(tr.amount)),
		),
		validation.Field(
			&tr.transactedAt,
			validation.Required,
		),
	)
}

func (tr *Transaction) IsSourceAccount(a *Account) bool {
	return tr.srcAccountID == a.id
}

func (tr *Transaction) IsDestinationAccount(a *Account) bool {
	return tr.destAccountID == a.id
}
