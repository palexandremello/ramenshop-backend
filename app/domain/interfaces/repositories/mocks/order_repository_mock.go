package repomocks

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *entities.Order) error {
	args := m.Called(order)
	return args.Error(0)

}

func (m *MockOrderRepository) AddOrderItem(orderItem *entities.OrderItem) error {
	args := m.Called(orderItem)
	return args.Error(0)
}

func (m *MockOrderRepository) List() ([]entities.Order, error) {
	args := m.Called()
	return args.Get(0).([]entities.Order), args.Error(1)
}

func (m *MockOrderRepository) GetOrder(orderID int) (*entities.Order, error) {
	args := m.Called(orderID)
	return args.Get(0).(*entities.Order), args.Error(1)
}
