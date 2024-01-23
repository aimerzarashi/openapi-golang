package stock_item

import (
	"context"
	"database/sql"
	"openapi/internal/datastore"
	"openapi/internal/domain"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository struct {}

func (r *Repository) Store(db *sql.DB, stockItem *domain.StockItem) error {
	activeRecord := datastore.StockItem{
		ID: stockItem.Id.String(),
		Name: stockItem.Name,
	}
	activeRecord.Insert(context.Background(), db, boil.Infer())
	return nil
}

func (r *Repository) Get(db *sql.DB, id uuid.UUID) (*domain.StockItem, error) {
	activeRecord, findErr := datastore.FindStockItem(context.Background(), db, id.String())
	if findErr != nil {
		return nil, findErr
	}

	stockItem := &domain.StockItem{
		Id: uuid.MustParse(activeRecord.ID),
		Name: activeRecord.Name,
	}
	return stockItem, nil
}