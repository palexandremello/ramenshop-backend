package usecases

import (
	"time"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type orderCreatorImpl struct {
	orderRepo repositories.OrderRepository // Use a interface aqui
}

var _ usecases.OrderCreator = &orderCreatorImpl{}

// NewOrderCreator usecase Factory
func NewOrderCreator(repo repositories.OrderRepository) usecases.OrderCreator { // Use a interface aqui
	return &orderCreatorImpl{orderRepo: repo}
}

func (oc *orderCreatorImpl) CreateOrder(client entities.Client, items []entities.OrderItem) (*entities.Order, error) {
	order := &entities.Order{
		Client:    client,
		Items:     items,
		CreatedAt: time.Now(),
	}

	err := oc.orderRepo.Save(order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
