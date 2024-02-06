package item

import (
	"context"
	"database/sql"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IRepository interface {
	Save(a *aggregate) error
	Get(id Id) (*aggregate, error)
	Find(id Id) (bool, error)
}

type Repository struct {
	IRepository
	Db *sql.DB
}

func (r *Repository) Save(a *aggregate) error {
	data := &sqlboiler.StockItem{
		ID:   a.id.UUID().String(),
		Name: a.name.String(),
		Deleted: a.deleted,
	}

	err := data.Upsert(
		context.Background(),
		r.Db,
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

func (r *Repository) Get(id Id) (*aggregate, error) {
	data, err := sqlboiler.FindStockItem(context.Background(), r.Db, id.UUID().String())
	if err != nil {
		return &aggregate{}, err
	}

	itemName, err := NewItemName(data.Name)
	if err != nil {
		return &aggregate{}, err
	}

	a := &aggregate{
		id:   id,
		name: *itemName,
		deleted: data.Deleted,
	}
	
	return a, nil
}


func (r *Repository) Find(id Id) (bool, error) {
	found, err := sqlboiler.StockItemExists(context.Background(), r.Db, id.UUID().String())
	if err != nil {
		return false, err
	}
	
	return found, nil
}