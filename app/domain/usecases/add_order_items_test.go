package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddOrderItems(t *testing.T) {

	t.Run("should add items to an open order", func(t *testing.T) {
		mockOrderRepo := new(repomocks.MockOrderRepository)
		usecase := NewAddOrderItems(mockOrderRepo)

		order := &entities.Order{
			ID:    1,
			Items: []entities.OrderItem{},
		}
		newItems := []entities.OrderItem{
			{OrderID: 1, Dish: entities.Dish{ID: 2, Name: "Ramen"}},
		}

		mockOrderRepo.On("GetOrder", 1).Return(order, nil)
		mockOrderRepo.On("Update", order).Return(nil)

		err := usecase.Add(1, newItems)

		assert.Nil(t, err)
		assert.Len(t, order.Items, 1)

	})
}
