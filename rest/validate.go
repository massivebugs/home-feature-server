package rest

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type validationError struct {
	error
	Messages map[string]string
}

type requestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *requestValidator {
	v := validator.New(validator.WithRequiredStructEnabled())

	// Copied straight from go-playground/validator documentation
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	return &requestValidator{
		validator: v,
	}
}

func (rv *requestValidator) Validate(req interface{}) error {
	if err := rv.validator.Struct(req); err != nil {
		ves := err.(validator.ValidationErrors)
		messages := map[string]string{}
		for _, fe := range ves {
			messages[fe.Field()] = fe.Tag()
		}

		return &validationError{
			error:    err,
			Messages: messages,
		}
	}

	return nil
}
