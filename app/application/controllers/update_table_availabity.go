package controllers

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type UpdateTableAvailabityController struct {
	updateTableAvailabityUseCase usecases.UpdateTableAvailability
}

var _ controllers.UpdateTableAvailabityController = &UpdateTableAvailabityController{}

func NewUpdateTableAvailabityController(useCase usecases.UpdateTableAvailability) controllers.UpdateTableAvailabityController {
	return &UpdateTableAvailabityController{updateTableAvailabityUseCase: useCase}
}

func (uta *UpdateTableAvailabityController) Execute(tableID int, isAvailable bool) (*entities.Table, error) {

	updatedTable, err := uta.updateTableAvailabityUseCase.Execute(tableID, isAvailable)

	if err != nil {
		return nil, err
	}

	return updatedTable, nil

}
