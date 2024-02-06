package item_test

import (
	"testing"

	"openapi/internal/domain/stock/item"
)

func TestNewItemName(t *testing.T) {
	// Given
	name := "test"

	// When
	a, err := item.NewItemName(name)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.String() != name {
		t.Errorf("expected %s, got %s", name, a.String())
	}
}

func TestNewItemNameEmpty(t *testing.T) {
	// Given
	name := ""

	// When
	a, err := item.NewItemName(name)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Then
	if a != nil {
		t.Errorf("expected nil, got %v", a)
	}

	if err != item.ErrEmptyItemName {
		t.Errorf("expected %v, got %v", item.ErrEmptyItemName, err)
	}
}