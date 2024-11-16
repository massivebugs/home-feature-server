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

	RecurrenceRule      *RecurrenceRule
	SourceAccount       *Account
	DestinationAccount  *Account
	TransactionCategory *TransactionCategory
}

func NewScheduledTransactionFromQueries(
	stData *queries.CashbunnyScheduledTransaction,
	rrData *queries.CashbunnyRecurrenceRule,
	sAData *queries.CashbunnyAccount,
	dAData *queries.CashbunnyAccount,
	trcaData *queries.CashbunnyTransactionCategory,
) (*ScheduledTransaction, error) {
	rr, err := NewRecurrenceRuleFromQueries(rrData)
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
		result.SourceAccount = NewAccountFromQueries(sAData, nil)
	}

	if dAData != nil {
		result.DestinationAccount = NewAccountFromQueries(dAData, nil)
	}

	if trcaData != nil {
		result.TransactionCategory = NewTransactionCategoryFromQueries(trcaData)
	}

	return result, nil
}

func (st *ScheduledTransaction) toTransactions(from time.Time, to time.Time) []*Transaction {
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

func (str *ScheduledTransaction) isRevenueTransaction() bool {
	return str.SourceAccount.Category == AccountCategoryRevenues && str.DestinationAccount.Category == AccountCategoryAssets
}

func (str *ScheduledTransaction) isLiabilityTransaction() bool {
	return str.SourceAccount.Category == AccountCategoryAssets && str.DestinationAccount.Category == AccountCategoryLiabilities
}

func (str *ScheduledTransaction) isExpenseTransaction() bool {
	return str.SourceAccount.Category == AccountCategoryAssets && str.DestinationAccount.Category == AccountCategoryExpenses
}
