package item

import (
	"github.com/google/uuid"
)

type aggregate struct {
	id   		Id
	name 		name
	deleted bool
}

func New(name *name) (*aggregate, error) {
	return &aggregate{
		id:   Id(uuid.New()),
		name: *name,
		deleted: false,
	}, nil
}

func (a aggregate) GetId() Id {
	return a.id
}

func (a aggregate) GetName() string {
	return a.name.value
}

func (a aggregate) IsDeleted() bool {
	return a.deleted
}

func (a *aggregate) ChangeName(name *name) {
	a.name = *name
}

func (a *aggregate) Delete() {
	a.deleted = true
}