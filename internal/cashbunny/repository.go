package cashbunny

import (
	"context"
	"database/sql"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

// ### Account ###
type CreateAccountParams struct {
	UserID      uint32
	Category    AccountCategory
	Name        string
	Description string
	Currency    string
	OrderIndex  sql.NullInt32
}

type ListAccountsAndAmountBetweenDatesParams struct {
	FromTransactedAt time.Time
	ToTransactedAt   time.Time
	UserID           uint32
}

type ListAccountsAndAmountByCategoryParams struct {
	UserID   uint32
	Category AccountCategory
}

type ListAccountsByIDsParams struct {
	UserID uint32
	IDs    []uint32
}

type DeleteAccountParams struct {
	UserID uint32
	ID     uint32
}

type GetAccountByIDParams struct {
	UserID uint32
	ID     uint32
}

type IncrementAccountIndicesParams struct {
	UserID     uint32
	OrderIndex uint32
}

type IAccountRepository interface {
	CreateAccount(ctx context.Context, db db.DB, params CreateAccountParams) (uint32, error)
	ListAccountsAndAmountBetweenDates(ctx context.Context, db db.DB, params ListAccountsAndAmountBetweenDatesParams) ([]*Account, error)
	ListAccountsAndAmountByCategory(ctx context.Context, db db.DB, params ListAccountsAndAmountByCategoryParams) ([]*Account, error)
	ListAccountsByIDs(ctx context.Context, db db.DB, params ListAccountsByIDsParams) ([]*Account, error)
	IncrementAccountIndices(ctx context.Context, db db.DB, params IncrementAccountIndicesParams) error
	DeleteAccount(ctx context.Context, db db.DB, params DeleteAccountParams) error
	GetAccountByID(ctx context.Context, db db.DB, params GetAccountByIDParams) (*Account, error)
}

// ### Scheduled Transaction ###

type IScheduledTransactionRepository interface {
	ListScheduledTransactionsWithAllRelations(ctx context.Context, db db.DB, userID uint32) ([]*ScheduledTransaction, error)
}

// ### Transaction ###

type CreateTransactionParams struct {
	UserID                 uint32
	ScheduledTransactionID sql.NullInt32
	CategoryID             sql.NullInt32
	SrcAccountID           uint32
	DestAccountID          uint32
	Description            string
	Amount                 float64
	Currency               string
	TransactedAt           time.Time
}

type DeleteTransactionParams struct {
	UserID uint32
	ID     uint32
}

type DeleteTransactionsByAccountIDParams struct {
	UserID    uint32
	AccountID uint32
}

type GetTransactionByIDParams struct {
	UserID uint32
	ID     uint32
}

type ListTransactionsBetweenDatesParams struct {
	UserID           uint32
	FromTransactedAt time.Time
	ToTransactedAt   time.Time
}

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, db db.DB, params CreateTransactionParams) (uint32, error)
	DeleteTransaction(ctx context.Context, db db.DB, params DeleteTransactionParams) error
	DeleteTransactionsByAccountID(ctx context.Context, db db.DB, params DeleteTransactionsByAccountIDParams) error
	GetTransactionByID(ctx context.Context, db db.DB, params GetTransactionByIDParams) (*Transaction, error)
	ListTransactions(ctx context.Context, db db.DB, userID uint32) ([]*Transaction, error)
	ListTransactionsBetweenDates(ctx context.Context, db db.DB, params ListTransactionsBetweenDatesParams) ([]*Transaction, error)
}

// ### Transaction Category ###

type CreateTransactionCategoryParams struct {
	UserID uint32
	Name   string
}

type ITransactionCategoryRepository interface {
	CreateTransactionCategory(ctx context.Context, db db.DB, arg CreateTransactionCategoryParams) (uint32, error)
	ListTransactionCategoriesWithAllocations(ctx context.Context, db db.DB, userID uint32) ([]*TransactionCategory, error)
}

// ### Transaction Category Allocation ###

type TransactionCategoryAllocation struct {
	ID         uint32
	CategoryID uint32
	Amount     *money.Money
}

func NewTransactionCategoryAllocationFromQueries(data *queries.CashbunnyTransactionCategoryAllocation) *TransactionCategoryAllocation {
	return &TransactionCategoryAllocation{
		ID:         data.ID,
		CategoryID: data.CategoryID,
		Amount:     money.NewFromFloat(data.Amount, data.Currency),
	}
}

// ### Recurrence Rule ###

type CreateRecurrenceRuleParams struct {
	Freq     string
	Dtstart  time.Time
	Count    int32
	Interval int32
	Until    time.Time
}

type IRecurrenceRuleRepository interface {
	CreateRecurrenceRule(ctx context.Context, db db.DB, params CreateRecurrenceRuleParams) (uint32, error)
}

// ### Currency ###

type ICurrencyRepository interface{}

// ### User Preferences ###

type CreateUserCurrencyParams struct {
	UserID       uint32
	CurrencyCode string
}

type IUserPreferencesRepository interface {
	GetUserPreferenceExistsByUserID(ctx context.Context, db db.DB, userID uint32) (bool, error)
	CreateUserPreferences(ctx context.Context, db db.DB, userID uint32) (uint32, error)
	GetUserPreferencesByUserID(ctx context.Context, db db.DB, userID uint32) (*UserPreferences, error)
	CreateUserCurrency(ctx context.Context, db db.DB, params CreateUserCurrencyParams) (uint32, error)
	ListUserCurrencies(ctx context.Context, db db.DB, userID uint32) ([]string, error)
}
