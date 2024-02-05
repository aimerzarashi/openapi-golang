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

	a := location.New(req.Name)

	err := r.Save(a)
	if err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: a.GetId().UUID(),
		Name: a.GetName(),
	}, nil
}