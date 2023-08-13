package usecases

import (
	"testing"
	"time"

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

	t.Run("should not add items to a closed order", func(t *testing.T) {
		now := time.Now()
		mockOrderRepo := new(repomocks.MockOrderRepository)
		usecase := NewAddOrderItems(mockOrderRepo)
		closedOrder := &entities.Order{ID: 1, ClosedAt: &now}
		newItems := []entities.OrderItem{
			{OrderID: 1, Dish: entities.Dish{
				ID:   1,
				Name: "Ramen",
			}},
		}
		mockOrderRepo.On("GetOrder", 1).Return(closedOrder, nil)

		err := usecase.Add(1, newItems)

		assert.NotNil(t, err)
		assert.Equal(t, "cannot add items to a closed order", err.Error())
	})
}
