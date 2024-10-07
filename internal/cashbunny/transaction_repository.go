package cashbunny

import (
	"context"
	"database/sql"
	"time"

	"github.com/massivebugs/home-feature-server/db"
)

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
