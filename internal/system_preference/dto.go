package system_preference

import validation "github.com/go-ozzo/ozzo-validation"

type UserSystemPreferenceDTO struct {
	Language *string `json:"language"`
}

func (r *UserSystemPreferenceDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Language,
			validation.In(
				"en",
				"ja",
			),
		),
	)
}
