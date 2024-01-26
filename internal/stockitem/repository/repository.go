package repository

import (
	"context"

	"openapi/internal/infra/database"
	"openapi/internal/infra/sqlboiler"
	"openapi/internal/stockitem/domain"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Save(stockItem *domain.StockItem) error {
	db, err := database.New()
	if err != nil {
		return err
	}
	defer db.Close()

	id := uuid.UUID(stockItem.Id)
	isExist, err := sqlboiler.StockItemExists(context.Background(), db, id.String())
	if err != nil {
		return err
	}

	data := &sqlboiler.StockItem{
		ID:   uuid.UUID(stockItem.Id).String(),
		Name: stockItem.Name,
	}
	if isExist {
		data.Update(context.Background(), db, boil.Infer())
		if err != nil {
			return err
		}
	} else {
		data.Insert(context.Background(), db, boil.Infer())
		if err != nil {
			return err
		}
	}

	return nil
}

func Get(id domain.StockItemId) (*domain.StockItem, error) {
	db, err := database.New()
	if err != nil {
		return &domain.StockItem{}, err
	}
	defer db.Close()

	uuid := uuid.UUID(id)
	data, err := sqlboiler.FindStockItem(context.Background(), db, uuid.String())
	if err != nil {
		return &domain.StockItem{}, err
	}

	stockItem := domain.NewStockItem(domain.StockItemId(uuid), data.Name)
	
	return stockItem, nil
}