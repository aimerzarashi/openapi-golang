package item_test

import (
	"fmt"
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

	if err := item.Delete(reqDeleteDto, repository); err != nil {
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

func TestDeleteError(t *testing.T) {	
	repository := &MockRepository{}
	
	// Given
	reqDto := &item.DeleteRequestDto{
		Id: uuid.New(),
	}

	// When
	err := item.Delete(reqDto, repository)
	if err == fmt.Errorf("not implemented") {
		t.Errorf("expected error, got nil")
	}
}