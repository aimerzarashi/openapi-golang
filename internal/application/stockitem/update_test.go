package stockitem_test

import (
	"openapi/internal/application/stockitem"
	"openapi/internal/domain/model"
	"openapi/internal/domain/repository"
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &repository.StockItem{DB: db}

	// Given
	beforeName := uuid.NewString()
	afterName := uuid.NewString()

	reqCreateDto := &stockitem.CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := stockitem.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqUpdateDto := &stockitem.UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	resUpdateDto, err := stockitem.Update(reqUpdateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resUpdateDto == nil {
		t.Errorf("expected not empty, actual empty")
	}

	model, err := repository.Get(model.StockItemId(resCreateDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if model.Name != afterName {
		t.Errorf("expected %s, got %s", afterName, model.Name)
	}
}