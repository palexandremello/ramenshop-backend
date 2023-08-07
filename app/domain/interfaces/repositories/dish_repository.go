package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type DishRepository interface {
	Save(dish *entities.Dish) error
	Update(dish *entities.Dish) error
	Delete(id int) error
	List() ([]entities.Dish, error)
	GetByID(id int) (*entities.Dish, error)
}
