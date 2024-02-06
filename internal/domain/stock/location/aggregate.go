package location

import (
	"fmt"

	"github.com/google/uuid"
)

type Aggregate struct {
	id   		Id
	name 		string
	deleted bool
}

func New(name string) (*Aggregate, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %s", name)
	}

	return &Aggregate{
		id:   Id(uuid.New()),
		name: name,
		deleted: false,
	}, nil
}

func (a Aggregate) GetId() Id {
	return a.id
}

func (a Aggregate) GetName() string {
	return a.name
}

func (a Aggregate) IsDeleted() bool {
	return a.deleted
}

func (a *Aggregate) ChangeName(name string) {
	a.name = name
}

func (a *Aggregate) Delete() {
	a.deleted = true
}