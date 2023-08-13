package controllers

import (
	"errors"
	"testing"

	controllermocks "github.com/palexandremello/ramenshop-backend/app/application/controllers/mocks"
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateClientController(t *testing.T) {
	t.Run("should create a client with successfull", func(t *testing.T) {
		mockUseCase := new(controllermocks.CreateClient)
		controller := NewCreateClientController(mockUseCase)
		client := &entities.Client{
			Name:   "Jeff Buckley",
			Gender: entities.Male,
			Age:    30,
		}

		mockUseCase.On("Create", client.ID, client.Name, client.Gender, client.Age).Return(client, nil)

		result, err := controller.Create(client)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, client.Name, result.Name)
	})

	t.Run("should return an error if CreateClient usecase fails", func(t *testing.T) {
		mockUseCase := new(controllermocks.CreateClient)
		controller := NewCreateClientController(mockUseCase)
		client := &entities.Client{
			Name:   "Jeff Buckley",
			Gender: entities.Male,
			Age:    30,
		}

		mockErr := errors.New("any_error")

		mockUseCase.On("Create", client.ID, client.Name, client.Gender, client.Age).Return(nil, mockErr)

		_, err := controller.Create(client)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "any_error")
	})
}
