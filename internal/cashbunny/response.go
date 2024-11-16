package cashbunny

import (
	"strconv"
	"time"
)

type planResponse struct {
}

func newPlanResponse(_ *Planner) planResponse {
	return planResponse{}
}

type plannerAssetResponse struct {
	AssetAccountID string  `json:"asset_account_id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
}

type plannerRevenueResponse struct {
	ScheduledTransactionID      string                       `json:"scheduled_transaction_id"`
	Description                 string                       `json:"description"`
	Amount                      float64                      `json:"amount"`
	Currency                    string                       `json:"currency"`
	SourceRevenueAccountID      string                       `json:"source_revenue_account_id"`
	SourceRevenueAccountName    string                       `json:"source_revenue_account_name"`
	DestinationAssetAccountID   string                       `json:"destination_asset_account_id"`
	DestinationAssetAccountName string                       `json:"destination_asset_account_name"`
	RecurrenceRule              recurrenceRuleResponse       `json:"recurrence_rule"`
	TransactionCategory         *transactionCategoryResponse `json:"transaction_category"`
}

type plannerLiabilityResponse struct {
	ScheduledTransactionID          string                       `json:"scheduled_transaction_id"`
	Description                     string                       `json:"description"`
	Amount                          float64                      `json:"amount"`
	Currency                        string                       `json:"currency"`
	SourceAssetAccountID            string                       `json:"source_asset_account_id"`
	SourceAssetAccountName          string                       `json:"source_asset_account_name"`
	DestinationLiabilityAccountID   string                       `json:"destination_liability_account_id"`
	DestinationLiabilityAccountName string                       `json:"destination_liability_account_name"`
	RecurrenceRule                  recurrenceRuleResponse       `json:"recurrence_rule"`
	TransactionCategory             *transactionCategoryResponse `json:"transaction_category"`
}

type plannerExpenseResponse struct {
	ScheduledTransactionID        string                       `json:"scheduled_transaction_id"`
	Description                   string                       `json:"description"`
	Amount                        float64                      `json:"amount"`
	Currency                      string                       `json:"currency"`
	SourceAssetAccountID          string                       `json:"source_asset_account_id"`
	SourceAssetAccountName        string                       `json:"source_asset_account_name"`
	DestinationExpenseAccountID   string                       `json:"destination_expense_account_id"`
	DestinationExpenseAccountName string                       `json:"destination_expense_account_name"`
	RecurrenceRule                recurrenceRuleResponse       `json:"recurrence_rule"`
	TransactionCategory           *transactionCategoryResponse `json:"transaction_category"`
}

type plannerParametersResponse struct {
	Assets                []plannerAssetResponse        `json:"assets"`
	Revenues              []plannerRevenueResponse      `json:"revenues"`
	Liabilities           []plannerLiabilityResponse    `json:"liabilities"`
	Expenses              []plannerExpenseResponse      `json:"expenses"`
	TransactionCategories []transactionCategoryResponse `json:"transaction_categories"`
}

func newPlannerParametersResponse(
	assetAccounts []*Account,
	scheduledRevenueTransactions []*ScheduledTransaction,
	scheduledLiabilityTransactions []*ScheduledTransaction,
	scheduledExpenseTransactions []*ScheduledTransaction,
	transactionCategories []*TransactionCategory,
) plannerParametersResponse {
	var assets []plannerAssetResponse
	for _, aa := range assetAccounts {
		assets = append(assets, plannerAssetResponse{
			AssetAccountID: strconv.FormatInt(int64(aa.ID), 10),
			Name:           aa.Name,
			Description:    aa.Description,
			Amount:         aa.Amount.AsMajorUnits(),
			Currency:       aa.Currency,
		})
	}

	revenues := make([]plannerRevenueResponse, len(scheduledRevenueTransactions))
	for idx, srt := range scheduledRevenueTransactions {
		e := plannerRevenueResponse{
			ScheduledTransactionID:      strconv.FormatInt(int64(srt.ID), 10),
			Description:                 srt.Description,
			Amount:                      srt.Amount.AsMajorUnits(),
			Currency:                    srt.Amount.Currency().Code,
			SourceRevenueAccountID:      strconv.FormatInt(int64(srt.SrcAccountID), 10),
			SourceRevenueAccountName:    srt.SourceAccount.Name,
			DestinationAssetAccountID:   strconv.FormatInt(int64(srt.DestAccountID), 10),
			DestinationAssetAccountName: srt.DestinationAccount.Name,
			RecurrenceRule:              newRecurrenceRuleResponse(srt.RecurrenceRule),
		}

		if srt.TransactionCategory != nil {
			trcRes := newTransactionCategoryResponse(srt.TransactionCategory)
			e.TransactionCategory = &trcRes
		}

		revenues[idx] = e
	}

	liabilities := make([]plannerLiabilityResponse, len(scheduledLiabilityTransactions))
	for idx, srt := range scheduledLiabilityTransactions {
		e := plannerLiabilityResponse{
			ScheduledTransactionID:          strconv.FormatInt(int64(srt.ID), 10),
			Description:                     srt.Description,
			Amount:                          srt.Amount.AsMajorUnits(),
			Currency:                        srt.Amount.Currency().Code,
			SourceAssetAccountID:            strconv.FormatInt(int64(srt.SrcAccountID), 10),
			SourceAssetAccountName:          srt.SourceAccount.Name,
			DestinationLiabilityAccountID:   strconv.FormatInt(int64(srt.DestAccountID), 10),
			DestinationLiabilityAccountName: srt.DestinationAccount.Name,
			RecurrenceRule:                  newRecurrenceRuleResponse(srt.RecurrenceRule),
		}

		if srt.TransactionCategory != nil {
			trcRes := newTransactionCategoryResponse(srt.TransactionCategory)
			e.TransactionCategory = &trcRes
		}

		liabilities[idx] = e
	}

	expenses := make([]plannerExpenseResponse, len(scheduledExpenseTransactions))
	for idx, srt := range scheduledExpenseTransactions {
		e := plannerExpenseResponse{
			ScheduledTransactionID:        strconv.FormatInt(int64(srt.ID), 10),
			Description:                   srt.Description,
			Amount:                        srt.Amount.AsMajorUnits(),
			Currency:                      srt.Amount.Currency().Code,
			SourceAssetAccountID:          strconv.FormatInt(int64(srt.SrcAccountID), 10),
			SourceAssetAccountName:        srt.SourceAccount.Name,
			DestinationExpenseAccountID:   strconv.FormatInt(int64(srt.DestAccountID), 10),
			DestinationExpenseAccountName: srt.DestinationAccount.Name,
			RecurrenceRule:                newRecurrenceRuleResponse(srt.RecurrenceRule),
		}

		if srt.TransactionCategory != nil {
			trcRes := newTransactionCategoryResponse(srt.TransactionCategory)
			e.TransactionCategory = &trcRes
		}

		expenses[idx] = e
	}

	trcs := make([]transactionCategoryResponse, len(transactionCategories))
	for idx, e := range transactionCategories {
		trcs[idx] = newTransactionCategoryResponse(e)
	}

	return plannerParametersResponse{
		Assets:                assets,
		Revenues:              revenues,
		Liabilities:           liabilities,
		Expenses:              expenses,
		TransactionCategories: trcs,
	}
}

type accountResponse struct {
	ID            uint32    `json:"id"`
	Category      string    `json:"category"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Currency      string    `json:"currency"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Amount        *float64  `json:"amount"`
	AmountDisplay *string   `json:"amount_display"`
}

func newAccountResponse(a *Account) accountResponse {
	amount := a.Amount.AsMajorUnits()
	amountDisplay := a.Amount.Display()
	return accountResponse{
		ID:            a.ID,
		Category:      string(a.Category),
		Name:          a.Name,
		Description:   a.Description,
		Currency:      a.Currency,
		Type:          string(a.GetType()),
		CreatedAt:     a.CreatedAt,
		UpdatedAt:     a.UpdatedAt,
		Amount:        &amount,
		AmountDisplay: &amountDisplay,
	}
}

func newListAccountsResponse(accounts []*Account) []accountResponse {
	result := make([]accountResponse, len(accounts))
	for idx, a := range accounts {
		result[idx] = newAccountResponse(a)
	}
	return result
}

type transactionCategoryResponse struct {
	ID          uint32                                  `json:"id"`
	Name        string                                  `json:"name"`
	Allocations []transactionCategoryAllocationResponse `json:"allocations"`
}

func newTransactionCategoryResponse(trc *TransactionCategory) transactionCategoryResponse {
	e := transactionCategoryResponse{
		ID:   trc.id,
		Name: trc.name,
	}

	e.Allocations = make([]transactionCategoryAllocationResponse, len(trc.allocations))
	for idx, alc := range trc.allocations {
		e.Allocations[idx] = *newTransactionCategoryAllocationResponse(alc)
	}

	return e
}

type transactionCategoryAllocationResponse struct {
	ID       uint32  `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func newTransactionCategoryAllocationResponse(e *TransactionCategoryAllocation) *transactionCategoryAllocationResponse {
	return &transactionCategoryAllocationResponse{
		ID:       e.ID,
		Amount:   e.Amount.AsMajorUnits(),
		Currency: e.Amount.Currency().Code,
	}
}

type transactionResponse struct {
	ID            uint32    `json:"id"`
	Description   string    `json:"description"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	AmountDisplay string    `json:"amount_display"`
	TransactedAt  time.Time `json:"transacted_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	SourceAccountID        uint32                       `json:"source_account_id"`
	SourceAccountName      string                       `json:"source_account_name"`
	DestinationAccountID   uint32                       `json:"destination_account_id"`
	DestinationAccountName string                       `json:"destination_account_name"`
	ScheduledTransaction   scheduledTransactionResponse `json:"scheduled_transaction"`
}

func newTransactionResponse(t *Transaction) transactionResponse {
	d := transactionResponse{
		ID:            t.ID,
		Description:   t.Description,
		Amount:        t.Amount.AsMajorUnits(),
		Currency:      t.Amount.Currency().Code,
		AmountDisplay: t.Amount.Display(),
		TransactedAt:  t.TransactedAt,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,

		SourceAccountID:        t.SourceAccount.ID,
		SourceAccountName:      t.SourceAccount.Name,
		DestinationAccountID:   t.DestinationAccount.ID,
		DestinationAccountName: t.DestinationAccount.Name,
	}

	if t.ScheduledTransaction != nil {
		d.ScheduledTransaction = newScheduledTransactionResponse(t.ScheduledTransaction)
	}

	return d
}

func newListTransactionsResponse(transactions []*Transaction) []transactionResponse {
	result := make([]transactionResponse, len(transactions))
	for idx, t := range transactions {
		result[idx] = newTransactionResponse(t)
	}
	return result
}

type scheduledTransactionResponse struct {
	ID            uint32    `json:"id"`
	Description   string    `json:"description"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	AmountDisplay string    `json:"amount_display"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	RecurrenceRule recurrenceRuleResponse `json:"recurrence_rule"`

	SourceAccountID        uint32 `json:"source_account_id"`
	SourceAccountName      string `json:"source_account_name"`
	DestinationAccountID   uint32 `json:"destination_account_id"`
	DestinationAccountName string `json:"destination_account_name"`
}

func newScheduledTransactionResponse(st *ScheduledTransaction) scheduledTransactionResponse {
	d := scheduledTransactionResponse{
		ID:            st.ID,
		Description:   st.Description,
		Amount:        st.Amount.AsMajorUnits(),
		Currency:      st.Amount.Currency().Code,
		AmountDisplay: st.Amount.Display(),
		CreatedAt:     st.CreatedAt,
		UpdatedAt:     st.UpdatedAt,

		RecurrenceRule: newRecurrenceRuleResponse(st.RecurrenceRule),

		SourceAccountID:        st.SourceAccount.ID,
		SourceAccountName:      st.SourceAccount.Name,
		DestinationAccountID:   st.DestinationAccount.ID,
		DestinationAccountName: st.DestinationAccount.Name,
	}

	return d
}

type recurrenceRuleResponse struct {
	Freq     string    `json:"freq"`
	Dtstart  time.Time `json:"dtstart"`
	Count    int       `json:"count"`
	Interval int       `json:"interval"`
	Until    time.Time `json:"until"`
}

func newRecurrenceRuleResponse(r *RecurrenceRule) recurrenceRuleResponse {
	d := recurrenceRuleResponse{
		Freq:     r.Rule.Options.Freq.String(),
		Dtstart:  r.Rule.Options.Dtstart,
		Count:    r.Rule.OrigOptions.Count,
		Interval: r.Rule.Options.Interval,
		Until:    r.Rule.Options.Until,
	}

	return d
}
