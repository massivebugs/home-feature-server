package repository

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type TransactionDBRepository struct {
	querier queries.Querier
}

func NewTransactionDBRepository(querier queries.Querier) *TransactionDBRepository {
	return &TransactionDBRepository{
		querier: querier,
	}
}

func (r *TransactionDBRepository) CreateTransaction(ctx context.Context, db db.DB, params cashbunny.CreateTransactionParams) (uint32, error) {
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

func (r *TransactionDBRepository) DeleteTransaction(ctx context.Context, db db.DB, params cashbunny.DeleteTransactionParams) error {
	return r.querier.DeleteTransaction(ctx, db, queries.DeleteTransactionParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
}

func (r *TransactionDBRepository) DeleteTransactionsByAccountID(ctx context.Context, db db.DB, params cashbunny.DeleteTransactionsByAccountIDParams) error {
	return r.querier.DeleteTransactionsByAccountID(ctx, db, queries.DeleteTransactionsByAccountIDParams{
		UserID:    params.UserID,
		AccountID: params.AccountID,
	})
}

func (r *TransactionDBRepository) GetTransactionByID(ctx context.Context, db db.DB, params cashbunny.GetTransactionByIDParams) (*cashbunny.Transaction, error) {
	data, err := r.querier.GetTransactionByID(ctx, db, queries.GetTransactionByIDParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
	if err != nil {
		return nil, err
	}

	return cashbunny.NewTransactionFromQueries(data), nil
}

func (r *TransactionDBRepository) ListTransactions(ctx context.Context, db db.DB, userID uint32) ([]*cashbunny.Transaction, error) {
	data, err := r.querier.ListTransactions(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	trs := make([]*cashbunny.Transaction, len(data))
	for idx, d := range data {
		trs[idx] = cashbunny.NewTransactionFromQueries(d)
	}

	return trs, nil
}

func (r *TransactionDBRepository) ListTransactionsBetweenDates(ctx context.Context, db db.DB, params cashbunny.ListTransactionsBetweenDatesParams) ([]*cashbunny.Transaction, error) {
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

	ts := make([]*cashbunny.Transaction, len(tRows))
	for idx, row := range tRows {
		ts[idx] = cashbunny.NewTransactionFromQueries(row)
	}

	return ts, nil
}
