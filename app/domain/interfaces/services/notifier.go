package services

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// Notifier is the interface that wraps the basic NotifyNewOrder method.
type Notifier interface {
	NotifyNewOrder(order *entities.Order) error
	NotifyOrderItem(order *entities.Order, orderItems []*entities.OrderItem) error
}
