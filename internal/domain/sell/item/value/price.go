package value

import "openapi/internal/domain/common/value/currency"

type Price struct {
	currency.Currency
}

func NewPrice(ammount float64, unit string) (Price, error) {
	currency, err := currency.NewCurrency(ammount, unit)
	if err != nil {
		return Price{}, err
	}
	return Price{
		Currency: currency,
	}, nil
}