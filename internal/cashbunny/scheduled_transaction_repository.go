package cashbunny

import (
	"context"
	"database/sql"

	"github.com/massivebugs/home-feature-server/db"
)

type CreateScheduledTransactionParams struct {
	UserID        uint32
	CategoryID    sql.NullInt32
	SrcAccountID  uint32
	DestAccountID uint32
	Description   string
	Amount        float64
	Currency      string
}

type CreateScheduledTransactionRecurrenceRuleRelationshipParams struct {
	ScheduledTransactionID uint32
	RecurrenceRuleID       uint32
}

type CreateScheduledTransactionWithRecurrenceRuleParams struct {
	ScheduledTransaction CreateScheduledTransactionParams
	RecurrenceRule       CreateRecurrenceRuleParams
}

type IScheduledTransactionRepository interface {
	ListScheduledTransactionsWithAllRelations(ctx context.Context, db db.DB, userID uint32) ([]*ScheduledTransaction, error)
}
