package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

type DeleteResponseDto struct {
}

func Delete(req *DeleteRequestDto, r item.IRepository) (*DeleteResponseDto, error) {
	
	id := item.Id(req.Id)
	a, err := r.Get(id)
	if err != nil {
		return &DeleteResponseDto{}, err
	}

	a.Delete()

	err = r.Save(a)
	if err != nil {
		return &DeleteResponseDto{}, err
	}
	
	return &DeleteResponseDto{}, nil
}