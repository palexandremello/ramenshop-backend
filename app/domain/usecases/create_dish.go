package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type dishUseCaseImpl struct {
	dishRepo           repositories.DishRepository
	createPhotoUseCase usecases.CreatePhoto
}

var _ usecases.CreateDish = &dishUseCaseImpl{}

// NewDishUseCase creates a new instance of CreateDish
func NewDishUseCase(repo repositories.DishRepository, useCase usecases.CreatePhoto) usecases.CreateDish {
	return &dishUseCaseImpl{dishRepo: repo, createPhotoUseCase: useCase}
}

func (du *dishUseCaseImpl) Create(name string, description string, photoURL string) (*entities.Dish, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	photo, err := du.createPhotoUseCase.Create(photoURL)

	if err != nil {
		return nil, err
	}

	dish := &entities.Dish{
		Name:        name,
		Description: description,
		Photo:       photo,
		Available:   true,
	}

	err = du.dishRepo.AddDish(dish)

	if err != nil {
		return nil, err
	}

	return dish, nil
}

func (du *dishUseCaseImpl) Execute(name string, description string, file []byte, fileName string, price float64, dishType string) (*entities.Dish, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	photo, err := du.createPhotoUseCase.Upload(file, fileName)

	if err != nil {
		return nil, err
	}

	dish := &entities.Dish{
		Name:        name,
		Description: description,
		Photo:       photo,
		Available:   true,
		Price:       price,
		Type:        dishType,
	}

	err = du.dishRepo.AddDish(dish)

	if err != nil {
		return nil, err
	}

	return dish, nil

}
