package usecase

import (
	"database/sql"

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

func Create(req *CreateRequestDto, db *sql.DB) (*CreateResponseDto, error) {

	id := domain.StockItemId(uuid.New())
	model := domain.NewStockItem(id, req.Name)

	err := repository.Save(db, model)
	if err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: uuid.UUID(model.Id),
	}, nil
}