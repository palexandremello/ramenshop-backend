package controllers

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type CreateClientController interface {
	Create(client *entities.Client) (*entities.Client, error)
}
