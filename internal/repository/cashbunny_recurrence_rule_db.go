package repository

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type RecurrenceRuleDBRepository struct{}

func NewRecurrenceRuleDBRepository(querier queries.Querier) *RecurrenceRuleDBRepository {
	return &RecurrenceRuleDBRepository{}
}

func (r *RecurrenceRuleDBRepository) CreateRecurrenceRule(ctx context.Context, db db.DB, params cashbunny.CreateRecurrenceRuleParams) (uint32, error) {
	return 0, errors.New("not implemented yet")
}
