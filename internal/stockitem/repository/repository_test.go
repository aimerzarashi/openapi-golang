package repository

import (
	"testing"

	"openapi/internal/stockitem/domain"

	"github.com/google/uuid"
)

func TestSaveSuccess(t *testing.T) {
	// Given
	id := domain.StockItemId(uuid.New())
	name := uuid.NewString()
	model := domain.NewStockItem(id, name)
	// When
	err := Save(model)

	// Then
	if err != nil {
		t.Fatal(err)
	}
}