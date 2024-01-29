package repository

import (
	"context"
	"testing"
	"time"

	"openapi/internal/infra/database"
	"openapi/internal/infra/sqlboiler"
	"openapi/internal/stockitem/domain"

	"github.com/google/uuid"
)

func TestCreateSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &Repository{DB: db}

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)
	name := uuid.NewString()
	model := domain.NewStockItem(id, name)
	currentDateTime := time.Now()

	// When
	err = r.Save(model)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	data, err := sqlboiler.FindStockItem(context.Background(), db, generatedUuid.String())
	if err != nil {
		t.Fatal(err)
	}

	if data.ID != generatedUuid.String() {
		t.Errorf("expected %s, got %s", generatedUuid.String(), data.ID)
	}
	
	if data.Name != name {
		t.Errorf("expected %s, got %s", name, data.Name)
	}

	if data.CreatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, data.CreatedAt)		
	}

	if data.UpdatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, data.UpdatedAt)		
	}
}

func TestUpdateSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &Repository{DB: db}

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)
	beforeName := uuid.NewString()
	afterName := uuid.NewString()
	model := domain.NewStockItem(id, beforeName)
	currentDateTime := time.Now()

	err = r.Save(model)
	if err != nil {
		t.Fatal(err)
	}

	beforeModel, err := r.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	beforeData, err := sqlboiler.FindStockItem(context.Background(), db, generatedUuid.String())
	if err != nil {
		t.Fatal(err)
	}

	// When
	beforeModel.Name = afterName
	err = r.Save( beforeModel)
	if err != nil {
		t.Fatal(err)
	}

	afterModel, err := r.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	data, err := sqlboiler.FindStockItem(context.Background(), db, generatedUuid.String())
	if err != nil {
		t.Fatal(err)
	}

	if data.ID != generatedUuid.String() {
		t.Errorf("expected %s, got %s", generatedUuid.String(), data.ID)
	}

	if data.Name != afterModel.Name {
		t.Errorf("expected %s, got %s", afterModel.Name, data.Name)
	}

	if data.CreatedAt.Equal(beforeData.CreatedAt) != true {
		t.Errorf("expected %s, got %s", beforeData.CreatedAt, data.CreatedAt)		
	}

	if data.UpdatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, data.UpdatedAt)		
	}
}

func TestFindSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &Repository{DB: db}

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)
	name := uuid.NewString()
	model := domain.NewStockItem(id, name)

	err = r.Save( model)
	if err != nil {
		t.Fatal(err)
	}

	// When
	found, err := r.Find( id)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if found != true {
		t.Errorf("expected %t, got %t", true, found)
	}

}


func TestFindFailure(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &Repository{DB: db}

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)

	// When
	found, err := r.Find(id)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if found != false {
		t.Errorf("expected %t, got %t", false, found)
	}

}