package response

import (
	"time"

	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type ListAccountResponseDTO struct {
	ID          uint32                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Balance     string                `json:"balance"`
	Currency    string                `json:"currency"`
	Type        cashbunny.AccountType `json:"type"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

func NewListAccountResponseDTO(accounts []cashbunny.Account) []ListAccountResponseDTO {
	result := make([]ListAccountResponseDTO, len(accounts))
	for idx, a := range accounts {
		result[idx] = ListAccountResponseDTO{
			ID:          a.ID,
			Name:        a.Name,
			Description: a.Description,
			Balance:     a.Balance.Display(),
			Currency:    a.Balance.Currency().Code,
			Type:        a.Type,
			CreatedAt:   a.CreatedAt,
			UpdatedAt:   a.UpdatedAt,
		}
	}
	return result
}
