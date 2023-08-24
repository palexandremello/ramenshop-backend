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

// NewCreateDishController creates a new instance of CreateDishController
func NewCreateDishController(useCase usecases.CreateDish) controllers.CreateDishController {
	return &createDishController{createDishUseCase: useCase}
}

func (cdc *createDishController) AddDish(name string, description string, photoURL string) (*entities.Dish, error) {
	return cdc.createDishUseCase.Create(name, description, photoURL)
}

func (cdc *createDishController) Execute(name string, description string, file []byte, fileName string, price float64, dishType string) (*entities.Dish, error) {
	return cdc.createDishUseCase.Execute(name, description, file, fileName, price, dishType)
}
