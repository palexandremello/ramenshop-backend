package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type ClientRepository interface {
	Save(client *entities.Client) error
	List() ([]entities.Client, error)
}
