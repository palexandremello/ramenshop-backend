package controllers

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type CloseOrderController struct {
	closeOrderUseCase usecases.CloseOrder
}

var _ controllers.CloseOrderController = &CloseOrderController{}

// NewCloseOrderController factory creates a new CloseOrderController
func NewCloseOrderController(useCase usecases.CloseOrder) controllers.CloseOrderController {
	return &CloseOrderController{closeOrderUseCase: useCase}
}

// Execute method the CloseOrder use case
func (coc *CloseOrderController) Execute(orderID int) error {

	err := coc.closeOrderUseCase.Execute(orderID)

	if err != nil {
		return err
	}

	return nil
}
