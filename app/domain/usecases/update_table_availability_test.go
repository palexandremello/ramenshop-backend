package usecases

import (
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

		err := useCase.Execute(1, true)

		assert.NoError(t, err)
		assert.True(t, mockTable.IsAvailable)
	})
}
