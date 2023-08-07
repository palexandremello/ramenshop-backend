package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type CreateDish interface {
	Create(name string, description string, photo *entities.Photo) (*entities.Dish, error)
}

type UpdateDish interface {
	UpdateDish(id int, name string, description string, photo *entities.Photo) (*entities.Dish, error)
}

type DeleteDish interface {
	Delete(id int) error
}

type ListDishes interface {
	List() ([]entities.Dish, error)
}

type GetDishByID interface {
	GetByID(id int) (*entities.Dish, error)
}
