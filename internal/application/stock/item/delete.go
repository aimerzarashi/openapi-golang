package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

func Delete(req *DeleteRequestDto, r item.IRepository) error {
	id := item.Id(req.Id)
	a, err := r.Get(id)
	if err != nil {
		return err
	}

	a.Delete()

	err = r.Save(a)
	if err != nil {
		return err
	}
	
	return nil
}