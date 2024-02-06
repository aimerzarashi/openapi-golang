package item

import "fmt"

type name struct {
	value string
}

var ErrEmptyItemName = fmt.Errorf("empty item name")

func NewItemName(value string) (*name, error) {
	if value == "" {
		return nil, ErrEmptyItemName
	}
	return &name{value}, nil
}

func (v *name) String() string {
	return v.value
}

