package cashbunny

import "github.com/massivebugs/home-feature-server/db/queries"

type TransactionCategory struct {
	id     uint32
	userID uint32
	name   string

	allocations []*TransactionCategoryAllocation
}

func NewTransactionCategoryFromQueries(data *queries.CashbunnyTransactionCategory) *TransactionCategory {
	return &TransactionCategory{
		id:     data.ID,
		userID: data.UserID,
		name:   data.Name,
	}
}

func NewTransactionCategoryWithAllocationsFromQueries(categoryData *queries.CashbunnyTransactionCategory, allocationsData []*queries.CashbunnyTransactionCategoryAllocation) *TransactionCategory {
	e := NewTransactionCategoryFromQueries(categoryData)

	e.allocations = make([]*TransactionCategoryAllocation, len(allocationsData))
	for idx, d := range allocationsData {
		e.allocations[idx] = NewTransactionCategoryAllocationFromQueries(d)
	}

	return e
}
