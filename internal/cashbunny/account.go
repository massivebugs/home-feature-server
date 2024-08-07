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
	AccountCategoryRevenue     AccountCategory = "revenue"
	AccountCategoryExpenses    AccountCategory = "expenses"
)

type Account struct {
	ID          uint32
	Category    AccountCategory
	Name        string
	Description string
	Balance     *money.Money
	Type        AccountType
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

func (e *Account) validate() error {
	return validation.ValidateStruct(
		e,
		validation.Field(
			&e.ID,
			validation.Required,
		),
		validation.Field(
			&e.Category,
			validation.Required,
			validation.In(
				AccountCategoryAssets,
				AccountCategoryLiabilities,
				AccountCategoryRevenue,
				AccountCategoryExpenses,
			),
		),
		validation.Field(
			&e.Name,
			validation.Required,
		),
		validation.Field(
			&e.Balance,
			validation.Required,
			validation.By(IsMoneyNotNegative(e.Balance)),
		),
		validation.Field(
			&e.Type,
			validation.Required,
			validation.In(AccountTypeCredit, AccountTypeDebit),
		),
	)
}
