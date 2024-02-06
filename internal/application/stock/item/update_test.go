package item_test

import (
	"openapi/internal/application/stock/item"
	domain "openapi/internal/domain/stock/item"
	"openapi/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateSuccess(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}

	// Given
	beforeName := uuid.NewString()
	afterName := uuid.NewString()

	reqCreateDto := &item.CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := item.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqUpdateDto := &item.UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	err = item.Update(reqUpdateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	a, err := repository.Get(domain.Id(resCreateDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if a.GetName() != afterName {
		t.Errorf("expected %s, got %s", afterName, a.GetName())
	}
}