package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type dishUseCaseImpl struct {
	dishRepo repositories.DishRepository
}

var _ usecases.CreateDish = &dishUseCaseImpl{}

// NewDishUseCase creates a new instance of CreateDish
func NewDishUseCase(repo repositories.DishRepository) usecases.CreateDish {
	return &dishUseCaseImpl{dishRepo: repo}
}

func (du *dishUseCaseImpl) Create(name string, description string, photo *entities.Photo) (*entities.Dish, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	dish := &entities.Dish{
		Name:        name,
		Description: description,
		Photo:       photo,
	}

	err := du.dishRepo.AddDish(dish)

	if err != nil {
		return nil, err
	}

	return dish, nil
}
