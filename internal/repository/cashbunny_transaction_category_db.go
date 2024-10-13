package repository

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type TransactionCategoryDBRepository struct {
	querier queries.Querier
}

func NewTransactionCategoryDBRepository(querier queries.Querier) *TransactionCategoryDBRepository {
	return &TransactionCategoryDBRepository{
		querier: querier,
	}
}

func (r *TransactionCategoryDBRepository) CreateTransactionCategory(ctx context.Context, db db.DB, arg cashbunny.CreateTransactionCategoryParams) (uint32, error) {
	return 0, nil
}

func (r *TransactionCategoryDBRepository) ListTransactionCategoriesWithAllocations(ctx context.Context, db db.DB, userID uint32) ([]*cashbunny.TransactionCategory, error) {
	// Preload category allocations
	alcRows, err := r.querier.ListTransactionCategoryAllocations(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	catRows, err := r.querier.ListTransactionCategories(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	catAlcsMap := make(map[uint32][]*queries.CashbunnyTransactionCategoryAllocation, len(catRows))
	for _, alc := range alcRows {
		catAlcsMap[alc.CategoryID] = append(catAlcsMap[alc.CategoryID], alc)
	}

	res := make([]*cashbunny.TransactionCategory, len(catRows))
	for idx, d := range catRows {
		res[idx] = cashbunny.NewTransactionCategoryWithAllocationsFromQueries(d, catAlcsMap[d.ID])
	}

	return res, nil
}
