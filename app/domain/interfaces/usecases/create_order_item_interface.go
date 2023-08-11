package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// CreateOrderItem UseCase interface
type CreateOrderItem interface {
	Create(orderID int, dishID int, amount int) (*entities.OrderItem, error)
}
