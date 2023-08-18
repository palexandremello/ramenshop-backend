package controllers

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type createDishController struct {
	createDishUseCase usecases.CreateDish
}

var _ controllers.CreateDishController = &createDishController{}

func NewCreateDishController(useCase usecases.CreateDish) controllers.CreateDishController {
	return &createDishController{createDishUseCase: useCase}
}

func (cdc *createDishController) AddDish(name string, description string, photoURL string) (*entities.Dish, error) {
	return cdc.createDishUseCase.Create(name, description, photoURL)
}
