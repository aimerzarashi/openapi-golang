package currency

import (
	"errors"
)

type (
	Currency struct {
		ammount float64
		unit    string
	}
)

var (
	ErrCurrencyUnitInvalid = errors.New("Currency: unit cannot be empty")
)

func NewCurrency(ammount float64, unit string) (Currency, error) {
	if err := valideteUnit(unit); err != nil {
		return Currency{}, err
	}
	return Currency{
		ammount: ammount,
		unit:    unit,
	}, nil
}

func valideteUnit(unit string) error {
	code := []string{"JPY"}

	for _, v := range code {
		if unit == v {
			return nil
		}
	}

	return ErrCurrencyUnitInvalid
}

func (v Currency) Ammount() float64 {
	return v.ammount
}

func (v Currency) Unit() string {
	return v.unit
}
