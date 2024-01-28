package repository

import (
	"context"
	"database/sql"

	"openapi/internal/infra/sqlboiler"
	"openapi/internal/stockitem/domain"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Save(db *sql.DB, model *domain.StockItem) error {

	// id := uuid.UUID(model.Id)
	// isExist, err := sqlboiler.StockItemExists(context.Background(), db, id.String())
	// if err != nil {
	// 	return err
	// }

	// data := &sqlboiler.StockItem{
	// 	ID:   uuid.UUID(model.Id).String(),
	// 	Name: model.Name,
	// }
	// if isExist {
	// 	data.Update(context.Background(), db, boil.Infer())
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
	// 	data.Insert(context.Background(), db, boil.Infer())
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	data := &sqlboiler.StockItem{
		ID:   uuid.UUID(model.Id).String(),
		Name: model.Name,
	}

	err := data.Upsert(
		context.Background(),
		db,
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

func Get(db *sql.DB, stockItemId domain.StockItemId) (*domain.StockItem, error) {

	id := uuid.UUID(stockItemId).String()
	data, err := sqlboiler.FindStockItem(context.Background(), db, id)
	if err != nil {
		return &domain.StockItem{}, err
	}

	model := domain.NewStockItem(stockItemId, data.Name)
	
	return model, nil
}


func Find(db *sql.DB, stockItemId domain.StockItemId) (bool, error) {
	id := uuid.UUID(stockItemId).String()
	found, err := sqlboiler.StockItemExists(context.Background(), db, id)
	if err != nil {
		return false, err
	}
	
	return found, nil
}