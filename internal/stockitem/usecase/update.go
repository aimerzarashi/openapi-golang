package usecase

import (
	"database/sql"
	"openapi/internal/stockitem/repository"

	"openapi/internal/stockitem/domain"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string	
}

type UpdateResponseDto struct {}

func Update(req *UpdateRequestDto, db *sql.DB) (*UpdateResponseDto, error) {

	id := domain.StockItemId(req.Id)
	model, err := repository.Get(db, id)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	model.Name = req.Name

	err = repository.Save(db, model)
	if err != nil {
		return &UpdateResponseDto{}, err
	}
	
	return &UpdateResponseDto{}, nil
}