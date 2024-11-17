package cashbunny

import (
	"errors"
	"reflect"

	"github.com/Rhymond/go-money"
	"github.com/go-playground/validator/v10"
	"github.com/massivebugs/home-feature-server/internal/util"
)

func isMoneyNotNegative(m *money.Money) func(v interface{}) error {
	return func(v interface{}) error {
		if m.IsNegative() {
			return errors.New("accounts/transactions cannot have negative balance/amount")
		}
		return nil
	}
}

func isValidCurrency(currencyCode string) func(v interface{}) error {
	return func(v interface{}) error {
		c := money.GetCurrency(currencyCode)
		if c == nil {
			return errors.New("unsupported or invalid currency")
		}

		return nil
	}
}

func IsValidCurrency(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}

	value := fl.Field().String()

	return util.SliceExists(supportedCurrencyCodes, value)
}
