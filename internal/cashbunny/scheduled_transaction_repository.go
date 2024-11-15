package cashbunny

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type IScheduledTransactionRepository interface {
	ListScheduledTransactionsWithAllRelations(ctx context.Context, db db.DB, userID uint32) ([]*ScheduledTransaction, error)
}

type ScheduledTransactionRepository struct {
	querier queries.Querier
}

var _ IScheduledTransactionRepository = (*ScheduledTransactionRepository)(nil)

func NewScheduledTransactionRepository(querier queries.Querier) *ScheduledTransactionRepository {
	return &ScheduledTransactionRepository{
		querier: querier,
	}
}

func (r *ScheduledTransactionRepository) ListScheduledTransactionsWithAllRelations(ctx context.Context, db db.DB, userID uint32) ([]*ScheduledTransaction, error) {
	stListRows, err := r.querier.ListScheduledTransactionsWithAllRelations(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	var strs []*ScheduledTransaction
	for _, row := range stListRows {
		str, err := NewScheduledTransactionFromQueries(
			&row.CashbunnyScheduledTransaction,
			&row.CashbunnyRecurrenceRule,
			&row.CashbunnyAccount,
			&row.CashbunnyAccount_2,
			&row.CashbunnyTransactionCategory,
		)
		if err != nil {
			return nil, err
		}

		strs = append(strs, str)
	}

	return strs, nil
}
