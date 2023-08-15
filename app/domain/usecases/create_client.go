package usecases

import (
	"errors"
	"strings"
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type createClientImpl struct {
	clientRepo repositories.ClientRepository // Use a interface aqui
}

var _ usecases.CreateClient = &createClientImpl{}

// NewCreateClient creates a new instance of CreateClient
func NewCreateClient(repo repositories.ClientRepository) usecases.CreateClient {
	return &createClientImpl{
		clientRepo: repo,
	}
}

func (cc *createClientImpl) Create(name string, gender entities.Gender, age int) (*entities.Client, error) {

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name is required")
	}

	if len(name) < 5 {
		return nil, errors.New("name must have at least 5 characters")
	}

	if gender != entities.Male && gender != entities.Female {
		return nil, errors.New("for now we have only two genders, sorry. We Will fix it soon")
	}

	if age < 0 || age > 120 {
		return nil, errors.New("age must be between 0 and 120")
	}

	client := &entities.Client{
		Name:   name,
		Gender: gender,
		Age:    age,
	}

	err := cc.clientRepo.Save(client)
	if err != nil {
		return nil, err
	}

	return client, nil

}
