package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	MIN_USERNAME_LENGTH int = 3
	MAX_USERNAME_LENGTH int = 100
	MIN_PASSWORD_LENGTH int = 8
	MAX_PASSWORD_LENGTH int = 72 // bcrypt length limit
)

type CreateUserRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUserRequestDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Username,
			validation.Required,
			validation.Length(MIN_USERNAME_LENGTH, MAX_USERNAME_LENGTH).Error("username length is invalid"),
			is.Alphanumeric,
		),
		validation.Field(
			&r.Password,
			validation.Required,
			validation.By(IsValidPassword(MIN_PASSWORD_LENGTH, MAX_PASSWORD_LENGTH)),
		),
	)
}

type UserAuthRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *UserAuthRequestDTO) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(
			&r.Username,
			validation.Required,
		),
		validation.Field(
			&r.Password,
			validation.Required,
		),
	)
}
