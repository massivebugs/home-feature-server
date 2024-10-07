package response

import (
	"strconv"
	"time"

	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type Summary map[string]struct {
	Revenue string `json:"revenue"`
	Expense string `json:"expense"`
	Profit  string `json:"profit"`
}

type GetOverviewResponseDTO struct {
	From                      time.Time                `json:"from"`
	To                        time.Time                `json:"to"`
	NetWorth                  map[string]string        `json:"net_worth"`
	ProfitLossSummary         Summary                  `json:"profit_loss_summary"`
	AssetAccounts             []AccountResponseDTO     `json:"asset_accounts"`
	LiabilityAccounts         []AccountResponseDTO     `json:"liability_accounts"`
	Transactions              []TransactionResponseDTO `json:"transactions"`
	TransactionsFromScheduled []TransactionResponseDTO `json:"transactions_from_scheduled"`
}

func NewGetOverviewResponseDTO(from time.Time, to time.Time, ledger *cashbunny.Ledger, transactionsFromScheduled []*cashbunny.Transaction) *GetOverviewResponseDTO {
	netWorth := map[string]string{}
	for k, money := range ledger.GetNetWorth(&to) {
		netWorth[k] = money.Display()
	}

	revenues, expenses, sums := ledger.GetProfitLoss(&from, &to)
	result := Summary{}
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

	return &GetOverviewResponseDTO{
		From:                      from,
		To:                        to,
		NetWorth:                  netWorth,
		ProfitLossSummary:         result,
		AssetAccounts:             NewListAccountsResponseDTO(ledger.GetAccountsByCategory(cashbunny.AccountCategoryAssets)),
		LiabilityAccounts:         NewListAccountsResponseDTO(ledger.GetAccountsByCategory(cashbunny.AccountCategoryLiabilities)),
		Transactions:              NewListTransactionsResponseDTO(ledger.GetTransactions(&from, &to)),
		TransactionsFromScheduled: NewListTransactionsResponseDTO(transactionsFromScheduled),
	}
}

type GetPlanResponseDTO struct {
}

func NewGetPlanResponseDTO(planner *cashbunny.Planner) GetPlanResponseDTO {
	return GetPlanResponseDTO{}
}

type PlannerAssetDTO struct {
	AssetAccountID string  `json:"asset_account_id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
}

type PlannerRevenueDTO struct {
	ScheduledTransactionID      string                    `json:"scheduled_transaction_id"`
	Description                 string                    `json:"description"`
	Amount                      float64                   `json:"amount"`
	Currency                    string                    `json:"currency"`
	SourceRevenueAccountID      string                    `json:"source_revenue_account_id"`
	SourceRevenueAccountName    string                    `json:"source_revenue_account_name"`
	DestinationAssetAccountID   string                    `json:"destination_asset_account_id"`
	DestinationAssetAccountName string                    `json:"destination_asset_account_name"`
	RecurrenceRule              RecurrenceRuleResponseDTO `json:"recurrence_rule"`
}

type PlannerLiabilityDTO struct {
	ScheduledTransactionID          string                    `json:"scheduled_transaction_id"`
	Description                     string                    `json:"description"`
	Amount                          float64                   `json:"amount"`
	Currency                        string                    `json:"currency"`
	SourceAssetAccountID            string                    `json:"source_asset_account_id"`
	SourceAssetAccountName          string                    `json:"source_asset_account_name"`
	DestinationLiabilityAccountID   string                    `json:"destination_liability_account_id"`
	DestinationLiabilityAccountName string                    `json:"destination_liability_account_name"`
	RecurrenceRule                  RecurrenceRuleResponseDTO `json:"recurrence_rule"`
}

type GetPlannerParametersResponseDTO struct {
	Assets      []PlannerAssetDTO     `json:"assets"`
	Revenues    []PlannerRevenueDTO   `json:"revenues"`
	Liabilities []PlannerLiabilityDTO `json:"liabilities"`
}

func NewGetPlannerParametersResponseDTO(
	assetAccounts []*cashbunny.Account,
	scheduledRevenueTransactions []*cashbunny.ScheduledTransaction,
	scheduledLiabilityTransactions []*cashbunny.ScheduledTransaction,
) GetPlannerParametersResponseDTO {
	var assets []PlannerAssetDTO
	for _, aa := range assetAccounts {
		assets = append(assets, PlannerAssetDTO{
			AssetAccountID: strconv.FormatInt(int64(aa.ID), 10),
			Name:           aa.Name,
			Description:    aa.Description,
			Amount:         aa.Amount.AsMajorUnits(),
			Currency:       aa.Currency,
		})
	}

	var revenues []PlannerRevenueDTO
	for _, srt := range scheduledRevenueTransactions {
		revenues = append(revenues, PlannerRevenueDTO{
			ScheduledTransactionID:      strconv.FormatInt(int64(srt.ID), 10),
			Description:                 srt.Description,
			Amount:                      srt.Amount.AsMajorUnits(),
			Currency:                    srt.Amount.Currency().Code,
			SourceRevenueAccountID:      strconv.FormatInt(int64(srt.SrcAccountID), 10),
			SourceRevenueAccountName:    srt.SourceAccount.Name,
			DestinationAssetAccountID:   strconv.FormatInt(int64(srt.DestAccountID), 10),
			DestinationAssetAccountName: srt.DestinationAccount.Name,
			RecurrenceRule:              NewRecurrenceRuleResponseDTO(srt.RecurrenceRule),
		})
	}

	var liabilities []PlannerLiabilityDTO
	for _, srt := range scheduledLiabilityTransactions {
		liabilities = append(liabilities, PlannerLiabilityDTO{
			ScheduledTransactionID:          strconv.FormatInt(int64(srt.ID), 10),
			Description:                     srt.Description,
			Amount:                          srt.Amount.AsMajorUnits(),
			Currency:                        srt.Amount.Currency().Code,
			SourceAssetAccountID:            strconv.FormatInt(int64(srt.SrcAccountID), 10),
			SourceAssetAccountName:          srt.SourceAccount.Name,
			DestinationLiabilityAccountID:   strconv.FormatInt(int64(srt.DestAccountID), 10),
			DestinationLiabilityAccountName: srt.DestinationAccount.Name,
			RecurrenceRule:                  NewRecurrenceRuleResponseDTO(srt.RecurrenceRule),
		})
	}

	return GetPlannerParametersResponseDTO{
		Assets:      assets,
		Revenues:    revenues,
		Liabilities: liabilities,
	}
}

type GetAllCurrenciesDTO struct {
	CurrenciesAndGrapheme map[string]string `json:"currencies_and_grapheme"`
}

func NewGetAllCurrenciesResponseDTO(cgMap map[string]string) GetAllCurrenciesDTO {
	return GetAllCurrenciesDTO{
		CurrenciesAndGrapheme: cgMap,
	}
}

type GetUserPreferencesDTO struct {
	UserCurrencies []string `json:"user_currencies"`
}

func NewGetUserPreferencesDTO(up *cashbunny.UserPreferences) GetUserPreferencesDTO {
	ucs := make([]string, len(up.UserCurrencies))
	for idx, ccode := range up.UserCurrencies {
		ucs[idx] = ccode
	}

	return GetUserPreferencesDTO{
		UserCurrencies: ucs,
	}
}

type AccountResponseDTO struct {
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

func NewAccountResponseDTO(a *cashbunny.Account) AccountResponseDTO {
	amount := a.Amount.AsMajorUnits()
	amountDisplay := a.Amount.Display()
	return AccountResponseDTO{
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

func NewListAccountsResponseDTO(accounts []*cashbunny.Account) []AccountResponseDTO {
	result := make([]AccountResponseDTO, len(accounts))
	for idx, a := range accounts {
		result[idx] = NewAccountResponseDTO(a)
	}
	return result
}

type TransactionResponseDTO struct {
	ID            uint32    `json:"id"`
	Description   string    `json:"description"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	AmountDisplay string    `json:"amount_display"`
	TransactedAt  time.Time `json:"transacted_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	SourceAccountID        uint32                          `json:"source_account_id"`
	SourceAccountName      string                          `json:"source_account_name"`
	DestinationAccountID   uint32                          `json:"destination_account_id"`
	DestinationAccountName string                          `json:"destination_account_name"`
	ScheduledTransaction   ScheduledTransactionResponseDTO `json:"scheduled_transaction"`
}

func NewTransactionResponseDTO(t *cashbunny.Transaction) TransactionResponseDTO {
	d := TransactionResponseDTO{
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
		d.ScheduledTransaction = NewScheduledTransactionResponseDTO(t.ScheduledTransaction)
	}

	return d
}

func NewListTransactionsResponseDTO(transactions []*cashbunny.Transaction) []TransactionResponseDTO {
	result := make([]TransactionResponseDTO, len(transactions))
	for idx, t := range transactions {
		result[idx] = NewTransactionResponseDTO(t)
	}
	return result
}

type ScheduledTransactionResponseDTO struct {
	ID            uint32    `json:"id"`
	Description   string    `json:"description"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	AmountDisplay string    `json:"amount_display"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	RecurrenceRule RecurrenceRuleResponseDTO `json:"recurrence_rule"`

	SourceAccountID        uint32 `json:"source_account_id"`
	SourceAccountName      string `json:"source_account_name"`
	DestinationAccountID   uint32 `json:"destination_account_id"`
	DestinationAccountName string `json:"destination_account_name"`
}

func NewScheduledTransactionResponseDTO(st *cashbunny.ScheduledTransaction) ScheduledTransactionResponseDTO {
	d := ScheduledTransactionResponseDTO{
		ID:            st.ID,
		Description:   st.Description,
		Amount:        st.Amount.AsMajorUnits(),
		Currency:      st.Amount.Currency().Code,
		AmountDisplay: st.Amount.Display(),
		CreatedAt:     st.CreatedAt,
		UpdatedAt:     st.UpdatedAt,

		RecurrenceRule: NewRecurrenceRuleResponseDTO(st.RecurrenceRule),

		SourceAccountID:        st.SourceAccount.ID,
		SourceAccountName:      st.SourceAccount.Name,
		DestinationAccountID:   st.DestinationAccount.ID,
		DestinationAccountName: st.DestinationAccount.Name,
	}

	return d
}

func NewListScheduledTransactionResponseDTO(sts []*cashbunny.ScheduledTransaction) []ScheduledTransactionResponseDTO {
	result := make([]ScheduledTransactionResponseDTO, len(sts))
	for idx, st := range sts {
		result[idx] = NewScheduledTransactionResponseDTO(st)
	}
	return result
}

type RecurrenceRuleResponseDTO struct {
	Freq     string    `json:"freq"`
	Dtstart  time.Time `json:"dtstart"`
	Count    int       `json:"count"`
	Interval int       `json:"interval"`
	Until    time.Time `json:"until"`
}

func NewRecurrenceRuleResponseDTO(r *cashbunny.RecurrenceRule) RecurrenceRuleResponseDTO {
	d := RecurrenceRuleResponseDTO{
		Freq:     r.Rule.Options.Freq.String(),
		Dtstart:  r.Rule.Options.Dtstart,
		Count:    r.Rule.OrigOptions.Count,
		Interval: r.Rule.Options.Interval,
		Until:    r.Rule.Options.Until,
	}

	return d
}
