package cashbunny

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/teambition/rrule-go"
)

type RecurrenceRule struct {
	rule *rrule.RRule
}

func NewRecurrenceRuleWithDefaultParams(freq rrule.Frequency, dtStart time.Time) (*RecurrenceRule, error) {
	rr, err := rrule.NewRRule(rrule.ROption{
		Freq:    freq,
		Dtstart: dtStart,
	})
	if err != nil {
		return nil, err
	}

	return &RecurrenceRule{
		rule: rr,
	}, err
}

func NewRecurrenceRuleFromQueries(data *queries.CashbunnyRecurrenceRule) (*RecurrenceRule, error) {
	rrFreq, err := rrule.StrToFreq(data.Freq)
	if err != nil {
		return nil, err
	}

	rr, err := rrule.NewRRule(rrule.ROption{
		Freq:     rrFreq,
		Dtstart:  data.Dtstart,
		Count:    int(data.Count),
		Interval: int(data.Interval),
		Until:    data.Until,
	})
	if err != nil {
		return nil, err
	}

	return &RecurrenceRule{
		rule: rr,
	}, nil
}
