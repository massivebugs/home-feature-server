package cashbunny

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/massivebugs/home-feature-server/internal/util"
)

// func isMoneyNotNegative(m *money.Money) func(v interface{}) error {
// 	return func(v interface{}) error {
// 		if m.IsNegative() {
// 			return errors.New("accounts/transactions cannot have negative balance/amount")
// 		}
// 		return nil
// 	}
// }

// Custom validator for go-playground/validator
// Checks if the field is a valid string of a currency which is supported by Cashbunny
func IsValidCurrency(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}

	value := fl.Field().String()

	return util.SliceExists(supportedCurrencyCodes, value)
}
