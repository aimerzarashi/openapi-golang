package item

import "fmt"

type ItemName struct {
	value string
}

var ErrEmptyItemName = fmt.Errorf("empty item name")

func NewItemName(value string) (*ItemName, error) {
	if value == "" {
		return nil, ErrEmptyItemName
	}
	return &ItemName{value}, nil
}

func (v *ItemName) String() string {
	return v.value
}

