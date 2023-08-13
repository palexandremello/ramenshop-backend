package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCloseOrder(t *testing.T) {

	t.Run("should to be able to close an order successfully", func(t *testing.T) {
		mockOrderRepo := new(repomocks.MockOrderRepository)
		usecase := NewCloseOrder(mockOrderRepo)

		order := &entities.Order{
			ID:     1,
			Client: entities.Client{ID: 1, Name: "Paulo", Gender: 1, Age: 29},
		}

		mockOrderRepo.On("GetOrder", 1).Return(order, nil)
		mockOrderRepo.On("Update", order).Return(nil)

		err := usecase.Execute(1)

		assert.Nil(t, err)
		assert.NotNil(t, order.ClosedAt)
	})
}
