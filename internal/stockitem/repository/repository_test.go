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

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)
	name := uuid.NewString()
	model := domain.NewStockItem(id, name)
	currentDateTime := time.Now()
	// When
	err = Save(db, model)
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

	if data.CreatedAt.UTC().Before(currentDateTime.UTC()) == true {
		t.Errorf("expected %s, got %s", currentDateTime.UTC(), data.CreatedAt.UTC())		
	}

	if data.UpdatedAt.UTC().Before(currentDateTime.UTC()) == true {
		t.Errorf("expected %s, got %s", currentDateTime.UTC(), data.UpdatedAt.UTC())		
	}
}

func TestUpdateSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)
	beforeName := uuid.NewString()
	afterName := uuid.NewString()
	model := domain.NewStockItem(id, beforeName)
	currentDateTime := time.Now()

	err = Save(db, model)
	if err != nil {
		t.Fatal(err)
	}

	beforeModel, err := Get(db, id)
	if err != nil {
		t.Fatal(err)
	}
	beforeData, err := sqlboiler.FindStockItem(context.Background(), db, generatedUuid.String())
	if err != nil {
		t.Fatal(err)
	}

	// When
	beforeModel.Name = afterName
	err = Save(db, beforeModel)
	if err != nil {
		t.Fatal(err)
	}

	afterModel, err := Get(db, id)
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

	if data.CreatedAt.UTC().Equal(beforeData.CreatedAt.UTC()) != true {
		t.Errorf("expected %s, got %s", beforeData.CreatedAt.UTC(), data.CreatedAt.UTC())		
	}

	if data.UpdatedAt.UTC().Before(currentDateTime.UTC()) == true {
		t.Errorf("expected %s, got %s", currentDateTime.UTC(), data.UpdatedAt.UTC())		
	}
}

func TestFindSuccess(t *testing.T) {
	// Setup
	db, err := database.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)
	name := uuid.NewString()
	model := domain.NewStockItem(id, name)

	err = Save(db, model)
	if err != nil {
		t.Fatal(err)
	}

	// When
	found, err := Find(db, id)
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

	// Given
	generatedUuid := uuid.New()
	id := domain.StockItemId(generatedUuid)

	// When
	found, err := Find(db, id)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if found != false {
		t.Errorf("expected %t, got %t", false, found)
	}

}