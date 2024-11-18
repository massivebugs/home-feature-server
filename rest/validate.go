package rest

import (
	"errors"
	"reflect"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
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

func NewRequestValidator(cfg *Config) *requestValidator {
	v := validator.New(validator.WithRequiredStructEnabled())

	v.RegisterValidation("_password", auth.IsValidPassword)
	v.RegisterValidation("_iso8601", IsValidDateTime(cfg.APIDateTimeFormat))
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

// Custom validator for go-playground/validator
// Checks if the field is a valid string time can parse
func IsValidDateTime(formatStr string) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() != reflect.String {
			return false
		}

		value := fl.Field().String()

		return validation.Validate(value, validation.Date(formatStr)) == nil
	}
}

// Custom rule for ozzo-validation
// Checks if the value is a valid DateTime format which time can parse.
// This rule does not check more than precision in seconds.
func IsValidDateTimeFormat(v interface{}) error {
	formatStr, ok := v.(string)
	if !ok {
		return errors.New("datetime format cannot be converted to string")
	}

	testDate := time.Date(2024, 12, 25, 7, 7, 7, 0, time.UTC)

	parsedTestDate, err := time.Parse(formatStr, testDate.Format(formatStr))
	if err != nil {
		return errors.New("unsupported or invalid datetime format")
	}

	if testDate.UnixNano() != parsedTestDate.UnixNano() {
		return errors.New("unsupported or invalid datetime format")
	}

	return nil
}
