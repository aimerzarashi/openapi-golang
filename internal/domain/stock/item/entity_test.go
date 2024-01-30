package item_test

import (
	"testing"

	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

func TestEntity(t *testing.T) {
	// When
	uuid := uuid.New()
	id := item.Id(uuid)

	// Then
	if id.UUID() != uuid {
		t.Errorf("expected %s, got %s", uuid, id.UUID())
	}
}