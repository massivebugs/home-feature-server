package response

import (
	"time"

	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type AccountResponseDTO struct {
	ID          uint32    `json:"id"`
	Category    string    `json:"category"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Balance     string    `json:"balance"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewAccountResponseDTO(a *cashbunny.Account) AccountResponseDTO {
	return AccountResponseDTO{
		ID:          a.ID,
		Category:    string(a.Category),
		Name:        a.Name,
		Description: a.Description,
		Balance:     a.Balance.Display(),
		Currency:    a.Balance.Currency().Code,
		Type:        string(a.Type),
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
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
	ID           uint32    `json:"id"`
	Description  string    `json:"description"`
	Amount       string    `json:"amount"`
	Currency     string    `json:"currency"`
	TransactedAt time.Time `json:"transacted_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	SourceAccountID        uint32 `json:"source_account_id"`
	SourceAccountName      string `json:"source_account_name"`
	DestinationAccountID   uint32 `json:"destination_account_id"`
	DestinationAccountName string `json:"destination_account_name"`
}

func NewTransactionResponseDTO(t *cashbunny.Transaction) TransactionResponseDTO {
	return TransactionResponseDTO{
		ID:           t.ID,
		Description:  t.Description,
		Amount:       t.Amount.Display(),
		Currency:     t.Amount.Currency().Code,
		TransactedAt: t.TransactedAt,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,

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
