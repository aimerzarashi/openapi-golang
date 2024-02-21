package item

import (
	"openapi/internal/domain/sell/item/value"
)

type (
	Aggregate struct {
		Id      Id
		Name    value.Name
		deleted bool
	}
)

func NewAggregate(id Id, name value.Name) *Aggregate {
	return &Aggregate{
		Id:      id,
		Name:    name,
		deleted: false,
	}
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
