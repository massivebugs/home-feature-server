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

func (m *CurrencySums) Add(amount *money.Money) error {
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

func (m *CurrencySums) Subtract(amount *money.Money) error {
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

func (m *CurrencySums) AddSums(o CurrencySums) error {
	for _, v := range o {
		err := m.Add(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *CurrencySums) SubtractSums(o CurrencySums) error {
	for _, v := range o {
		err := m.Subtract(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func SliceFilter[T any](ss []T, test func(T) bool) []T {
	var ret []T
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}

	return ret
}

func SliceFind[T any](ss []T, test func(T) bool) *T {
	for _, s := range ss {
		if test(s) {
			return &s
		}
	}

	return nil
}
