package repomocks

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockDishRepository struct {
	mock.Mock
}

func (m *MockDishRepository) GetDish(dishID int) (*entities.Dish, error) {
	args := m.Called(dishID)
	return args.Get(0).(*entities.Dish), args.Error(1)
}

func (m *MockDishRepository) ListDishesByType(dishType entities.DishType) ([]*entities.Dish, error) {
	args := m.Called(dishType)
	return args.Get(0).([]*entities.Dish), args.Error(1)
}

func (m *MockDishRepository) AddDish(dish *entities.Dish) error {
	args := m.Called(dish)
	return args.Error(0)
}

func (m *MockDishRepository) UpdateDish(dish *entities.Dish) error {
	args := m.Called(dish)
	return args.Error(0)
}

func (m *MockDishRepository) DeleteDish(dishID int) error {
	args := m.Called(dishID)
	return args.Error(0)
}
