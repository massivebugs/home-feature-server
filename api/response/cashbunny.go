package response

import (
	"time"

	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type Summaries map[string]struct {
	Revenue string `json:"revenue"`
	Expense string `json:"expense"`
	Profit  string `json:"profit"`
}

type GetOverviewResponseDTO struct {
	From                      time.Time                `json:"from"`
	To                        time.Time                `json:"to"`
	NetWorth                  map[string]string        `json:"net_worth"`
	Summaries                 Summaries                `json:"summaries"`
	Transactions              []TransactionResponseDTO `json:"transactions"`
	TransactionsFromScheduled []TransactionResponseDTO `json:"transactions_from_scheduled"`
}

func NewGetOverviewResponseDTO(from time.Time, to time.Time, ledger *cashbunny.Ledger, transactionsFromScheduled []*cashbunny.Transaction) *GetOverviewResponseDTO {
	netWorth := map[string]string{}
	for k, money := range ledger.GetNetWorth(&to) {
		netWorth[k] = money.Display()
	}

	revenues, expenses, sums := ledger.GetProfitLoss(&from, &to)
	result := Summaries{}
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
		Summaries:                 result,
		Transactions:              NewListTransactionsResponseDTO(ledger.GetTransactions(&from, &to)),
		TransactionsFromScheduled: NewListTransactionsResponseDTO(transactionsFromScheduled),
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
	for idx, uc := range up.UserCurrencies {
		ucs[idx] = uc.CurrencyCode
	}

	return GetUserPreferencesDTO{
		UserCurrencies: ucs,
	}
}

type AccountResponseDTO struct {
	ID             uint32    `json:"id"`
	Category       string    `json:"category"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Balance        float64   `json:"balance"`
	Currency       string    `json:"currency"`
	BalanceDisplay string    `json:"balance_display"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewAccountResponseDTO(a *cashbunny.Account) AccountResponseDTO {
	return AccountResponseDTO{
		ID:             a.ID,
		Category:       string(a.Category),
		Name:           a.Name,
		Description:    a.Description,
		Balance:        a.Balance.AsMajorUnits(),
		Currency:       a.Balance.Currency().Code,
		BalanceDisplay: a.Balance.Display(),
		Type:           string(a.GetType()),
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
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
