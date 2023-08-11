package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// OrderRepository interface
type OrderRepository interface {
	Save(order *entities.Order) error
	AddOrderItem(orderItem *entities.OrderItem) error
	List() ([]entities.Order, error)
	GetOrder(orderID int) (*entities.Order, error)
}
