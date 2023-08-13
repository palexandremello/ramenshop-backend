package usecases

import (
	"errors"
	"time"

	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type closeOrderImpl struct {
	OrderRepo repositories.OrderRepository
}

var _ usecases.CloseOrder = &closeOrderImpl{}

func NewCloseOrder(repo repositories.OrderRepository) usecases.CloseOrder {
	return &closeOrderImpl{OrderRepo: repo}
}

func (co *closeOrderImpl) Execute(orderID int) error {
	order, err := co.OrderRepo.GetOrder(orderID)

	if err != nil {
		return err
	}

	if order.ClosedAt != nil {
		return errors.New("order was closed already")
	}
	now := time.Now()
	order.ClosedAt = &now
	return co.OrderRepo.Update(order)
}
