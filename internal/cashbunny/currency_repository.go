package cashbunny

import "github.com/massivebugs/home-feature-server/db/queries"

type ICurrencyRepository interface{}

type CurrencyRepository struct {
	querier queries.Querier
}

var _ ICurrencyRepository = (*CurrencyRepository)(nil)

func NewCurrencyRepository(querier queries.Querier) *CurrencyRepository {
	return &CurrencyRepository{
		querier: querier,
	}
}
