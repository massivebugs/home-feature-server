package cashbunny

import (
	"context"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type CreateTransactionParams struct {
	UserID                 uint32
	ScheduledTransactionID *int32
	CategoryID             *int32
	SrcAccountID           uint32
	DestAccountID          uint32
	Description            string
	Amount                 float64
	Currency               string
	TransactedAt           time.Time
}

type UpdateTransactionParams struct {
	Description  string
	Amount       float64
	TransactedAt time.Time
	UserID       uint32
	ID           uint32
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
	UpdateTransaction(ctx context.Context, db db.DB, arg UpdateTransactionParams) error
	DeleteTransaction(ctx context.Context, db db.DB, params DeleteTransactionParams) error
	DeleteTransactionsByAccountID(ctx context.Context, db db.DB, params DeleteTransactionsByAccountIDParams) error
	GetTransactionByID(ctx context.Context, db db.DB, params GetTransactionByIDParams) (*Transaction, error)
	ListTransactions(ctx context.Context, db db.DB, userID uint32) ([]*Transaction, error)
	ListTransactionsBetweenDates(ctx context.Context, db db.DB, params ListTransactionsBetweenDatesParams) ([]*Transaction, error)
}

type TransactionRepository struct {
	querier queries.Querier
}

var _ ITransactionRepository = (*TransactionRepository)(nil)

func NewTransactionRepository(querier queries.Querier) *TransactionRepository {
	return &TransactionRepository{
		querier: querier,
	}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, db db.DB, params CreateTransactionParams) (uint32, error) {
	result, err := r.querier.CreateTransaction(ctx, db, queries.CreateTransactionParams{
		UserID:        params.UserID,
		SrcAccountID:  params.SrcAccountID,
		DestAccountID: params.DestAccountID,
		Description:   params.Description,
		Amount:        params.Amount,
		Currency:      params.Currency,
		TransactedAt:  params.TransactedAt,
	})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r *TransactionRepository) UpdateTransaction(ctx context.Context, db db.DB, arg UpdateTransactionParams) error {
	return r.querier.UpdateCashbunnyTransaction(ctx, db, queries.UpdateCashbunnyTransactionParams(arg))
}

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, db db.DB, params DeleteTransactionParams) error {
	return r.querier.DeleteTransaction(ctx, db, queries.DeleteTransactionParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
}

func (r *TransactionRepository) DeleteTransactionsByAccountID(ctx context.Context, db db.DB, params DeleteTransactionsByAccountIDParams) error {
	return r.querier.DeleteTransactionsByAccountID(ctx, db, queries.DeleteTransactionsByAccountIDParams{
		UserID:    params.UserID,
		AccountID: params.AccountID,
	})
}

func (r *TransactionRepository) GetTransactionByID(ctx context.Context, db db.DB, params GetTransactionByIDParams) (*Transaction, error) {
	data, err := r.querier.GetTransactionByID(ctx, db, queries.GetTransactionByIDParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
	if err != nil {
		return nil, err
	}

	return NewTransactionFromQueries(data), nil
}

func (r *TransactionRepository) ListTransactions(ctx context.Context, db db.DB, userID uint32) ([]*Transaction, error) {
	data, err := r.querier.ListTransactions(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	trs := make([]*Transaction, len(data))
	for idx, d := range data {
		trs[idx] = NewTransactionFromQueries(d)
	}

	return trs, nil
}

func (r *TransactionRepository) ListTransactionsBetweenDates(ctx context.Context, db db.DB, params ListTransactionsBetweenDatesParams) ([]*Transaction, error) {
	tRows, err := r.querier.ListTransactionsBetweenDates(
		ctx,
		db,
		queries.ListTransactionsBetweenDatesParams{
			UserID:           params.UserID,
			FromTransactedAt: params.FromTransactedAt,
			ToTransactedAt:   params.ToTransactedAt,
		},
	)
	if err != nil {
		return nil, err
	}

	ts := make([]*Transaction, len(tRows))
	for idx, row := range tRows {
		ts[idx] = NewTransactionFromQueries(row)
	}

	return ts, nil
}
