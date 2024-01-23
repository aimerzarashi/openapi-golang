package stock_item

import (
	"database/sql"
	"openapi/internal/domain"
	"testing"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStockItem(t *testing.T) {
	db, dbOpenErr := sql.Open(
		"postgres",
		"host=openapi-db port=5432 user=user password=password dbname=openapi sslmode=disable",
	)
	if dbOpenErr != nil {
		t.Fatal(dbOpenErr)
	}
	defer db.Close()
	
	repository := &Repository{}
	stockItem := &domain.StockItem{
		Id: uuid.New(),
		Name: uuid.NewString(),
	}
	err := repository.Store(db, stockItem)
	if err != nil {
		t.Fatal(err)
	}

	stockItem2, err := repository.Get(db, stockItem.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, stockItem, stockItem2)
}