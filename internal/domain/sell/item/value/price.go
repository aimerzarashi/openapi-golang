package value

import "openapi/internal/domain/common/value"

type Price struct {
	value.Currency
}

func NewPrice(ammount float64, unit string) (Price, error) {
	currency, err := value.NewCurrency(ammount, unit)
	if err != nil {
		return Price{}, err
	}
	return Price{
		Currency: currency,
	}, nil
}