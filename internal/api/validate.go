package api

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestValidator struct{}

func (*RequestValidator) Validate(i interface{}) error {
	if v, ok := i.(validation.Validatable); ok {
		return v.Validate()
	}
	return NewAPIError(CodeInternalServerError, errors.New("request cannot be validated"))
}
