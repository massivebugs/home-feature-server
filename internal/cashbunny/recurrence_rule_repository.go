package cashbunny

import (
	"context"
	"time"

	"github.com/massivebugs/home-feature-server/db"
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
