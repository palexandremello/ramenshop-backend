package repomocks

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockTableRepository struct {
	mock.Mock
}

func (m *MockTableRepository) FindByID(id int) (*entities.Table, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Table), args.Error(1)
}

// List implements repositories.TableRepository.
func (*MockTableRepository) List() ([]*entities.Table, error) {
	panic("unimplemented")
}

// Remove implements repositories.TableRepository.
func (*MockTableRepository) Remove(id int) error {
	panic("unimplemented")
}

func (m *MockTableRepository) Update(table *entities.Table) error {
	args := m.Called(table)
	return args.Error(0)
}

func (m *MockTableRepository) Add(table *entities.Table) (*entities.Table, error) {
	args := m.Called(table)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Table), args.Error(1)
}
