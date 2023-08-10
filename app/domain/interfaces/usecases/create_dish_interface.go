package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// CreateDish UseCase interface
type CreateDish interface {
	Create(name string, description string, photo *entities.Photo) (*entities.Dish, error)
}
