package item

import (
	"github.com/google/uuid"
)

type aggregate struct {
	id   		Id
	name 		string
	deleted bool
}

func New(name string) (*aggregate, error) {
	return &aggregate{
		id:   Id(uuid.New()),
		name: name,
		deleted: false,
	}, nil
}

func (a aggregate) GetId() Id {
	return a.id
}

func (a aggregate) GetName() string {
	return string(a.name)
}

func (a aggregate) IsDeleted() bool {
	return a.deleted
}

func (a *aggregate) ChangeName(name string) {
	a.name = name
}

func (a *aggregate) Delete() {
	a.deleted = true
}