package usecases

import (
	"errors"
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestViewOrders(t *testing.T) {
	t.Run("should return all orders", func(t *testing.T) {
		mockRepo := new(repomocks.MockOrderRepository)
		useCase := NewViewOrders(mockRepo)
		expectedOrders := []entities.Order{
			{ID: 1},
			{ID: 2},
		}

		mockRepo.On("List").Return(expectedOrders, nil)

		orders, err := useCase.GetAllOrders()

		assert.NoError(t, err)
		assert.Equal(t, expectedOrders, orders)

	})

	t.Run("should return an error when repository fails", func(t *testing.T) {
		mockRepo := new(repomocks.MockOrderRepository)
		useCase := NewViewOrders(mockRepo)
		mockRepo.On("List").Return(nil, errors.New("database error"))

		_, err := useCase.GetAllOrders()

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

	})

	t.Run("should return a specific order", func(t *testing.T) {
		mockRepo := new(repomocks.MockOrderRepository)
		useCase := NewViewOrders(mockRepo)
		expectedOrder := &entities.Order{ID: 1}
		mockRepo.On("GetOrder", 1).Return(expectedOrder, nil)

		order, err := useCase.GetOrder(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedOrder, order)

	})
}
