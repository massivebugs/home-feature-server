package cashbunny

import (
	"errors"
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

func GetAccountTypeForCategory(c AccountCategory) (AccountType, error) {
	switch c {
	case AccountCategoryAssets, AccountCategoryExpenses:
		return AccountTypeDebit, nil
	case AccountCategoryLiabilities, AccountCategoryRevenues:
		return AccountTypeCredit, nil
	}
	return "", errors.New("failed to get account type as account category is invalid")
}

type Account struct {
	ID          uint32
	Category    AccountCategory
	Name        string
	Description string
	Balance     *money.Money
	Type        AccountType
	CreatedAt   time.Time
	UpdatedAt   time.Time

	IncomingTransactions []*Transaction
	OutgoingTransactions []*Transaction
}

func NewAccount(account *cashbunny_repository.CashbunnyAccount) (*Account, error) {
	a := &Account{
		ID:          account.ID,
		Category:    AccountCategory(account.Category),
		Name:        account.Name,
		Description: account.Description,
		Balance:     money.NewFromFloat(account.Balance, account.Currency),
		Type:        AccountType(account.Type),
		CreatedAt:   account.CreatedAt,
		UpdatedAt:   account.UpdatedAt,
	}

	return a, a.validate()
}

func (a *Account) validate() error {
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
		validation.Field(
			&a.Type,
			validation.Required,
			validation.In(AccountTypeCredit, AccountTypeDebit),
		),
	)
}

func (a *Account) SumTransactions() CurrencySums {
	cs := NewCurrencySums(nil)

	for _, tr := range a.IncomingTransactions {
		if a.Type == AccountTypeCredit {
			cs.Subtract(tr.Amount)
		} else {
			cs.Add(tr.Amount)
		}
	}

	for _, tr := range a.OutgoingTransactions {
		if a.Type == AccountTypeCredit {
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
