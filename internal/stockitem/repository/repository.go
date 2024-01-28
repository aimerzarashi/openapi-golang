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

func Get(db *sql.DB, id domain.StockItemId) (*domain.StockItem, error) {

	uuid := uuid.UUID(id)
	data, err := sqlboiler.FindStockItem(context.Background(), db, uuid.String())
	if err != nil {
		return &domain.StockItem{}, err
	}

	model := domain.NewStockItem(domain.StockItemId(uuid), data.Name)
	
	return model, nil
}


func Find(db *sql.DB, id domain.StockItemId) (bool, error) {
	uuid := uuid.UUID(id).String()
	found, err := sqlboiler.StockItemExists(context.Background(), db, uuid)
	if err != nil {
		return false, err
	}
	
	return found, nil
}