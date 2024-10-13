package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/massivebugs/home-feature-server/db/queries"
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
	id          uint32
	category    AccountCategory
	name        string
	description string
	currency    string
	createdAt   time.Time
	updatedAt   time.Time

	incomingTransactions []*Transaction
	outgoingTransactions []*Transaction
	amount               *money.Money
}

func NewAccountFromQueries(account *queries.CashbunnyAccount, amount *float64) *Account {
	a := &Account{
		id:          account.ID,
		category:    AccountCategory(account.Category),
		name:        account.Name,
		description: account.Description,
		currency:    account.Currency,
		createdAt:   account.CreatedAt,
		updatedAt:   account.UpdatedAt,
	}

	if amount != nil {
		a.amount = money.NewFromFloat(*amount, a.currency)
	}

	return a
}

func (a *Account) validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(
			&a.id,
			validation.Required,
		),
		validation.Field(
			&a.category,
			validation.Required,
			validation.In(
				AccountCategoryAssets,
				AccountCategoryLiabilities,
				AccountCategoryRevenues,
				AccountCategoryExpenses,
			),
		),
		validation.Field(
			&a.name,
			validation.Required,
		),
		validation.Field(
			&a.currency,
			validation.Required,
		),
	)
}

func (a *Account) getType() AccountType {
	switch a.category {
	case AccountCategoryAssets, AccountCategoryExpenses:
		return AccountTypeDebit
	case AccountCategoryLiabilities, AccountCategoryRevenues:
		return AccountTypeCredit
	}
	return ""
}

func (a *Account) sumTransactions(from *time.Time, to *time.Time) CurrencySums {
	cs := NewCurrencySums(nil)

	for _, tr := range a.incomingTransactions {
		if from != nil && tr.transactedAt.Before(*from) {
			continue
		}

		if to != nil && tr.transactedAt.After(*to) {
			continue
		}

		if a.getType() == AccountTypeCredit {
			cs.subtract(tr.amount)
		} else {
			cs.add(tr.amount)
		}
	}

	for _, tr := range a.outgoingTransactions {
		if from != nil && tr.transactedAt.Before(*from) {
			continue
		}

		if to != nil && tr.transactedAt.After(*to) {
			continue
		}

		if a.getType() == AccountTypeCredit {
			cs.add(tr.amount)
		} else {
			cs.subtract(tr.amount)
		}
	}

	return cs
}

func (a *Account) addIncomingTransaction(tr *Transaction) {
	a.incomingTransactions = append(a.incomingTransactions, tr)
}

func (a *Account) addOutgoingTransaction(tr *Transaction) {
	a.outgoingTransactions = append(a.outgoingTransactions, tr)
}
