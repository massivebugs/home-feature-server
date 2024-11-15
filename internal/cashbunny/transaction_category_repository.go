package cashbunny

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type CreateTransactionCategoryParams struct {
	UserID uint32
	Name   string
}

type ITransactionCategoryRepository interface {
	CreateTransactionCategory(ctx context.Context, db db.DB, arg CreateTransactionCategoryParams) (uint32, error)
	ListTransactionCategoriesWithAllocations(ctx context.Context, db db.DB, userID uint32) ([]*TransactionCategory, error)
}

type TransactionCategoryRepository struct {
	querier queries.Querier
}

var _ ITransactionCategoryRepository = (*TransactionCategoryRepository)(nil)

func NewTransactionCategoryRepository(querier queries.Querier) *TransactionCategoryRepository {
	return &TransactionCategoryRepository{
		querier: querier,
	}
}

func (r *TransactionCategoryRepository) CreateTransactionCategory(ctx context.Context, db db.DB, arg CreateTransactionCategoryParams) (uint32, error) {
	return 0, nil
}

func (r *TransactionCategoryRepository) ListTransactionCategoriesWithAllocations(ctx context.Context, db db.DB, userID uint32) ([]*TransactionCategory, error) {
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

	res := make([]*TransactionCategory, len(catRows))
	for idx, d := range catRows {
		res[idx] = NewTransactionCategoryWithAllocationsFromQueries(d, catAlcsMap[d.ID])
	}

	return res, nil
}
