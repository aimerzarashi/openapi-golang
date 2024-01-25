package model

import (
	"testing"

	"github.com/google/uuid"
)

func NewTestModel(t *testing.T) {
	// When
	uuid := uuid.New()
	id := uuid
	name := "test"
	stockItem := New(id, name)

	// Then
	if stockItem.Id != uuid {
		t.Errorf("expected %s, got %s", uuid, stockItem.Id)
	}
	if stockItem.Name != name {
		t.Errorf("expected %s, got %s", name, stockItem.Name)
	}
}