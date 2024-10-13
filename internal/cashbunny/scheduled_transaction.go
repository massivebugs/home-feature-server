package cashbunny

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type ScheduledTransaction struct {
	id            uint32
	srcAccountID  uint32
	destAccountID uint32
	description   string
	amount        *money.Money
	createdAt     time.Time
	updatedAt     time.Time

	recurrenceRule      *RecurrenceRule
	sourceAccount       *Account
	destinationAccount  *Account
	transactionCategory *TransactionCategory
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
		id:            stData.ID,
		srcAccountID:  stData.SrcAccountID,
		destAccountID: stData.DestAccountID,
		description:   stData.Description,
		amount:        money.NewFromFloat(stData.Amount, stData.Currency),
		createdAt:     stData.CreatedAt,
		updatedAt:     stData.UpdatedAt,

		recurrenceRule: rr,
	}

	if sAData != nil {
		result.sourceAccount = NewAccountFromQueries(sAData, nil)
	}

	if dAData != nil {
		result.destinationAccount = NewAccountFromQueries(dAData, nil)
	}

	if trcaData != nil {
		result.transactionCategory = NewTransactionCategoryFromQueries(trcaData)
	}

	return result, nil
}

func (st *ScheduledTransaction) toTransactions(from time.Time, to time.Time) []*Transaction {
	var result []*Transaction

	timeSlice := st.recurrenceRule.rule.Between(from, to, true)
	for _, t := range timeSlice {
		result = append(result, &Transaction{
			srcAccountID:         st.srcAccountID,
			destAccountID:        st.destAccountID,
			description:          st.description,
			amount:               money.New(st.amount.Amount(), st.amount.Currency().Code),
			transactedAt:         t,
			sourceAccount:        st.sourceAccount,
			destinationAccount:   st.destinationAccount,
			scheduledTransaction: st,
		})
	}

	return result
}

func (str *ScheduledTransaction) isRevenueTransaction() bool {
	return str.sourceAccount.category == AccountCategoryRevenues && str.destinationAccount.category == AccountCategoryAssets
}

func (str *ScheduledTransaction) isLiabilityTransaction() bool {
	return str.sourceAccount.category == AccountCategoryAssets && str.destinationAccount.category == AccountCategoryLiabilities
}

func (str *ScheduledTransaction) isExpenseTransaction() bool {
	return str.sourceAccount.category == AccountCategoryAssets && str.destinationAccount.category == AccountCategoryExpenses
}
