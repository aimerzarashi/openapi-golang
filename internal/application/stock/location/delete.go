package location

import (
	"openapi/internal/domain/stock/location"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

func Delete(req *DeleteRequestDto, r location.IRepository) error {
	id := location.Id(req.Id)
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