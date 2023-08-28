package services

import (
	"encoding/json"
	"os"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/services"
)

// RedisNotifier is a struct that implements the Notifier interface
type RedisNotifier struct {
	publisher services.PublisherEvent
}

// NewRedisNotifier is a factory function that creates a new RedisNotifier
func NewRedisNotifier(pub services.PublisherEvent) services.Notifier {
	return &RedisNotifier{publisher: pub}
}

// NotifyNewOrder is a method that notifies a new order
func (rn *RedisNotifier) NotifyNewOrder(order *entities.Order) error {
	channel := os.Getenv("REDIS_PUBLISHER_CHANNEL")
	message := formatOrder(order)

	return rn.publisher.Execute(channel, message)
}

// NotifyOrderItem is a method that notifies a new order item
func (rn *RedisNotifier) NotifyOrderItem(order *entities.Order, orderItems []*entities.OrderItem) error {
	channel := os.Getenv("REDIS_PUBLISHER_CHANNEL")
	data, err := json.Marshal(orderItems)
	if err != nil {
		return err
	}

	return rn.publisher.Execute(channel, string(data))
}
