package usecase

import (
	"openapi/internal/infra/database"
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
	
	// Given
	reqDto := &CreateRequestDto{
		Name: uuid.NewString(),
	}

	// When	
	resDto, err := Create(reqDto, db)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDto.Id == uuid.Nil {
		t.Errorf("expected %s, got %s", uuid.Nil, resDto.Id)
	}

}