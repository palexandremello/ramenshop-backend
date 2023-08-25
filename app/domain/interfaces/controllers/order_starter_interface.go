package controllers

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type OrderStarterController interface {
	Execute(customerName *string, tableID int, dishInputs []usecases.DishOrderInput) (*entities.Order, error)
}
