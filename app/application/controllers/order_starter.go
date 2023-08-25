package controllers

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type orderStarterController struct {
	orderStarterUseCase usecases.OrderStarter
}

var _ controllers.OrderStarterController = &orderStarterController{}

// NewOrderStarterController creates a new instance of OrderStarterController
func NewOrderStarterController(useCase usecases.OrderStarter) controllers.OrderStarterController {
	return &orderStarterController{orderStarterUseCase: useCase}
}

func (osc *orderStarterController) Execute(customerName *string, tableID int, dishInputs []usecases.DishOrderInput) (*entities.Order, error) {

	if tableID < 0 {
		return nil, errors.New("tableID should be greater than 0")
	}

	return osc.orderStarterUseCase.StartOrder(customerName, tableID, dishInputs)
}
