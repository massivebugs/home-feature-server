package cashbunny_repository

import (
	"context"
)

type CreateScheduledTransactionWithRecurrenceRuleParams struct {
	ScheduledTransaction CreateScheduledTransactionParams
	RecurrenceRule       CreateRecurrenceRuleParams
}

type ICashbunnyRepository interface {
	Querier
	CreateScheduledTransactionWithRecurrenceRule(ctx context.Context, tx DBTX, arg CreateScheduledTransactionWithRecurrenceRuleParams) error
}

type CashbunnyRepository struct {
	Querier
}

func NewCashbunnyRepository() *CashbunnyRepository {
	return &CashbunnyRepository{
		Querier: New(),
	}
}

func (r *CashbunnyRepository) CreateScheduledTransactionWithRecurrenceRule(ctx context.Context, tx DBTX, arg CreateScheduledTransactionWithRecurrenceRuleParams) error {
	stResult, err := r.CreateScheduledTransaction(ctx, tx, arg.ScheduledTransaction)
	if err != nil {
		return err
	}

	// Retrieve ScheduledTransaction ID
	stID, err := stResult.LastInsertId()
	if err != nil {
		return err
	}

	rrResult, err := r.CreateRecurrenceRule(ctx, tx, arg.RecurrenceRule)
	if err != nil {
		return err
	}

	// Retrieve ScheduledTransaction ID
	rrID, err := rrResult.LastInsertId()
	if err != nil {
		return err
	}

	_, err = r.CreateScheduledTransactionRecurrenceRuleRelationship(ctx, tx, CreateScheduledTransactionRecurrenceRuleRelationshipParams{
		ScheduledTransactionID: uint32(stID),
		RecurrenceRuleID:       uint32(rrID),
	})

	return err
}
