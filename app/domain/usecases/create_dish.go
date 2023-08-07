package usecases

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

type dishUseCaseImpl struct {
	dishRepo repositories.DishRepository
}

func NewDishUsecase(repo repositories.DishRepository) *dishUseCaseImpl {
	return &dishUseCaseImpl{dishRepo: repo}
}
