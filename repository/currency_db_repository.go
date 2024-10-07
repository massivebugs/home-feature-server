package repository

import "github.com/massivebugs/home-feature-server/db/queries"

type CurrencyDBRepository struct {
	querier queries.Querier
}

func NewCurrencyDBRepository(querier queries.Querier) *CurrencyDBRepository {
	return &CurrencyDBRepository{
		querier: querier,
	}
}
