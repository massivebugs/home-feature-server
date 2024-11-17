package rest

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/auth"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
	"github.com/massivebugs/home-feature-server/rest/oapi"
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

	v.RegisterValidation("_password", auth.IsValidPassword)
	v.RegisterValidation("_cashbunny_currency", cashbunny.IsValidCurrency)

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
			messages[fe.Field()] = fe.ActualTag()
		}

		return &validationError{
			error:    err,
			Messages: messages,
		}
	}

	return nil
}

func RequestValidatorStrictHandlerFunc(f oapi.StrictHandlerFunc, operationID string) oapi.StrictHandlerFunc {
	return func(c echo.Context, req interface{}) (interface{}, error) {
		if err := c.Validate(req); err != nil {
			return nil, err
		}

		return f(c, req)
	}
}
