package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type createOrderItemImpl struct {
	dishRepo  repositories.DishRepository
	orderRepo repositories.OrderRepository
}

var _ usecases.CreateOrderItem = &createOrderItemImpl{}

// NewCreateOrderItemUseCase factory
func NewCreateOrderItemUseCase(dishRepo repositories.DishRepository, orderRepo repositories.OrderRepository) usecases.CreateOrderItem {
	return &createOrderItemImpl{dishRepo: dishRepo, orderRepo: orderRepo}
}

func (ou *createOrderItemImpl) Create(orderID int, dishID int, amount int) (*entities.OrderItem, error) {

	if amount <= 0 {
		return nil, errors.New("amount should be greater than 0")
	}

	order, err := ou.orderRepo.GetOrder(orderID)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, errors.New("Order not found")
	}

	dish, err := ou.dishRepo.GetDish(dishID)

	if err != nil {
		return nil, err
	}

	if dish == nil {
		return nil, errors.New("Dish not found")
	}

	orderItem := &entities.OrderItem{
		OrderID: orderID,
		Dish:    *dish,
		Amount:  amount,
	}

	err = ou.orderRepo.AddOrderItem(orderItem)

	if err != nil {
		return nil, err
	}

	return orderItem, nil
}
