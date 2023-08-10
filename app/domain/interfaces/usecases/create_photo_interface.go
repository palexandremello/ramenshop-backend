package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// CreatePhoto UseCase interface
type CreatePhoto interface {
	Create(url string) (*entities.Photo, error)
}
