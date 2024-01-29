package stockitem_test

import (
	"openapi/internal/application/stockitem"
	"openapi/internal/domain/repository"
	"openapi/internal/infra/database"
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
	repository := &repository.StockItem{DB: db}

	// Given
	reqCreateDto := &stockitem.CreateRequestDto{
		Name: uuid.NewString(),
	}

	resCreateDto, err := stockitem.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDeleteDto := &stockitem.DeleteRequestDto{
		Id: resCreateDto.Id,
	}

	resDeleteDto, err := stockitem.Delete(reqDeleteDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDeleteDto == nil {
		t.Errorf("expected not empty, actual empty")
	}
}