package usecases

import (
	"errors"
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderItem(t *testing.T) {

	t.Run("should create an order item with success", func(t *testing.T) {
		orderID := 1
		dishID := 100
		amount := 2

		mockDishRepo := new(repomocks.MockDishRepository)
		mockOrderRepo := new(repomocks.MockOrderRepository)

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

	t.Run("should return an error if amount is less then 0", func(t *testing.T) {
		mockDishRepo := new(repomocks.MockDishRepository)
		mockOrderRepo := new(repomocks.MockOrderRepository)

		usecase := NewCreateOrderItemUseCase(mockDishRepo, mockOrderRepo)

		_, err := usecase.Create(1, 1, 0)
		assert.Error(t, err)
		assert.Equal(t, "amount should be greater than 0", err.Error())

	})

	t.Run("should return an error if order not found", func(t *testing.T) {
		mockDishRepo := new(repomocks.MockDishRepository)
		mockOrderRepo := new(repomocks.MockOrderRepository)

		usecase := NewCreateOrderItemUseCase(mockDishRepo, mockOrderRepo)
		mockOrderRepo.On("GetOrder", mock.Anything).Return((*entities.Order)(nil), errors.New("Order not found"))
		_, err := usecase.Create(1, 1, 2)
		assert.Error(t, err)
		assert.Equal(t, "Order not found", err.Error())
	})

	t.Run("should return an error if GetDish fails", func(t *testing.T) {
		mockDishRepo := new(repomocks.MockDishRepository)
		mockOrderRepo := new(repomocks.MockOrderRepository)

		usecase := NewCreateOrderItemUseCase(mockDishRepo, mockOrderRepo)

		// Mock GetOrder
		mockOrderRepo.On("GetOrder", 1).Return(&entities.Order{ID: 1}, nil)

		// Mock GetDish
		mockDishRepo.On("GetDish", 1).Return(nil, errors.New("any_error"))

		_, err := usecase.Create(1, 1, 2)
		assert.Error(t, err)
		assert.Equal(t, "any_error", err.Error())
	})

	t.Run("should return an error if Dish not found", func(t *testing.T) {
		mockDishRepo := new(repomocks.MockDishRepository)
		mockOrderRepo := new(repomocks.MockOrderRepository)

		usecase := NewCreateOrderItemUseCase(mockDishRepo, mockOrderRepo)

		// Mock GetOrder
		mockOrderRepo.On("GetOrder", 1).Return(&entities.Order{ID: 1}, nil)

		// Mock GetDish
		mockDishRepo.On("GetDish", 1).Return(nil, nil)

		_, err := usecase.Create(1, 1, 2)
		assert.Error(t, err)
		assert.Equal(t, "Dish not found", err.Error())
	})

	t.Run("should return an error if failed to add an order item", func(t *testing.T) {
		orderID := 1
		dishID := 100
		amount := 2

		mockDishRepo := new(repomocks.MockDishRepository)
		mockOrderRepo := new(repomocks.MockOrderRepository)

		usecase := NewCreateOrderItemUseCase(mockDishRepo, mockOrderRepo)

		mockOrderRepo.On("GetOrder", orderID).Return(&entities.Order{ID: orderID}, nil)
		mockDishRepo.On("GetDish", dishID).Return(&entities.Dish{ID: dishID, Name: "Ramen"}, nil)
		mockOrderRepo.On("AddOrderItem", mock.Anything).Return(errors.New("Failed to add order item"))

		_, err := usecase.Create(orderID, dishID, amount)
		assert.Error(t, err)
		assert.Equal(t, "Failed to add order item", err.Error())
	})

}
