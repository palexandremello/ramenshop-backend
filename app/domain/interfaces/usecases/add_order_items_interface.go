package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// AddOrderItems
type AddOrderItems interface {
	Add(orderID int, items []entities.OrderItem) error
}
