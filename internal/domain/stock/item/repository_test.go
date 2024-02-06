package item_test

import (
	"context"
	"testing"
	"time"

	"openapi/internal/domain/stock/item"

	"openapi/internal/infrastructure/database"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &item.Repository{Db: db}

	// Given
	name, err := item.NewItemName(uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.New(name)
	if err != nil {
		t.Fatal(err)
	}
	currentDateTime := time.Now().UTC()

	// When
	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	data, err := sqlboiler.FindStockItem(context.Background(), db, a.GetId().UUID().String())
	if err != nil {
		t.Fatal(err)
	}

	if data.ID != a.GetId().UUID().String() {
		t.Errorf("expected %s, got %s", a.GetId().UUID().String(), data.ID)
	}
	
	if data.Name != name.String() {
		t.Errorf("expected %s, got %s", name, data.Name)
	}

	if data.Deleted != false {
		t.Errorf("expected %t, got %t", false, data.Deleted)
	}

	if data.CreatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, data.CreatedAt)		
	}

	if data.UpdatedAt.Equal(data.CreatedAt) != true {
		t.Errorf("expected %s, got %s", data.CreatedAt, data.UpdatedAt)
	}
}

func TestUpdate(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &item.Repository{Db: db}

	// Given
	beforeName, err := item.NewItemName(uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}
	afterName, err := item.NewItemName(uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.New(beforeName)
	if err != nil {
		t.Fatal(err)
	}
	currentDateTime := time.Now().UTC()
	dataFormat := "2006-01-02 15:04:05.000000 +09:00"

	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	beforeData, err := sqlboiler.FindStockItem(context.Background(), db, a.GetId().UUID().String())
	if err != nil {
		t.Fatal(err)
	}

	// When
	a.ChangeName(afterName)
	a.Delete()
	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	afterData, err := sqlboiler.FindStockItem(context.Background(), db, a.GetId().UUID().String())
	if err != nil {
		t.Fatal(err)
	}

	if afterData.ID != beforeData.ID{
		t.Errorf("expected %s, got %s", beforeData.ID, afterData.ID)
	}

	if afterData.Name != a.GetName() {
		t.Errorf("expected %s, got %s", a.GetName(), afterData.Name)
	}

	if afterData.Deleted != a.IsDeleted() {
		t.Errorf("expected %t, got %t", a.IsDeleted(), afterData.Deleted)
	}

	if afterData.CreatedAt.Format(dataFormat) != beforeData.CreatedAt.Format(dataFormat) {
		t.Errorf("expected %s, got %s", beforeData.CreatedAt.Format(dataFormat), afterData.CreatedAt.Format(dataFormat))
	}

	if afterData.UpdatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, afterData.UpdatedAt)
	}
}

func TestFind(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := &item.Repository{Db: db}

	// Given
	name, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.New(name)
	if err != nil {
		t.Fatal(err)
	}

	// When
	beforeFound, err := r.Find(a.GetId())
	if err != nil {
		t.Fatal(err)
	}
	err = r.Save(a)
	if err != nil {
		t.Fatal(err)
	}
	afterFound, err := r.Find(a.GetId())
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if beforeFound != false {
		t.Errorf("expected %t, got %t", false, beforeFound)
	}
	if afterFound != true {
		t.Errorf("expected %t, got %t", true, afterFound)
	}
}