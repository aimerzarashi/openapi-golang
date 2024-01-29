package usecase

import (
	"openapi/internal/stockitem/domain"
	"openapi/internal/stockitem/repository"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

type DeleteResponseDto struct {
}

func Delete(req *DeleteRequestDto, r repository.IRepository) (*DeleteResponseDto, error) {
	
	id := domain.StockItemId(req.Id)
	model, err := r.Get(id)
	if err != nil {
		return &DeleteResponseDto{}, err
	}

	model.Deleted = true

	err = r.Save( model)
	if err != nil {
		return &DeleteResponseDto{}, err
	}
	
	return &DeleteResponseDto{}, nil
}