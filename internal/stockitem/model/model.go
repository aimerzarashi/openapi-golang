package domain

import "github.com/google/uuid"

type StockItemId uuid.UUID

type StockItem struct {
	Id   StockItemId
	Name string
}

func NewStockItem(id StockItemId, name string) *StockItem {
	return &StockItem{
		Id:   id,
		Name: name,
	}
}