package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
)

type ScheduledTransaction struct {
	ID            uint32
	SrcAccountID  uint32
	DestAccountID uint32
	Description   string
	Amount        *money.Money
	CreatedAt     time.Time
	UpdatedAt     time.Time

	RecurrenceRule     *RecurrenceRule
	SourceAccount      *Account
	DestinationAccount *Account
}

func NewScheduledTransaction(
	stData *cashbunny_repository.CashbunnyScheduledTransaction,
	rrData *cashbunny_repository.CashbunnyRecurrenceRule,
	sAData *cashbunny_repository.CashbunnyAccount,
	dAData *cashbunny_repository.CashbunnyAccount,
) (*ScheduledTransaction, error) {
	rr, err := NewRecurrenceRuleFromData(rrData)
	if err != nil {
		return nil, err
	}

	result := &ScheduledTransaction{
		ID:            stData.ID,
		SrcAccountID:  stData.SrcAccountID,
		DestAccountID: stData.DestAccountID,
		Description:   stData.Description,
		Amount:        money.NewFromFloat(stData.Amount, stData.Currency),
		CreatedAt:     stData.CreatedAt,
		UpdatedAt:     stData.UpdatedAt,

		RecurrenceRule: rr,
	}

	if sAData != nil {
		result.SourceAccount = NewAccount(sAData)
	}

	if dAData != nil {
		result.DestinationAccount = NewAccount(dAData)
	}

	return result, nil
}

func (st *ScheduledTransaction) ToTransactions(from time.Time, to time.Time) []*Transaction {
	var result []*Transaction

	timeSlice := st.RecurrenceRule.Rule.Between(from, to, true)
	for _, t := range timeSlice {
		result = append(result, &Transaction{
			SrcAccountID:         st.SrcAccountID,
			DestAccountID:        st.DestAccountID,
			Description:          st.Description,
			Amount:               money.New(st.Amount.Amount(), st.Amount.Currency().Code),
			TransactedAt:         t,
			SourceAccount:        st.SourceAccount,
			DestinationAccount:   st.DestinationAccount,
			ScheduledTransaction: st,
		})
	}

	return result
}
