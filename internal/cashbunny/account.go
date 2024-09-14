package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
)

type AccountCategory string
type AccountType string

const (
	AccountTypeCredit AccountType = "credit"
	AccountTypeDebit  AccountType = "debit"

	AccountCategoryAssets      AccountCategory = "assets"
	AccountCategoryLiabilities AccountCategory = "liabilities"
	AccountCategoryRevenues    AccountCategory = "revenues"
	AccountCategoryExpenses    AccountCategory = "expenses"
)

type Account struct {
	ID          uint32
	Category    AccountCategory
	Name        string
	Description string
	Balance     *money.Money
	CreatedAt   time.Time
	UpdatedAt   time.Time

	IncomingTransactions []*Transaction
	OutgoingTransactions []*Transaction
}

func NewAccount(account *cashbunny_repository.CashbunnyAccount) *Account {
	a := &Account{
		ID:          account.ID,
		Category:    AccountCategory(account.Category),
		Name:        account.Name,
		Description: account.Description,
		Balance:     money.NewFromFloat(account.Balance, account.Currency),
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}

	return a
}

func (a *Account) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(
			&a.ID,
			validation.Required,
		),
		validation.Field(
			&a.Category,
			validation.Required,
			validation.In(
				AccountCategoryAssets,
				AccountCategoryLiabilities,
				AccountCategoryRevenues,
				AccountCategoryExpenses,
			),
		),
		validation.Field(
			&a.Name,
			validation.Required,
		),
		validation.Field(
			&a.Balance,
			validation.Required,
			validation.By(IsMoneyNotNegative(a.Balance)),
		),
	)
}

func (a *Account) GetType() AccountType {
	switch a.Category {
	case AccountCategoryAssets, AccountCategoryExpenses:
		return AccountTypeDebit
	case AccountCategoryLiabilities, AccountCategoryRevenues:
		return AccountTypeCredit
	}
	return ""
}

func (a *Account) SumTransactions(from *time.Time, to *time.Time) CurrencySums {
	cs := NewCurrencySums(nil)

	for _, tr := range a.IncomingTransactions {
		if from != nil && tr.TransactedAt.Before(*from) {
			continue
		}

		if to != nil && tr.TransactedAt.After(*to) {
			continue
		}

		if a.GetType() == AccountTypeCredit {
			cs.Subtract(tr.Amount)
		} else {
			cs.Add(tr.Amount)
		}
	}

	for _, tr := range a.OutgoingTransactions {
		if from != nil && tr.TransactedAt.Before(*from) {
			continue
		}

		if to != nil && tr.TransactedAt.After(*to) {
			continue
		}

		if a.GetType() == AccountTypeCredit {
			cs.Add(tr.Amount)
		} else {
			cs.Subtract(tr.Amount)
		}
	}

	return cs
}

func (a *Account) AddIncomingTransaction(tr *Transaction) {
	a.IncomingTransactions = append(a.IncomingTransactions, tr)
}

func (a *Account) AddOutgoingTransaction(tr *Transaction) {
	a.OutgoingTransactions = append(a.OutgoingTransactions, tr)
}
