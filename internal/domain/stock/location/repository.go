package location

import (
	"context"
	"database/sql"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IRepository interface {
	Save(a *Aggregate) error
	Get(id Id) (*Aggregate, error)
	Find(id Id) (bool, error)
}

type Repository struct {
	IRepository
	Db *sql.DB
}

func (r *Repository) Save(a *Aggregate) error {
	data := &sqlboiler.StockLocation{
		ID:   a.id.UUID().String(),
		Name: a.name,
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

func (r *Repository) Get(id Id) (*Aggregate, error) {
	data, err := sqlboiler.FindStockLocation(context.Background(), r.Db, id.UUID().String())
	if err != nil {
		return &Aggregate{}, err
	}

	a := &Aggregate{
		id:   id,
		name: data.Name,
		deleted: data.Deleted,
	}
	
	return a, nil
}


func (r *Repository) Find(id Id) (bool, error) {
	found, err := sqlboiler.StockLocationExists(context.Background(), r.Db, id.UUID().String())
	if err != nil {
		return false, err
	}
	
	return found, nil
}