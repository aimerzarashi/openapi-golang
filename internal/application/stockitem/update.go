package stockitem

import (
	"openapi/internal/domain/repository"

	"openapi/internal/domain/model"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string	
}

type UpdateResponseDto struct {}

func Update(req *UpdateRequestDto, r repository.IStockItem) (*UpdateResponseDto, error) {

	id := model.StockItemId(req.Id)
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