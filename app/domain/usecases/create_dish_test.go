package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreatePhoto struct {
	mock.Mock
}

func (m *MockCreatePhoto) Create(url string) (*entities.Photo, error) {
	args := m.Called(url)
	return args.Get(0).(*entities.Photo), args.Error(1)
}

func TestDishUseCase_Create(t *testing.T) {

	t.Run("should be able to create a new dish", func(t *testing.T) {
		mockDishRepo := new(repomocks.MockDishRepository)
		mockPhotoUseCase := new(MockCreatePhoto) // Mock para CreatePhoto

		photoURL := "https://www.google.com"
		photo := &entities.Photo{
			URL: photoURL,
		}

		// Configure o mock para retornar a foto quando a URL for passada
		mockPhotoUseCase.On("Create", photoURL).Return(photo, nil)

		du := NewDishUseCase(mockDishRepo, mockPhotoUseCase) // Inclua a dependência mock
		dishName := "Ramen"
		description := "Tasty ramen with vegetables"

		mockDishRepo.On("AddDish", mock.AnythingOfType("*entities.Dish")).Return(nil)

		dish, err := du.Create(dishName, description, photoURL) // Atualize para passar photoURL

		assert.NoError(t, err)
		assert.NotNil(t, dish)
		assert.Equal(t, dishName, dish.Name)
		assert.Equal(t, description, dish.Description)
		assert.Equal(t, photo, dish.Photo)
		mockDishRepo.AssertExpectations(t)
		mockPhotoUseCase.AssertExpectations(t) // Adicione isso

	})

	t.Run("should not be able to create a dish without a name", func(t *testing.T) {
		mockDishRepo := new(repomocks.MockDishRepository)
		mockPhotoUseCase := new(MockCreatePhoto) // Mock para CreatePhoto

		photoURL := "https://www.google.com"
		photo := &entities.Photo{
			URL: photoURL,
		}

		// Configure o mock para retornar a foto quando a URL for passada
		mockPhotoUseCase.On("Create", photoURL).Return(photo, nil)

		du := NewDishUseCase(mockDishRepo, mockPhotoUseCase)
		dishName := ""
		description := "Tasty ramen with vegetables"

		dish, err := du.Create(dishName, description, photoURL)

		assert.Error(t, err)
		assert.Nil(t, dish)
		assert.Equal(t, err.Error(), "name is required")
	})
}
