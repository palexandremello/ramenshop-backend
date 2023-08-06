package usecases

import (
	"errors"
	"fmt"
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *entities.Order) error {
	fmt.Println("Save method in mock is called")
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepository) List() ([]entities.Order, error) {

	args := m.Called()
	return args.Get(0).([]entities.Order), args.Error(1)
}

func TestOrderCreator(t *testing.T) {
	client := entities.Client{ID: 1, Name: "Daft Punk", Gender: entities.Male, Age: 29}
	items := []entities.OrderItem{
		{ID: 1, Dish: entities.Dish{ID: 1, Name: "Ramen"}, Amount: 1},
	}

	t.Run("should be able to create a new order", func(t *testing.T) {
		mockRepo := new(MockOrderRepository)
		NewOrderCreator()
		uc := &orderCreatorImpl{orderRepo: mockRepo}

		mockRepo.On("Save", mock.AnythingOfType("*entities.Order")).Return(nil)

		order, err := uc.CreateOrder(client, items)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, client, order.Client)
		assert.Equal(t, items, order.Items)
		mockRepo.AssertExpectations(t)
	})
	t.Run("should return an error when repository fails", func(t *testing.T) {
		mockRepo := new(MockOrderRepository)

		uc := &orderCreatorImpl{orderRepo: mockRepo}

		mockRepo.On("Save", mock.AnythingOfType("*entities.Order")).Return(errors.New("DB error"))

		order, err := uc.CreateOrder(client, []entities.OrderItem{})

		assert.Error(t, err)
		assert.Nil(t, order)
		assert.Equal(t, err.Error(), "DB error")
		mockRepo.AssertExpectations(t)
	})

}
