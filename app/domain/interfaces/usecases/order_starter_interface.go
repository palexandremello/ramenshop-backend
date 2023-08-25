package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// DishOrderInput is the input to start an order
type DishOrderInput struct {
	DishID int
	Amount int
}

// OrderStarter is the interface that wraps the StartOrder method.
type OrderStarter interface {
	StartOrder(customerName *string, tableID int, dishInputs []DishOrderInput) (*entities.Order, error)
}
