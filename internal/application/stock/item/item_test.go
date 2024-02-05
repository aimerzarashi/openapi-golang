package item_test

import (
	"fmt"

	domain "openapi/internal/domain/stock/item"
)

type MockRepository struct {
	domain.IRepository
}

var ErrMockRepository = fmt.Errorf("mock repository error")

func (m *MockRepository) Save(a *domain.Aggregate) error {
	return ErrMockRepository
}

func (m *MockRepository) Get(id domain.Id) (*domain.Aggregate, error) {
	return nil, ErrMockRepository
}

func (m *MockRepository) Delete(id domain.Id) error {
	return ErrMockRepository
}
