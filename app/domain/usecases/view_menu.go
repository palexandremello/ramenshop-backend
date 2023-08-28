package usecases

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type viewMenuImpl struct {
	dishRepo repositories.DishRepository
}

var _ usecases.ViewMenu = &viewMenuImpl{}

// NewViewMenu is a factory function that creates a new viewMenuImpl
func NewViewMenu(repo repositories.DishRepository) usecases.ViewMenu {
	return &viewMenuImpl{dishRepo: repo}
}

func (vm *viewMenuImpl) Execute() ([]*entities.Dish, error) {
	return vm.dishRepo.ListAllDishes()
}
