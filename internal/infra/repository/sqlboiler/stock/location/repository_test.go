package location_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
	"openapi/internal/infra/sqlboiler"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestNewRepository(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// When
	repo, err := infra.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if repo == nil {
		t.Errorf("expected not nil, got nil")
	}
}

func TestNewRepositoryFail(t *testing.T) {
	t.Parallel()

	// When
	repo, err := infra.NewRepository(nil)

	// Then
	if err == nil {
		t.Fatal("expected error but returned nil")
	}

	if repo != nil {
		t.Errorf("expected nil, got not nil")
	}
}

func TestSave(t *testing.T) {
	t.Parallel()

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

	currentDateTime := time.Now().UTC()

	// Given
	id, err := domain.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := domain.NewName("TestName")
	if err != nil {
		t.Fatal(err)
	}

	a := domain.NewAggregate(id, name)

	// When
	before, err := repo.Get(a.Id)
	if err == nil {
		t.Fatalf("expected error but returned nil, %+v", before)
	}

	if err = repo.Save(a); err != nil {
		t.Fatal(err)
	}

	after, err := repo.Get(a.Id)
	if err != nil {
		t.Fatalf("expected error but returned nil, %+v", err)
	}

	// Then
	if reflect.DeepEqual(after,before) {
		t.Errorf("%T %+v want %+v", after, after, before)
	}

	data, err := sqlboiler.FindStockLocation(context.Background(), db, a.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	if data.ID != id.String() {
		t.Errorf("%T %+v want %+v", data.ID, data.ID, id)
	}
	
	if data.Name != name.String() {
		t.Errorf("%T %+v want %+v", data.Name, data.Name, name)
	}

	if data.Deleted != false {
		t.Errorf("%T %+v want %+v", data.Deleted, data.Deleted, false)
	}

	if data.CreatedAt.Before(currentDateTime) == true {
		t.Errorf("expected %s, got %s", currentDateTime, data.CreatedAt)		
	}

	if data.UpdatedAt.Equal(data.CreatedAt) != true {
		t.Errorf("expected %s, got %s", data.CreatedAt, data.UpdatedAt)
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

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
	id, err := domain.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := domain.NewName("before")
	if err != nil {
		t.Fatal(err)
	}

	before := domain.NewAggregate(id, name)

	currentDateTime := time.Now().UTC()
	dataFormat := "2006-01-02 15:04:05.000000 +09:00"

	if err = repo.Save(before); err != nil {
		t.Fatal(err)
	}

	beforeData, err := sqlboiler.FindStockLocation(context.Background(), db, before.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	// When
	after, err := repo.Get(before.Id)
	if err != nil {
		t.Fatal(err)
	}

	changedName, err := domain.NewName("after")
	if err != nil {
		t.Fatal(err)
	}

	after.Name = changedName
	after.Delete()

	if err = repo.Save(after); err != nil {
		t.Fatal(err)
	}

	// Then
	afterData, err := sqlboiler.FindStockLocation(context.Background(), db, after.Id.String())
	if err != nil {
		t.Fatal(err)
	}

	if afterData.ID != after.Id.String(){
		t.Errorf("%T %+v want %+v", afterData.ID, afterData.ID, after.Id.String())
	}

	if afterData.Name != after.Name.String() {
		t.Errorf("%T %+v want %+v", afterData.Name, afterData.Name, after.Name.String())
	}

	if afterData.Deleted != after.IsDeleted() {
		t.Errorf("%T %+v want %+v", afterData.Deleted, afterData.Deleted, after.IsDeleted())
	}

	if afterData.CreatedAt.Format(dataFormat) != beforeData.CreatedAt.Format(dataFormat) {
		t.Errorf("%T %+v want %+v", afterData.CreatedAt, afterData.CreatedAt, beforeData.CreatedAt.Format(dataFormat))
	}

	if afterData.UpdatedAt.Before(currentDateTime) == true {
		t.Errorf("%T %+v want greater than %+v ", afterData.UpdatedAt, afterData.UpdatedAt, currentDateTime)
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

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
	id, err := domain.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := domain.NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	a := domain.NewAggregate(id, name)

	// When
	notFound, err := repo.Find(a.Id)
	if err != nil {
		t.Fatal(err)
	}

	if err = repo.Save(a); err != nil {
		t.Fatal(err)
	}

	found, err := repo.Find(a.Id)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if notFound != false {
		t.Errorf("%T %+v want %+v", notFound, notFound, false)
	}

	if found != true {
		t.Errorf("%T %+v want %+v", found, found, true)
	}
}

func TestGetFailInvalidData(t *testing.T) {
	t.Parallel()

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
	id, err := domain.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := domain.NewName("TestName")
	if err != nil {
		t.Fatal(err)
	}

	a := domain.NewAggregate(id, name)

	if err := repo.Save(a); err != nil {
		t.Fatal(err)
	}

	data := &sqlboiler.StockLocation{
		ID:   a.Id.String(),
		Name: "",
		Deleted: a.IsDeleted(),
	}

	err = data.Upsert(
		context.Background(),
		db,
		true,
		[]string{"id"},
		boil.Whitelist("name","deleted"),
		boil.Infer(),
	)
	if err != nil {
		t.Fatal(err)
	}

	// When
	_, err = repo.Get(id)

	// Then
	if err == nil {
		t.Fatal("expected error but returned nil")
	}
}

func TestFailDbClose(t *testing.T) {
	t.Parallel()

	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	db.Close()

	repo, err := infra.NewRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	// Given
	id, err := domain.NewId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	name, err := domain.NewName("test")
	if err != nil {
		t.Fatal(err)
	}

	a := domain.NewAggregate(id, name)

	// Then
	_, err = repo.Get(a.Id)
	if err == nil {
		t.Fatal("expected error but returned nil")
	}

	_, err = repo.Find(a.Id)
	if err == nil {
		t.Fatal("expected error but returned nil")
	}

	err = repo.Save(a)
	if err == nil {
		t.Fatal("expected error but returned nil")
	}
}