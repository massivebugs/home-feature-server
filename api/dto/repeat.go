package dto

import validation "github.com/go-ozzo/ozzo-validation"

type RepeatRequestDTO struct {
	Message string `json:"message"`
}

func (r *RepeatRequestDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Message,
			validation.Required,
			validation.Length(1, 20),
		),
	)
}
