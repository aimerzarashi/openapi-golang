package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	// When
	generatedUuid := uuid.New()
	id := StockItemId(generatedUuid)
	name := "test"
	stockItem := New(id, name)

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