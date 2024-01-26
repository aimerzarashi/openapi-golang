package usecase

import (
	"openapi/internal/stockitem/repository"

	"openapi/internal/stockitem/domain"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string	
}

type UpdateResponseDto struct {}

func Update(r *UpdateRequestDto) (*UpdateResponseDto, error) {

	id := domain.StockItemId(r.Id)
	model, err := repository.Get(id)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	model.Name = r.Name

	err = repository.Save(model)
	if err != nil {
		return &UpdateResponseDto{}, err
	}
	
	return &UpdateResponseDto{}, nil
}