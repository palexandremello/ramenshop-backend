package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type ViewOrders interface {
	GetAllOrders() ([]entities.Order, error)
	GetOrder(orderID int) (*entities.Order, error)
}
