package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string	
}

type UpdateResponseDto struct {
	Id   uuid.UUID
	Name string	
}

func Update(req *UpdateRequestDto, r item.IRepository) (*UpdateResponseDto, error) {

	id := item.Id(req.Id)
	a, err := r.Get( id)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	a.ChangeName(req.Name)

	err = r.Save(a)
	if err != nil {
		return &UpdateResponseDto{}, err
	}
	
	return &UpdateResponseDto{
		Id: a.GetId().UUID(),
		Name: a.GetName(),
	}, nil
}