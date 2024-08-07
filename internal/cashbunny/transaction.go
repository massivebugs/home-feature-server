package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
)

type Transaction struct {
	ID           uint32
	Description  string
	Amount       *money.Money
	TransactedAt time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time

	SourceAccount      *Account
	DestinationAccount *Account
}

func NewTransaction(transaction *cashbunny_repository.CashbunnyTransaction, srcAccount *cashbunny_repository.CashbunnyAccount, destAccount *cashbunny_repository.CashbunnyAccount) (*Transaction, error) {
	sa, err := NewAccount(srcAccount)
	if err != nil {
		return nil, err
	}

	da, err := NewAccount(destAccount)
	if err != nil {
		return nil, err
	}

	t := &Transaction{

		ID:           transaction.ID,
		Description:  transaction.Description,
		Amount:       money.NewFromFloat(transaction.Amount, transaction.Currency),
		TransactedAt: transaction.TransactedAt,
		CreatedAt:    transaction.CreatedAt,
		UpdatedAt:    transaction.UpdatedAt,

		SourceAccount:      sa,
		DestinationAccount: da,
	}

	return t, t.validate()
}

func (e *Transaction) validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(
			&e.ID,
			validation.Required,
		),
		validation.Field(
			&e.Description,
			validation.Required,
		),
		validation.Field(
			&e.Amount,
			validation.Required,
			validation.By(IsMoneyNotNegative(e.Amount)),
		),
		validation.Field(
			&e.TransactedAt,
			validation.Required,
		),
		validation.Field(
			&e.SourceAccount,
			validation.Required,
		),
		validation.Field(
			&e.DestinationAccount,
			validation.Required,
		),
	)
}
