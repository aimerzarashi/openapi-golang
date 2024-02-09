package location_test

import (
	app "openapi/internal/app/stock/location"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
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

	resCreate, err := app.Create(reqCreate, repo, uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	// When
	newName := "NewTestName"
	reqUpdate, err := app.NewUpdateRequest(resCreate.Id, newName)
	if err != nil {
		t.Fatal(err)
	}

	err = app.Update(reqUpdate, repo)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	a, err := repo.Get(reqUpdate.Id)
	if err != nil {
		t.Fatal(err)
	}

	if a.Name.String() != newName {
		t.Errorf("%T = %v, want %v", a.Name.String(), a.Name.String(), newName)
	}
}