package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"

	"github.com/google/uuid"
)

func TestNewAggregate(t *testing.T) {
	// When
	name, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.New(name)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.GetId().UUID() == uuid.Nil {
		t.Errorf("expected %s, got %s", uuid.Nil, a.GetId().UUID())
	}
	if a.GetName() != name.String() {
		t.Errorf("expected %s, got %s", name, a.GetName())
	}
	if a.IsDeleted() != false {
		t.Errorf("expected %t, got %t", false, a.IsDeleted())
	}
}

func TestChangeName(t *testing.T) {
	// Given
	beforeName, err := item.NewItemName("test1")
	if err != nil {
		t.Fatal(err)
	}
	afterName, err := item.NewItemName("test2")
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.New(beforeName)
	if err != nil {
		t.Fatal(err)
	}

	// When
	a.ChangeName(afterName)

	// Then
	if a.GetName() != "test2" {
		t.Errorf("expected %s, got %s", "test2", a.GetName())
	}
}

func TestDelete(t *testing.T) {
	// When
	name, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.New(name)
	if err != nil {
		t.Fatal(err)
	}

	// When
	a.Delete()

	// Then
	if a.IsDeleted() != true {
		t.Errorf("expected %t, got %t", true, a.IsDeleted())
	}
}