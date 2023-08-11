package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDishUseCase_Create(t *testing.T) {

	t.Run("should be able to create a new dish", func(t *testing.T) {
		mockRepo := new(repomocks.MockDishRepository)
		du := NewDishUseCase(mockRepo)
		photo := &entities.Photo{
			URL: "https://www.google.com",
		}

		dishName := "Ramen"
		description := "Tasty ramen with vegetables"

		mockRepo.On("AddDish", mock.AnythingOfType("*entities.Dish")).Return(nil)

		dish, err := du.Create(dishName, description, photo)

		assert.NoError(t, err)
		assert.NotNil(t, dish)
		assert.Equal(t, dishName, dish.Name)
		assert.Equal(t, description, dish.Description)
		assert.Equal(t, photo, dish.Photo)
		mockRepo.AssertExpectations(t)

	})

	t.Run("should not be able to create a dish without a name", func(t *testing.T) {
		mockRepo := new(repomocks.MockDishRepository)
		du := NewDishUseCase(mockRepo)
		photo := &entities.Photo{
			URL: "https://www.google.com",
		}
		dishName := ""
		description := "Tasty ramen with vegetables"

		dish, err := du.Create(dishName, description, photo)

		assert.Error(t, err)
		assert.Nil(t, dish)
		assert.Equal(t, err.Error(), "name is required")
	})
}
