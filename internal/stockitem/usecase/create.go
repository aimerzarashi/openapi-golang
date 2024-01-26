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

func Create(r *CreateRequestDto) (*CreateResponseDto, error) {

	id := domain.StockItemId(uuid.New())
	model := domain.NewStockItem(id, r.Name)

	err := repository.Save(model)
	if err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: uuid.New(),
	}, nil
}