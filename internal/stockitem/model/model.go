package model

import "github.com/google/uuid"

type StockItem struct {
	Id   uuid.UUID
	Name string
}



func New(id uuid.UUID, name string) *StockItem {
	return &StockItem{
		Id:   id,
		Name: name,
	}
}