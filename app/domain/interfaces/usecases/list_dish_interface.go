package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// ListDish UseCase interface
type ListDish interface {
	List() ([]*entities.Dish, error)
	ListDishesByType(dishType string) ([]*entities.Dish, error)
	ListDishesByID(dishID int) (*entities.Dish, error)
}
