package usecases

import (
	"fmt"
	"time"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type orderCreatorImpl struct {
	orderRepo repositories.OrderRepository // Use a interface aqui
}

// Verifique se orderCreatorImpl implementa a interface OrderCreator
var _ usecases.OrderCreator = &orderCreatorImpl{}

func NewOrderCreator() usecases.OrderCreator { // Use a interface aqui
	return &orderCreatorImpl{}
}

func (oc *orderCreatorImpl) CreateOrder(client entities.Client, items []entities.OrderItem) (*entities.Order, error) {
	order := &entities.Order{
		Client:    client,
		Items:     items,
		CreatedAt: time.Now(),
	}

	err := oc.orderRepo.Save(order)

	fmt.Println("Error after saving:", err)

	if err != nil {
		fmt.Println("Returning due to error.")
		return nil, err
	}

	fmt.Println("Order created successfully.")
	return order, nil
}
