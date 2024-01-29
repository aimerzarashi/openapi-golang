package model

import "github.com/google/uuid"

type StockItemId uuid.UUID

type StockItem struct {
	Id   		StockItemId
	Name 		string
	Deleted bool
}

func NewStockItem(id StockItemId, name string) *StockItem {
	return &StockItem{
		Id:   id,
		Name: name,
		Deleted: false,
	}
}

func (s *StockItem) Delete() {
	s.Deleted = true
}