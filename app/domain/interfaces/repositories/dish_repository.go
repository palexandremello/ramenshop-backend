package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// DishRepository interface
type DishRepository interface {
	GetDish(dishID int) (*entities.Dish, error)
	ListDishesByType(dishType string) ([]*entities.Dish, error)
	ListAllDishes() ([]*entities.Dish, error)
	AddDish(dish *entities.Dish) error
	UpdateDish(dish *entities.Dish) error
	DeleteDish(dishID int) error
}
