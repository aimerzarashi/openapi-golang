package location_test

import (
	"openapi/internal/application/stock/location"
	domain "openapi/internal/domain/stock/location"
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

	reqCreateDto := &location.CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := location.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqUpdateDto := &location.UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	err = location.Update(reqUpdateDto, repository)
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