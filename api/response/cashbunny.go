package response

import (
	"time"

	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type GetOverviewResponseDTO struct {
	Revenues     map[string]string        `json:"revenues"`
	Expenses     map[string]string        `json:"expenses"`
	Sums         map[string]string        `json:"sums"`
	Totals       map[string]string        `json:"totals"`
	Transactions []*cashbunny.Transaction `json:"transactions"`
}

func NewGetOverviewResponseDTO(ledger *cashbunny.Ledger) *GetOverviewResponseDTO {
	revenues, expenses, sums := ledger.GetProfitLoss()

	return &GetOverviewResponseDTO{
		Revenues:     revenues.ToDefaultDisplayMap(),
		Expenses:     expenses.ToDefaultDisplayMap(),
		Sums:         sums.ToDefaultDisplayMap(),
		Transactions: ledger.GetTransactions(),
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
		Type:           string(a.Type),
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

	SourceAccountID        uint32 `json:"source_account_id"`
	SourceAccountName      string `json:"source_account_name"`
	DestinationAccountID   uint32 `json:"destination_account_id"`
	DestinationAccountName string `json:"destination_account_name"`
}

func NewTransactionResponseDTO(t *cashbunny.Transaction) TransactionResponseDTO {
	return TransactionResponseDTO{
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
}

func NewListTransactionsResponseDTO(transactions []*cashbunny.Transaction) []TransactionResponseDTO {
	result := make([]TransactionResponseDTO, len(transactions))
	for idx, t := range transactions {
		result[idx] = NewTransactionResponseDTO(t)
	}
	return result
}
