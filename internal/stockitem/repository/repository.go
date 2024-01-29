package repository

import (
	"context"
	"database/sql"

	"openapi/internal/infra/sqlboiler"
	"openapi/internal/stockitem/domain"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IRepository interface {
	Save(model *domain.StockItem) error
	Get(stockItemId domain.StockItemId) (*domain.StockItem, error)
	Find(stockItemId domain.StockItemId) (bool, error)
}

type Repository struct {
	IRepository
	*sql.DB
}

func (r *Repository) Save(model *domain.StockItem) error {

	data := &sqlboiler.StockItem{
		ID:   uuid.UUID(model.Id).String(),
		Name: model.Name,
	}

	err := data.Upsert(
		context.Background(),
		r.DB,
		true,
		[]string{"id"},
		boil.Whitelist("name","deleted"),
		boil.Infer(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get(stockItemId domain.StockItemId) (*domain.StockItem, error) {

	id := uuid.UUID(stockItemId).String()
	data, err := sqlboiler.FindStockItem(context.Background(), r.DB, id)
	if err != nil {
		return &domain.StockItem{}, err
	}

	model := domain.NewStockItem(stockItemId, data.Name)
	
	return model, nil
}


func (r *Repository) Find(stockItemId domain.StockItemId) (bool, error) {
	id := uuid.UUID(stockItemId).String()
	found, err := sqlboiler.StockItemExists(context.Background(), r.DB, id)
	if err != nil {
		return false, err
	}
	
	return found, nil
}