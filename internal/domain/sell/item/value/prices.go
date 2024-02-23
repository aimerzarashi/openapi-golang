package value

import (
	"errors"
)

type (
	Price struct {
		ammount  float64
		currency string
	}
)

var (
	ErrPriceAmmount  = errors.New("Price: ammount must be greater than 0")
	ErrPriceCurrency = errors.New("Price: currency cannot be empty")
)

func NewPrice(ammount float64, currency string) (Price, error) {
	if ammount <= 0 {
		return Price{}, ErrPriceAmmount
	}
	if err := valideteCurrency(currency); err != nil {
		return Price{}, err
	}
	return Price{
		ammount:  ammount,
		currency: currency,
	}, nil
}

func valideteCurrency(currency string) error {
	code := []string{"JPY"}

	for _, v := range code {
		if currency == v {
			return nil
		}
	}

	return ErrPriceCurrency
}
