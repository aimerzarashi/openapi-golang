package location_test

import (
	"openapi/internal/application/stock/location"
	domain "openapi/internal/domain/stock/location"
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
	reqCreateDto := &location.CreateRequestDto{
		Name: uuid.NewString(),
	}

	resCreateDto, err := location.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDeleteDto := &location.DeleteRequestDto{
		Id: resCreateDto.Id,
	}

	if err := location.Delete(reqDeleteDto, repository); err != nil {
		t.Fatal(err)		
	}

	// Then
	a, err := repository.Get(domain.Id(resCreateDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if !a.IsDeleted() {
		t.Errorf("expected %t, got %t", true, a.IsDeleted())
	}
}