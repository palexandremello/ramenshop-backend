package usecases

import (
	"errors"
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTableAvailability(t *testing.T) {

	t.Run("should update a table availability successfully", func(t *testing.T) {
		mockRepo := new(repomocks.MockTableRepository)
		useCase := NewUpdateTableAvailabilty(mockRepo)

		mockTable := &entities.Table{
			ID:          1,
			IsAvailable: false,
		}

		mockRepo.On("FindByID", 1).Return(mockTable, nil)
		mockRepo.On("Update", mockTable).Return(nil)

		_, err := useCase.Execute(1, true)

		assert.NoError(t, err)
		assert.True(t, mockTable.IsAvailable)
	})

	t.Run("should return an error if table does not exist", func(t *testing.T) {
		mockRepo := new(repomocks.MockTableRepository)
		useCase := NewUpdateTableAvailabilty(mockRepo)

		mockRepo.On("FindByID", 2).Return(nil, errors.New("table not found"))

		_, err := useCase.Execute(2, true)

		assert.Error(t, err)
		assert.Equal(t, "table does not exists", err.Error())
	})

	t.Run("should return an error if update fails", func(t *testing.T) {
		mockRepo := new(repomocks.MockTableRepository)
		useCase := NewUpdateTableAvailabilty(mockRepo)

		mockTable := &entities.Table{
			ID:          3,
			IsAvailable: false,
		}

		mockRepo.On("FindByID", 3).Return(mockTable, nil)
		mockRepo.On("Update", mockTable).Return(errors.New("database error"))

		_, err := useCase.Execute(3, true)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())

	})
}
