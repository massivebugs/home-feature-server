package repeat

import validation "github.com/go-ozzo/ozzo-validation"

type RepeatRequest struct {
	Message string `json:"message"`
}

func (r *RepeatRequest) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Message,
			validation.Required,
			validation.Length(1, 20),
		),
	)
}
