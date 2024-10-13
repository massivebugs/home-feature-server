package repository

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type RecurrenceRuleRepository struct{}

var _ cashbunny.IRecurrenceRuleRepository = (*RecurrenceRuleRepository)(nil)

func NewRecurrenceRuleRepository(querier queries.Querier) *RecurrenceRuleRepository {
	return &RecurrenceRuleRepository{}
}

func (r *RecurrenceRuleRepository) CreateRecurrenceRule(ctx context.Context, db db.DB, params cashbunny.CreateRecurrenceRuleParams) (uint32, error) {
	return 0, errors.New("not implemented yet")
}
