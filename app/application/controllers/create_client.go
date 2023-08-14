package controllers

import (
	"fmt"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type CreateClientController struct {
	creteClientUseCase usecases.CreateClient
}

var _ controllers.CreateClientController = &CreateClientController{}

func NewCreateClientController(useCase usecases.CreateClient) controllers.CreateClientController {
	return &CreateClientController{creteClientUseCase: useCase}
}

func (cc *CreateClientController) Create(client *entities.Client) (*entities.Client, error) {
	createdClient, err := cc.creteClientUseCase.Create(client.Name, client.Gender, client.Age)

	if err != nil {
		return nil, fmt.Errorf("could not ceate a consumer %w", err)
	}
	return createdClient, nil
}
