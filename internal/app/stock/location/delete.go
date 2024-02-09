package location

import (
	"openapi/internal/domain/stock/location"

	"github.com/google/uuid"
)

type deleteRequest struct {
	Id location.Id
}

func NewDeleteRequest(id uuid.UUID) (deleteRequest, error) {
	validId, err := location.NewId(id)
	if err != nil {
		return deleteRequest{}, err
	}

	return deleteRequest{
		Id: validId,
	}, nil
}

func Delete(req deleteRequest, r location.IRepository) error {
	a, err := r.Get(req.Id)
	if err != nil {
		return err
	}

	a.Delete()

	err = r.Save(a)
	if err != nil {
		return err
	}
	
	return nil
}