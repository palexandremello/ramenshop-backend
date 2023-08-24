package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// ListDish UseCase interface
type ListDish interface {
	List() ([]*entities.Dish, error)
}
