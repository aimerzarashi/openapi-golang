package usecase

import (
	"database/sql"
	"openapi/internal/stockitem/domain"
	"openapi/internal/stockitem/repository"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

type DeleteResponseDto struct {
}

func Delete(req *DeleteRequestDto, db *sql.DB) (*DeleteResponseDto, error) {
	
	id := domain.StockItemId(req.Id)
	model, err := repository.Get(db, id)
	if err != nil {
		return &DeleteResponseDto{}, err
	}

	model.Deleted = true

	err = repository.Save(db, model)
	if err != nil {
		return &DeleteResponseDto{}, err
	}
	
	return &DeleteResponseDto{}, nil
}