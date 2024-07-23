package cashbunny

import (
	"errors"

	"github.com/Rhymond/go-money"
)

func IsMoneyNotNegative(m *money.Money) func(v interface{}) error {
	return func(v interface{}) error {
		if m.IsNegative() {
			return errors.New("money amount must not be negative")
		}
		return nil
	}
}

func IsValidCurrency(currencyCode string) func(v interface{}) error {
	return func(v interface{}) error {
		c := money.GetCurrency(currencyCode)
		if c == nil {
			return errors.New("unsupported or invalid currency")
		}

		return nil
	}
}