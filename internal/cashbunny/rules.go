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
