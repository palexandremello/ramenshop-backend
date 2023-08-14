package ports

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type ClientPort interface {
	CreateClient(client *entities.Client) (*entities.Client, error)
}
