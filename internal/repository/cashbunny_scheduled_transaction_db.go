package repository

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type ScheduledTransactionDBRepository struct {
	querier queries.Querier
}

func NewScheduledTransactionDBRepository(querier queries.Querier) *ScheduledTransactionDBRepository {
	return &ScheduledTransactionDBRepository{
		querier: querier,
	}
}

func (r *ScheduledTransactionDBRepository) ListScheduledTransactionsWithAllRelations(ctx context.Context, db db.DB, userID uint32) ([]*cashbunny.ScheduledTransaction, error) {
	stListRows, err := r.querier.ListScheduledTransactionsWithAllRelations(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	var strs []*cashbunny.ScheduledTransaction
	for _, row := range stListRows {
		str, err := cashbunny.NewScheduledTransactionFromQueries(
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
