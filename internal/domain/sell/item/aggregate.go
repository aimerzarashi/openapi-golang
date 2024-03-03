package item

import (
	"openapi/internal/domain/sell/item/value"

	"github.com/aimerzarashi/timeslice"
)

type (
	Aggregate struct {
		Id      Id
		Name    value.Name
		Prices  *timeslice.Collection[value.Price]
		deleted bool
	}
)

func NewAggregate(id Id, name value.Name) (*Aggregate, error) {
	prices, err := timeslice.NewCollection[value.Price]()
	if err != nil {
		return nil, err		
	}

	return &Aggregate{
		Id:      id,
		Name:    name,
		Prices:  prices,
		deleted: false,
	}, nil
}

func RestoreAggregate(id Id, name value.Name, deleted bool) *Aggregate {
	return &Aggregate{
		Id:      id,
		Name:    name,
		deleted: deleted,
	}
}

func (a Aggregate) IsDeleted() bool {
	return a.deleted
}

func (a *Aggregate) Delete() {
	a.deleted = true
}
