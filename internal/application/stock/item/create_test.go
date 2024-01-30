package item_test

import (
	"openapi/internal/application/stock/item"
	domain "openapi/internal/domain/stock/item"
	"openapi/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
)


func TestCreateSuccess(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}
	
	// Given
	reqDto := &item.CreateRequestDto{
		Name: uuid.NewString(),
	}

	// When	
	resDto, err := item.Create(reqDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDto.Id == uuid.Nil {
		t.Errorf("expected %s, got %s", uuid.Nil, resDto.Id)
	}

	a, err := repository.Get(domain.Id(resDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if a.GetName() != reqDto.Name {
		t.Errorf("expected %s, got %s", reqDto.Name, a.GetName())
	}
}