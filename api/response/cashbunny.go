package response

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/service/cashbunny_account"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type AccountCategoryResponseDTO struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewAccountCategoryResponseDTO(c *cashbunny_account.CashbunnyAccountCategory) AccountCategoryResponseDTO {
	return AccountCategoryResponseDTO{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func NewListAccountCategoriesResponseDTO(categories []*cashbunny_account.CashbunnyAccountCategory) []AccountCategoryResponseDTO {
	result := make([]AccountCategoryResponseDTO, len(categories))
	for idx, c := range categories {
		result[idx] = NewAccountCategoryResponseDTO(c)
	}
	return result
}

type ListAccountResponseDTO struct {
	ID          uint32                     `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Balance     string                     `json:"balance"`
	Currency    string                     `json:"currency"`
	Type        cashbunny.AccountType      `json:"type"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
	Category    AccountCategoryResponseDTO `json:"category"`
}

func NewListAccountResponseDTO(accounts []*cashbunny.Account) []ListAccountResponseDTO {
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
			// TODO: Let's group all the queries together for each module... gosh
			Category: NewAccountCategoryResponseDTO((*cashbunny_account.CashbunnyAccountCategory)(a.Category)),
		}
	}
	return result
}
