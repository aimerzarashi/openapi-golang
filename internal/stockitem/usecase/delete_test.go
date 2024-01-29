package usecase

import (
	"openapi/internal/infra/database"
	"openapi/internal/stockitem/repository"
	"testing"

	"github.com/google/uuid"
)


func TestDeleteSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &repository.Repository{DB: db}

	// Given
	reqCreateDto := &CreateRequestDto{
		Name: uuid.NewString(),
	}

	resCreateDto, err := Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDeleteDto := &DeleteRequestDto{
		Id: resCreateDto.Id,
	}

	resDeleteDto, err := Delete(reqDeleteDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDeleteDto == nil {
		t.Errorf("expected not empty, actual empty")
	}
}