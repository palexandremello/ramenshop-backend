package controllers

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type listDishController struct {
	listDishUseCase usecases.ListDish
}

var _ controllers.ListDishController = &listDishController{}

// NewListDishController creates a new instance of ListDishController
func NewListDishController(useCase usecases.ListDish) controllers.ListDishController {
	return &listDishController{listDishUseCase: useCase}
}

func (ld *listDishController) Execute(dishType string, dishID int) ([]*entities.Dish, error) {
	if dishType == "" && dishID == -1 {
		return ld.listDishUseCase.List()
	}

	if dishType != "" && dishID == -1 {
		return ld.listDishUseCase.ListDishesByType(dishType)
	}

	if dishID > 0 {
		dish, err := ld.listDishUseCase.ListDishesByID(dishID)

		if err != nil {
			return nil, err
		}

		if dish != nil {
			return []*entities.Dish{dish}, nil
		}

	}

	return nil, errors.New("Invalid filter parameters")

}
