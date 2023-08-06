package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// OrderCreator UseCase Interface
type OrderCreator interface {
	CreateOrder(client entities.Client, items []entities.OrderItem) (*entities.Order, error)
}
