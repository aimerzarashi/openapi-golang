package usecase

import (
	"github.com/google/uuid"

	"openapi/internal/stockitem/domain"
	"openapi/internal/stockitem/repository"
)

type CreateRequestDto struct {
	Name string	
}

type CreateResponseDto struct {
	Id uuid.UUID
}

func Create(req *CreateRequestDto, r repository.IRepository) (*CreateResponseDto, error) {

	id := domain.StockItemId(uuid.New())
	model := domain.NewStockItem(id, req.Name)

	err := r.Save(model)
	if err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: uuid.UUID(model.Id),
	}, nil
}