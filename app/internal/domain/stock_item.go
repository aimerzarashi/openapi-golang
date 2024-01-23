package domain

import (
	"github.com/google/uuid"
)

type StockItem struct {
	Id        uuid.UUID
	Name      string
}

func NewStockItem(name string) *StockItem {
	return &StockItem{
		Id:        uuid.New(),
		Name:      name,
	}
}