package location

import (
	"github.com/google/uuid"

	"openapi/internal/domain/stock/location"
)

type CreateRequestDto struct {
	Name string	
}

type CreateResponseDto struct {
	Id uuid.UUID
	Name string	
}

func Create(req *CreateRequestDto, r location.IRepository) (*CreateResponseDto, error) {

	a, err := location.New(req.Name)
	if err != nil {
		return nil, err
	}

	if err := r.Save(a); err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: a.GetId().UUID(),
		Name: a.GetName(),
	}, nil
}