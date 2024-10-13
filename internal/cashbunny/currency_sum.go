package cashbunny

import (
	"github.com/Rhymond/go-money"
)

type CurrencySums map[string]*money.Money

func NewCurrencySums(sums []*money.Money) CurrencySums {
	csMap := map[string]*money.Money{}

	for _, s := range sums {
		csMap[s.Currency().Code] = s
	}

	return CurrencySums(csMap)
}

func (m *CurrencySums) add(amount *money.Money) error {
	cc := amount.Currency().Code
	csMap := map[string]*money.Money(*m)

	sum, exists := csMap[cc]
	var err error
	if exists {
		sum, err = sum.Add(amount)
	} else {
		sum, err = money.New(0, cc).Add(amount)
	}
	if err != nil {
		return err
	}

	csMap[cc] = sum

	return nil
}

func (m *CurrencySums) subtract(amount *money.Money) error {
	cc := amount.Currency().Code
	csMap := map[string]*money.Money(*m)

	sum, exists := csMap[cc]
	var err error
	if exists {
		sum, err = sum.Subtract(amount)
	} else {
		sum, err = money.New(0, cc).Subtract(amount)
	}
	if err != nil {
		return err
	}

	csMap[cc] = sum

	return nil
}

func (m *CurrencySums) addSums(o CurrencySums) error {
	for _, v := range o {
		err := m.add(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *CurrencySums) subtractSums(o CurrencySums) error {
	for _, v := range o {
		err := m.subtract(v)
		if err != nil {
			return err
		}
	}
	return nil
}
