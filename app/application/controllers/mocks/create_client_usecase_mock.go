package controllermocks

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/stretchr/testify/mock"
)

type CreateClient struct {
	mock.Mock
}

func (m *CreateClient) Create(id int, name string, gender entities.Gender, age int) (*entities.Client, error) {
	args := m.Called(id, name, gender, age)
	return args.Get(0).(*entities.Client), args.Error(1)
}
