package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientRepository struct {
	mock.Mock
}

func (m *MockClientRepository) Save(client *entities.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) List() ([]entities.Client, error) {

	args := m.Called()
	return args.Get(0).([]entities.Client), args.Error(1)
}

func TestCreateClient(t *testing.T) {

	t.Run("should be able to create a new client", func(t *testing.T) {

		mockRepo := new(MockClientRepository)
		cc := NewCreateClient(mockRepo)

		mockRepo.On("Save", mock.AnythingOfType("*entities.Client")).Return(nil)

		client, err := cc.Create(1, "Maicum", entities.Male, 29)

		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, "Maicum", client.Name)
		mockRepo.AssertExpectations(t)

	})

}
