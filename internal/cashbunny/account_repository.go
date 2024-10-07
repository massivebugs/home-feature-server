package cashbunny

import (
	"context"
	"database/sql"
	"time"

	"github.com/massivebugs/home-feature-server/db"
)

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
