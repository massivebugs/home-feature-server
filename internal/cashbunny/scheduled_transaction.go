package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db/queries"
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

func NewScheduledTransactionFromDBGateway(
	stData *queries.CashbunnyScheduledTransaction,
	rrData *queries.CashbunnyRecurrenceRule,
	sAData *queries.CashbunnyAccount,
	dAData *queries.CashbunnyAccount,
) (*ScheduledTransaction, error) {
	rr, err := NewRecurrenceRuleFromDBGateway(rrData)
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
		result.SourceAccount = NewAccountFromDBGateway(sAData, nil)
	}

	if dAData != nil {
		result.DestinationAccount = NewAccountFromDBGateway(dAData, nil)
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

func (str *ScheduledTransaction) IsRevenueTransaction() bool {
	return str.SourceAccount.Category == AccountCategoryRevenues && str.DestinationAccount.Category == AccountCategoryAssets
}

func (str *ScheduledTransaction) IsLiabilityTransaction() bool {
	return str.SourceAccount.Category == AccountCategoryAssets && str.DestinationAccount.Category == AccountCategoryLiabilities
}
