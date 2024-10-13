package cashbunny

import (
	"strconv"
	"time"
)

type summary map[string]struct {
	Revenue string `json:"revenue"`
	Expense string `json:"expense"`
	Profit  string `json:"profit"`
}

type overviewResponse struct {
	From                      time.Time             `json:"from"`
	To                        time.Time             `json:"to"`
	NetWorth                  map[string]string     `json:"net_worth"`
	ProfitLossSummary         summary               `json:"profit_loss_summary"`
	AssetAccounts             []accountResponse     `json:"asset_accounts"`
	LiabilityAccounts         []accountResponse     `json:"liability_accounts"`
	Transactions              []transactionResponse `json:"transactions"`
	TransactionsFromScheduled []transactionResponse `json:"transactions_from_scheduled"`
}

func newOverviewResponse(from time.Time, to time.Time, ledger *Ledger, transactionsFromScheduled []*Transaction) overviewResponse {
	netWorth := map[string]string{}
	for k, money := range ledger.getNetWorth(&to) {
		netWorth[k] = money.Display()
	}

	revenues, expenses, sums := ledger.getProfitLoss(&from, &to)
	result := summary{}
	for k, money := range revenues {
		values := result[k]
		values.Revenue = money.Display()
		result[k] = values
	}

	for k, money := range expenses {
		values := result[k]
		values.Expense = money.Display()
		result[k] = values
	}

	for k, money := range sums {
		values := result[k]
		values.Profit = money.Display()
		result[k] = values
	}

	return overviewResponse{
		From:                      from,
		To:                        to,
		NetWorth:                  netWorth,
		ProfitLossSummary:         result,
		AssetAccounts:             newListAccountsResponse(ledger.getAccountsByCategory(AccountCategoryAssets)),
		LiabilityAccounts:         newListAccountsResponse(ledger.getAccountsByCategory(AccountCategoryLiabilities)),
		Transactions:              newListTransactionsResponse(ledger.getTransactions(&from, &to)),
		TransactionsFromScheduled: newListTransactionsResponse(transactionsFromScheduled),
	}
}

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
			AssetAccountID: strconv.FormatInt(int64(aa.id), 10),
			Name:           aa.name,
			Description:    aa.description,
			Amount:         aa.amount.AsMajorUnits(),
			Currency:       aa.currency,
		})
	}

	revenues := make([]plannerRevenueResponse, len(scheduledRevenueTransactions))
	for idx, srt := range scheduledRevenueTransactions {
		e := plannerRevenueResponse{
			ScheduledTransactionID:      strconv.FormatInt(int64(srt.id), 10),
			Description:                 srt.description,
			Amount:                      srt.amount.AsMajorUnits(),
			Currency:                    srt.amount.Currency().Code,
			SourceRevenueAccountID:      strconv.FormatInt(int64(srt.srcAccountID), 10),
			SourceRevenueAccountName:    srt.sourceAccount.name,
			DestinationAssetAccountID:   strconv.FormatInt(int64(srt.destAccountID), 10),
			DestinationAssetAccountName: srt.destinationAccount.name,
			RecurrenceRule:              newRecurrenceRuleResponse(srt.recurrenceRule),
		}

		if srt.transactionCategory != nil {
			trcRes := newTransactionCategoryResponse(srt.transactionCategory)
			e.TransactionCategory = &trcRes
		}

		revenues[idx] = e
	}

	liabilities := make([]plannerLiabilityResponse, len(scheduledLiabilityTransactions))
	for idx, srt := range scheduledLiabilityTransactions {
		e := plannerLiabilityResponse{
			ScheduledTransactionID:          strconv.FormatInt(int64(srt.id), 10),
			Description:                     srt.description,
			Amount:                          srt.amount.AsMajorUnits(),
			Currency:                        srt.amount.Currency().Code,
			SourceAssetAccountID:            strconv.FormatInt(int64(srt.srcAccountID), 10),
			SourceAssetAccountName:          srt.sourceAccount.name,
			DestinationLiabilityAccountID:   strconv.FormatInt(int64(srt.destAccountID), 10),
			DestinationLiabilityAccountName: srt.destinationAccount.name,
			RecurrenceRule:                  newRecurrenceRuleResponse(srt.recurrenceRule),
		}

		if srt.transactionCategory != nil {
			trcRes := newTransactionCategoryResponse(srt.transactionCategory)
			e.TransactionCategory = &trcRes
		}

		liabilities[idx] = e
	}

	expenses := make([]plannerExpenseResponse, len(scheduledExpenseTransactions))
	for idx, srt := range scheduledExpenseTransactions {
		e := plannerExpenseResponse{
			ScheduledTransactionID:        strconv.FormatInt(int64(srt.id), 10),
			Description:                   srt.description,
			Amount:                        srt.amount.AsMajorUnits(),
			Currency:                      srt.amount.Currency().Code,
			SourceAssetAccountID:          strconv.FormatInt(int64(srt.srcAccountID), 10),
			SourceAssetAccountName:        srt.sourceAccount.name,
			DestinationExpenseAccountID:   strconv.FormatInt(int64(srt.destAccountID), 10),
			DestinationExpenseAccountName: srt.destinationAccount.name,
			RecurrenceRule:                newRecurrenceRuleResponse(srt.recurrenceRule),
		}

		if srt.transactionCategory != nil {
			trcRes := newTransactionCategoryResponse(srt.transactionCategory)
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

type getAllCurrenciesResponse struct {
	CurrenciesAndGrapheme map[string]string `json:"currencies_and_grapheme"`
}

func newGetAllCurrenciesResponse(cgMap map[string]string) getAllCurrenciesResponse {
	return getAllCurrenciesResponse{
		CurrenciesAndGrapheme: cgMap,
	}
}

type userPreferencesResponse struct {
	UserCurrencies []string `json:"user_currencies"`
}

func newUserPreferencesResponse(up *UserPreferences) userPreferencesResponse {
	return userPreferencesResponse{
		UserCurrencies: up.userCurrencies,
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
	amount := a.amount.AsMajorUnits()
	amountDisplay := a.amount.Display()
	return accountResponse{
		ID:            a.id,
		Category:      string(a.category),
		Name:          a.name,
		Description:   a.description,
		Currency:      a.currency,
		Type:          string(a.getType()),
		CreatedAt:     a.createdAt,
		UpdatedAt:     a.updatedAt,
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
		ID:            t.id,
		Description:   t.description,
		Amount:        t.amount.AsMajorUnits(),
		Currency:      t.amount.Currency().Code,
		AmountDisplay: t.amount.Display(),
		TransactedAt:  t.transactedAt,
		CreatedAt:     t.createdAt,
		UpdatedAt:     t.updatedAt,

		SourceAccountID:        t.sourceAccount.id,
		SourceAccountName:      t.sourceAccount.name,
		DestinationAccountID:   t.destinationAccount.id,
		DestinationAccountName: t.destinationAccount.name,
	}

	if t.scheduledTransaction != nil {
		d.ScheduledTransaction = newScheduledTransactionResponse(t.scheduledTransaction)
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
		ID:            st.id,
		Description:   st.description,
		Amount:        st.amount.AsMajorUnits(),
		Currency:      st.amount.Currency().Code,
		AmountDisplay: st.amount.Display(),
		CreatedAt:     st.createdAt,
		UpdatedAt:     st.updatedAt,

		RecurrenceRule: newRecurrenceRuleResponse(st.recurrenceRule),

		SourceAccountID:        st.sourceAccount.id,
		SourceAccountName:      st.sourceAccount.name,
		DestinationAccountID:   st.destinationAccount.id,
		DestinationAccountName: st.destinationAccount.name,
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
		Freq:     r.rule.Options.Freq.String(),
		Dtstart:  r.rule.Options.Dtstart,
		Count:    r.rule.OrigOptions.Count,
		Interval: r.rule.Options.Interval,
		Until:    r.rule.Options.Until,
	}

	return d
}
