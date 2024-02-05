package item_test

import (
	"fmt"

	domain "openapi/internal/domain/stock/item"
)

type MockRepository struct {
	domain.IRepository
}

func (m *MockRepository) Save(a *domain.Aggregate) error {
	return fmt.Errorf("not implemented")
}

func (m *MockRepository) Get(id domain.Id) (*domain.Aggregate, error) {
	return nil, fmt.Errorf("not implemented")
}

func (m *MockRepository) Delete(id domain.Id) error {
	return fmt.Errorf("not implemented")
}
