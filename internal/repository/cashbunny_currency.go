package repository

import (
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type CurrencyRepository struct {
	querier queries.Querier
}

var _ cashbunny.ICurrencyRepository = (*CurrencyRepository)(nil)

func NewCurrencyRepository(querier queries.Querier) *CurrencyRepository {
	return &CurrencyRepository{
		querier: querier,
	}
}
