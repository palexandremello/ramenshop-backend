package controllers

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type RegisterTableController struct {
	registerTableUseCase usecases.RegisterTable
}

var _ controllers.RegisterTableController = &RegisterTableController{}

func NewRegisterTableController(useCase usecases.RegisterTable) controllers.RegisterTableController {
	return &RegisterTableController{registerTableUseCase: useCase}
}

func (rtc *RegisterTableController) Execute(capacity int) (*entities.Table, error) {
	registeredTable, err := rtc.registerTableUseCase.Execute(capacity)

	if err != nil {
		return nil, err
	}

	return registeredTable, nil
}
