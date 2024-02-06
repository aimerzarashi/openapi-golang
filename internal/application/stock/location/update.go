package location

import (
	"openapi/internal/domain/stock/location"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string	
}

func Update(req *UpdateRequestDto, r location.IRepository) error {
	id := location.Id(req.Id)
	a, err := r.Get(id)
	if err != nil {
		return err
	}

	a.ChangeName(req.Name)

	err = r.Save(a)
	if err != nil {
		return err
	}
	
	return nil
}