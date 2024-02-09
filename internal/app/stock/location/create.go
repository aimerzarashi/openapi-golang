package location

import (
	"github.com/google/uuid"

	"openapi/internal/domain/stock/location"
)

type createRequest struct {
	Name location.Name	
}

type createResponse struct {
	Id uuid.UUID
}

func Create(req createRequest, r location.IRepository, newId uuid.UUID) (createResponse, error) {
	id, err := location.NewId(newId)
	if err != nil {
		return createResponse{}, err
	}

	a := location.NewAggregate(id, req.Name)

	if err := r.Save(a); err != nil {
		return createResponse{}, err
	}

	res := newCreateResponse(a.Id, a.Name)
	return res, nil
}

func NewCreateRequest(name string) (createRequest, error) {
	validName, err := location.NewName(name)
	if err != nil {
		return createRequest{}, err
	}

	return createRequest{
		Name: validName,
	}, nil
}

func newCreateResponse(id location.Id, name location.Name) createResponse {
	return createResponse{
		Id: id.UUID(),
	}
}