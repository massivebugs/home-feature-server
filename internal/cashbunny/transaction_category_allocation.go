package cashbunny

import (
	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db/queries"
)

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
