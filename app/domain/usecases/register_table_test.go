package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRegisterTableUseCase(t *testing.T) {

	mockRepo := new(repomocks.MockTableRepository)
	useCase := NewRegisterTable(mockRepo)

	t.Run("should register a table successfully", func(t *testing.T) {

		mockTable := &entities.Table{
			Capacity:    5,
			IsAvailable: true,
		}

		mockRepo.On("Add", mockTable).Return(mockTable, nil)

		table, err := useCase.Execute(5)

		assert.NoError(t, err)
		assert.NotNil(t, table)
		assert.Equal(t, 5, table.Capacity)
	})
}
