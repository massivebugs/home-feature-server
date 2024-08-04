package response

import (
	"time"

	"github.com/massivebugs/home-feature-server/db/service/cashbunny_account"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type CategoryResponseDTO struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCategoryResponseDTO(c *cashbunny_account.CashbunnyCategory) CategoryResponseDTO {
	return CategoryResponseDTO{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func NewListCategoriesResponseDTO(categories []*cashbunny_account.CashbunnyCategory) []CategoryResponseDTO {
	result := make([]CategoryResponseDTO, len(categories))
	for idx, c := range categories {
		result[idx] = NewCategoryResponseDTO(c)
	}
	return result
}

type ListAccountResponseDTO struct {
	ID          uint32                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Balance     string                `json:"balance"`
	Currency    string                `json:"currency"`
	Type        cashbunny.AccountType `json:"type"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	Category    CategoryResponseDTO   `json:"category"`
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
			Category: NewCategoryResponseDTO((*cashbunny_account.CashbunnyCategory)(a.Category)),
		}
	}
	return result
}
