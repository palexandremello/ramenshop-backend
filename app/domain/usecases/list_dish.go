package usecases

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type listDishUseCaseImpl struct {
	dishRepo repositories.DishRepository
}

var _ usecases.ListDish = &listDishUseCaseImpl{}

// NewListDishUseCase creates a new instance of ListDish
func NewListDishUseCase(repo repositories.DishRepository) usecases.ListDish {
	return &listDishUseCaseImpl{dishRepo: repo}
}

func (ld *listDishUseCaseImpl) List() ([]*entities.Dish, error) {
	return ld.dishRepo.ListAllDishes()
}

func (ld *listDishUseCaseImpl) ListDishesByID(dishID int) (*entities.Dish, error) {

	return ld.dishRepo.GetDish(dishID)
}

func (ld *listDishUseCaseImpl) ListDishesByType(dishType string) ([]*entities.Dish, error) {
	return ld.dishRepo.ListDishesByType(dishType)
}
