package usecase

import (
	"openapi/internal/infra/database"
	"openapi/internal/stockitem/domain"
	"openapi/internal/stockitem/repository"
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

	// Given
	beforeName := uuid.NewString()
	afterName := uuid.NewString()

	reqCreateDto := &CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := Create(reqCreateDto, db)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqUpdateDto := &UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	resUpdateDto, err := Update(reqUpdateDto, db)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resUpdateDto == nil {
		t.Errorf("expected not empty, actual empty")
	}

	model, err := repository.Get(db, domain.StockItemId(resCreateDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if model.Name != afterName {
		t.Errorf("expected %s, got %s", afterName, model.Name)
	}
}