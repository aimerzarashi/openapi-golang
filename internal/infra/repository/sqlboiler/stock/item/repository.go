package item

import (
	"context"
	"database/sql"
	"errors"
	"openapi/internal/infra/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"openapi/internal/domain/stock/item"
)

type (
	repository struct {
		item.IRepository
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) (item.IRepository, error) {
	if db == nil {
		return nil, item.ErrIRepositoryDbEmpty
	}
	return &repository{
		db: db,
	}, nil
}

func (r *repository) Save(a *item.Aggregate) error {
	data := &sqlboiler.StockItem{
		ID:      a.Id.String(),
		Name:    a.Name.String(),
		Deleted: a.IsDeleted(),
	}

	err := data.Upsert(
		context.Background(),
		r.db,
		true,
		[]string{"id"},
		boil.Whitelist("name", "deleted"),
		boil.Infer(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Get(id item.Id) (*item.Aggregate, error) {
	data, err := sqlboiler.FindStockItem(context.Background(), r.db, id.UUID().String())
	if err != nil {
		return nil, err
	}

	if data.Deleted {
		return nil, item.ErrIRepositoryRowDeleted
	}

	name, err := item.NewName(data.Name)
	if err != nil {
		return nil, errors.Join(item.ErrIRepositoryInvalidData, err)
	}

	a := item.RestoreAggregate(id, name, data.Deleted)

	return a, nil
}

func (r *repository) Find(id item.Id) (bool, error) {
	data, err := sqlboiler.FindStockItem(context.Background(), r.db, id.UUID().String())
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	if data.Deleted {
		return false, nil
	}

	return true, nil
}
