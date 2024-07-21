package cashbunny

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateCategoryRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateCategoryRequestDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Name,
			validation.Required,
		),
		validation.Field(
			&r.Description,
			validation.Required,
		),
	)
}

type CreateAccountRequestDTO struct {
	CategoryID  uint32  `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Balance     float64 `json:"balance"`
	Currency    string  `json:"currency"`
	Type        string  `json:"type"`
	OrderIndex  uint32  `json:"order_index"`
}

func (r *CreateAccountRequestDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.CategoryID,
			validation.Required,
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
			&r.Balance,
			validation.Required,
		),
		validation.Field(
			&r.Currency,
			validation.Required,
			validation.By(IsValidCurrency(r.Currency)),
		),
		validation.Field(
			&r.Type,
			validation.Required,
			validation.In(string(AccountTypeCredit), string(AccountTypeDebit)),
		),
		validation.Field(
			&r.OrderIndex,
			validation.Required,
		),
	)
}
