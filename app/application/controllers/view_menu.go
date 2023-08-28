package controllers

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type ViewMenuController struct {
	viewMenuUseCase usecases.ViewMenu
}

var _ controllers.ViewMenuController = &ViewMenuController{}

func NewViewMenuController(useCase usecases.ViewMenu) controllers.ViewMenuController {
	return &ViewMenuController{viewMenuUseCase: useCase}
}

func (vm *ViewMenuController) Execute() ([]*entities.Dish, error) {
	return vm.viewMenuUseCase.Execute()
}
