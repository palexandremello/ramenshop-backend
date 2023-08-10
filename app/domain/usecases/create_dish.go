package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

type dishUseCaseImpl struct {
	dishRepo repositories.DishRepository
}

func NewDishUseCase(repo repositories.DishRepository) *dishUseCaseImpl {
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

	err := du.dishRepo.Save(dish)

	if err != nil {
		return nil, err
	}

	return dish, nil
}
