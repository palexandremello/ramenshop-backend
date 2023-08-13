package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type addOrderItemsImpl struct {
	orderRepo repositories.OrderRepository // Use a interface aqui
}

var _ usecases.AddOrderItems = &addOrderItemsImpl{}

func NewAddOrderItems(repo repositories.OrderRepository) usecases.AddOrderItems {
	return &addOrderItemsImpl{orderRepo: repo}
}

func (aoi *addOrderItemsImpl) Add(orderID int, items []entities.OrderItem) error {
	order, err := aoi.orderRepo.GetOrder(orderID)

	if err != nil {
		return err
	}

	if order.ClosedAt != nil {
		return errors.New("cannot add items to a closed order")
	}

	order.Items = append(order.Items, items...)
	return aoi.orderRepo.Update(order)
}
