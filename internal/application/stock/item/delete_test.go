package item_test

import (
	"openapi/internal/application/stock/item"
	domain "openapi/internal/domain/stock/item"
	"openapi/internal/infrastructure/database"
	"testing"

	"github.com/google/uuid"
)


func TestDeleteSuccess(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &domain.Repository{Db: db}

	// Given
	reqCreateDto := &item.CreateRequestDto{
		Name: uuid.NewString(),
	}

	resCreateDto, err := item.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDeleteDto := &item.DeleteRequestDto{
		Id: resCreateDto.Id,
	}

	resDeleteDto, err := item.Delete(reqDeleteDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDeleteDto == nil {
		t.Errorf("expected not empty, actual empty")
	}
}