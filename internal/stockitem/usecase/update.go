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

func Update(req *UpdateRequestDto, r repository.IRepository) (*UpdateResponseDto, error) {

	id := domain.StockItemId(req.Id)
	model, err := r.Get( id)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	model.Name = req.Name

	err = r.Save(model)
	if err != nil {
		return &UpdateResponseDto{}, err
	}
	
	return &UpdateResponseDto{}, nil
}