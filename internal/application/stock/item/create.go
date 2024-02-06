package item

import (
	"github.com/google/uuid"

	"openapi/internal/domain/stock/item"
)

type CreateRequestDto struct {
	Name string	
}

type CreateResponseDto struct {
	Id uuid.UUID
	Name string	
}

func Create(req *CreateRequestDto, r item.IRepository) (*CreateResponseDto, error) {
	itemName, err := item.NewItemName(req.Name)
	if err != nil {
		return nil, err
	}
	a, err := item.New(itemName)
	if err != nil {
		return nil, err
	}

	if err := r.Save(a); err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: a.GetId().UUID(),
		Name: a.GetName(),
	}, nil
}