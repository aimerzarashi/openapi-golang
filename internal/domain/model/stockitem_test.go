package model_test

import (
	"openapi/internal/domain/model"
	"testing"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	// When
	generatedUuid := uuid.New()
	id := model.StockItemId(generatedUuid)
	name := "test"
	stockItem := model.NewStockItem(id, name)

	// Then
	if stockItem.Id != id {
		t.Errorf("expected %s, got %s", id, stockItem.Id)
	}
	castedId := uuid.UUID(stockItem.Id)
	if castedId != generatedUuid {
		t.Errorf("expected %s, got %s", generatedUuid, castedId)
	}
	if stockItem.Name != name {
		t.Errorf("expected %s, got %s", name, stockItem.Name)
	}
}


func TestDelete(t *testing.T) {
	// Given
	generatedUuid := uuid.New()
	id := model.StockItemId(generatedUuid)
	name := "test"
	stockItem := model.NewStockItem(id, name)

	// When
	stockItem.Delete()

	// Then
	if stockItem.Deleted != true {
		t.Errorf("expected %t, got %t", true, stockItem.Deleted)
	}
}