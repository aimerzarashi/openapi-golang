package location_test

import (
	"testing"

	"github.com/google/uuid"

	"openapi/internal/domain/stock/location"
)

func TestEntity(t *testing.T) {
	// When
	uuid := uuid.New()
	id := location.Id(uuid)

	// Then
	if id.UUID() != uuid {
		t.Errorf("expected %s, got %s", uuid, id.UUID())
	}
}