package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDishRepository struct {
	mock.Mock
}

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockDishRepository) GetDish(dishID int) (*entities.Dish, error) {
	args := m.Called(dishID)
	return args.Get(0).(*entities.Dish), args.Error(1)
}

func (m *MockDishRepository) GetOrder(orderID int) (*entities.Order, error) {
	args := m.Called(orderID)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (m *MockOrderRepository) AddOrderItem(orderItem *entities.OrderItem) error {
	args := m.Called(orderItem)
	return args.Error(0)
}

func TestCreateOrderItem(t *testing.T) {

	t.Run("should create an order item with success", func(t *testing.T) {
		orderID := 1
		dishID := 100
		amount := 2

		mockDishRepo := new(MockDishRepository)
		mockOrderRepo := new(MockOrderRepository)

		usecase := NewCreateOrderItemUseCase(mockDishRepo, mockOrderRepo)

		mockOrderRepo.On("GetOrder", orderID).Return(&entities.Order{ID: orderID}, nil)
		mockDishRepo.On("GetDish", dishID).Return(&entities.Dish{ID: dishID, Name: "Ramen"}, nil)
		mockOrderRepo.On("AddOrderItem", mock.Anything).Return(nil)

		orderItem, err := usecase.Create(orderID, dishID, amount)

		assert.NoError(t, err)
		assert.NotNil(t, orderItem)
		assert.Equal(t, orderID, orderItem.OrderID)
		assert.Equal(t, dishID, orderItem.Dish.ID)
		assert.Equal(t, amount, orderItem.Amount)

	})
}
