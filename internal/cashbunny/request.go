package cashbunny

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateAccountRequestDTO struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
	OrderIndex  uint32 `json:"order_index"`
}

func (r *CreateAccountRequestDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Category,
			validation.Required,
			validation.In(
				string(AccountCategoryAssets),
				string(AccountCategoryLiabilities),
				string(AccountCategoryRevenues),
				string(AccountCategoryExpenses),
			),
		),
		validation.Field(
			&r.Name,
			validation.Required,
			validation.Length(0, 255),
		),
		validation.Field(
			&r.Description,
			validation.Length(0, 255),
		),
		validation.Field(
			&r.Currency,
			validation.Required,
			validation.By(IsValidCurrency(r.Currency)),
		),
		validation.Field(
			&r.OrderIndex,
		),
	)
}

type CreateTransactionRequestDTO struct {
	Description          string  `json:"description"`
	Amount               float64 `json:"amount"`
	Currency             string  `json:"currency"`
	SourceAccountID      uint32  `json:"source_account_id"`
	DestinationAccountID uint32  `json:"destination_account_id"`
	TransactedAt         string  `json:"transacted_at"`
}

func (r *CreateTransactionRequestDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Description,
			validation.Required,
		),
		validation.Field(
			&r.Amount,
			validation.Required,
		),
		validation.Field(
			&r.Currency,
			validation.Required,
			validation.By(IsValidCurrency(r.Currency)),
		),
		validation.Field(
			&r.SourceAccountID,
			validation.Required,
		),
		validation.Field(
			&r.DestinationAccountID,
			validation.Required,
		),
		validation.Field(
			&r.TransactedAt,
			validation.Required,
			validation.Date(time.DateTime),
		),
	)
}
