package location_test

import (
	app "openapi/internal/app/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
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

	repo, err := infra.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	} 

	// Given
	name := "TestName"
	reqCreate, err := app.NewCreateRequest(name)
	if err != nil {
		t.Fatal(err)
	}

	newId := uuid.New()
	resCreate, err := app.Create(reqCreate, repo, newId)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDelete, err := app.NewDeleteRequest(resCreate.Id)
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Delete(reqDelete, repo); err != nil {
		t.Fatal(err)		
	}

	// Then
	id, err := domain.NewId(resCreate.Id)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.Get(id)
	if err == nil {
		t.Errorf("expected not nil, got nil")
	}
}

func TestDeleteFail(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo, err := infra.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqDelete, err := app.NewDeleteRequest(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Delete(reqDelete, repo); err == nil {
		t.Errorf("expected not nil, got nil")
	}

	// Then
	found, err := repo.Find(reqDelete.Id)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if found {
		t.Errorf("%T = %v, want %v", found, found, false)
	}

	_, err = repo.Get(reqDelete.Id)
	if err == nil {
		t.Errorf("expected not nil, got nil")
	}
}