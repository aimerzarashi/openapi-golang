package location_test

import (
	app "openapi/internal/app/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
	"testing"

	"github.com/google/uuid"
)


func TestCreateSuccess(t *testing.T) {
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
	name := "test"
	req, err := app.NewCreateRequest(name)
	if err != nil {
		t.Fatal(err)
	}

	// When
	newId := uuid.New()
	res, err := app.Create(req, repo, newId)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if res.Id == uuid.Nil {
		t.Errorf("expected not nil, got nil")
	}

	if res.Id != newId {
		t.Errorf("%T = %v, want %v", res.Id, res.Id, newId)
	}

	id, err := domain.NewId(res.Id) 
	a, err := repo.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if a.Name.String() != name {
		t.Errorf("%T = %v, want %v", a.Name.String(), a.Name.String(), name)
	}
}