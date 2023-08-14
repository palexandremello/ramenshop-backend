package ports

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type CreateClientHandler interface {
	Handle(client *entities.Client) (*entities.Client, error)
}
