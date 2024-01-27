package usecase

import (
	"database/sql"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

type DeleteResponseDto struct {
}

func Delete(reqDto *DeleteRequestDto, db *sql.DB) (*DeleteResponseDto, error) {

	return &DeleteResponseDto{}, nil
}