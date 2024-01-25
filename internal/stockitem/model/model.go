package model

import "github.com/google/uuid"

type StockItemId uuid.UUID

type StockItem struct {
	Id   StockItemId
	Name string
}



func New(id StockItemId, name string) *StockItem {
	return &StockItem{
		Id:   id,
		Name: name,
	}
}