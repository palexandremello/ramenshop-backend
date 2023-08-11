package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// DishTypeRepository interface
type DishTypeRepository interface {
	GetDishType(typeID int) (*entities.DishType, error)
	ListDishTypes() ([]*entities.DishType, error)
	AddDishType(dishType *entities.DishType) error
	UpdateDishType(dishType *entities.DishType) error
	DeleteDishType(typeID int) error
}
