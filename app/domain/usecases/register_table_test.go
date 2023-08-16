package usecases

import (
	"errors"
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	repomocks "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterTableUseCase(t *testing.T) {

	mockRepo := new(repomocks.MockTableRepository)
	useCase := NewRegisterTable(mockRepo)
	mockTable := &entities.Table{
		Capacity:    5,
		IsAvailable: true,
	}

	t.Run("should register a table successfully", func(t *testing.T) {

		mockRepo.On("Add", mockTable).Return(mockTable, nil)

		table, err := useCase.Execute(5)

		assert.NoError(t, err)
		assert.NotNil(t, table)
		assert.Equal(t, 5, table.Capacity)
	})

	t.Run("should return an error if capacity is non-positive", func(t *testing.T) {
		_, err := useCase.Execute(-1)

		assert.Error(t, err)
		assert.Equal(t, "table capacity should be more than zero", err.Error())
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo := new(repomocks.MockTableRepository)
		useCase := NewRegisterTable(mockRepo)

		mockRepo.On("Add", mock.Anything).Return(nil, errors.New("database error"))

		_, err := useCase.Execute(5)

		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})
}
