package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string	
}

func Update(req *UpdateRequestDto, r item.IRepository) error {
	id := item.Id(req.Id)
	a, err := r.Get( id)
	if err != nil {
		return err
	}

	itemName, err := item.NewItemName(req.Name)
	if err != nil {
		return err
	}
	a.ChangeName(itemName)

	err = r.Save(a)
	if err != nil {
		return err
	}
	
	return nil
}