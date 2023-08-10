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

func (m *MockDishRepository) Save(dish *entities.Dish) error {
	args := m.Called(dish)
	return args.Error(0)
}

func (m *MockDishRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockDishRepository) Update(dish *entities.Dish) error {
	args := m.Called(dish)
	return args.Error(0)
}

func (m *MockDishRepository) List() ([]entities.Dish, error) {
	args := m.Called()
	return args.Get(0).([]entities.Dish), args.Error(1)
}

func (m *MockDishRepository) GetByID(id int) (*entities.Dish, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Dish), args.Error(1)
}

func TestDishUseCase_Create(t *testing.T) {

	t.Run("should be able to create a new dish", func(t *testing.T) {
		mockRepo := new(MockDishRepository)
		du := NewDishUseCase(mockRepo)
		photo := &entities.Photo{
			URL: "https://www.google.com",
		}

		dishName := "Ramen"
		description := "Tasty ramen with vegetables"

		mockRepo.On("Save", mock.AnythingOfType("*entities.Dish")).Return(nil)

		dish, err := du.Create(dishName, description, photo)

		assert.NoError(t, err)
		assert.NotNil(t, dish)
		assert.Equal(t, dishName, dish.Name)
		assert.Equal(t, description, dish.Description)
		assert.Equal(t, photo, dish.Photo)
		mockRepo.AssertExpectations(t)

	})
}
