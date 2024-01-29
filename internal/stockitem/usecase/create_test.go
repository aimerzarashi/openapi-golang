package usecase

import (
	"openapi/internal/infra/database"
	"openapi/internal/stockitem/domain"
	"openapi/internal/stockitem/repository"
	"testing"

	"github.com/google/uuid"
)


func TestCreateSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &repository.Repository{DB: db}
	
	// Given
	reqDto := &CreateRequestDto{
		Name: uuid.NewString(),
	}

	// When	
	resDto, err := Create(reqDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDto.Id == uuid.Nil {
		t.Errorf("expected %s, got %s", uuid.Nil, resDto.Id)
	}

	model, err := repository.Get(domain.StockItemId(resDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if model.Name != reqDto.Name {
		t.Errorf("expected %s, got %s", reqDto.Name, model.Name)
	}	
}