package cashbunny

import (
	"context"
	"errors"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type CreateRecurrenceRuleParams struct {
	Freq     string
	Dtstart  time.Time
	Count    int32
	Interval int32
	Until    time.Time
}

type IRecurrenceRuleRepository interface {
	CreateRecurrenceRule(ctx context.Context, db db.DB, params CreateRecurrenceRuleParams) (uint32, error)
}

type RecurrenceRuleRepository struct{}

var _ IRecurrenceRuleRepository = (*RecurrenceRuleRepository)(nil)

func NewRecurrenceRuleRepository(querier queries.Querier) *RecurrenceRuleRepository {
	return &RecurrenceRuleRepository{}
}

func (r *RecurrenceRuleRepository) CreateRecurrenceRule(ctx context.Context, db db.DB, params CreateRecurrenceRuleParams) (uint32, error) {
	return 0, errors.New("not implemented yet")
}
