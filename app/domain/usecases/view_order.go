package usecases

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type viewOrdersImpl struct {
	orderRepo repositories.OrderRepository
}

var _ usecases.ViewOrders = &viewOrdersImpl{}

func NewViewOrders(repo repositories.OrderRepository) usecases.ViewOrders {
	return &viewOrdersImpl{orderRepo: repo}
}

func (vo *viewOrdersImpl) GetAllOrders() ([]entities.Order, error) {
	return vo.orderRepo.List()
}

func (vo *viewOrdersImpl) GetOrder(orderID int) (*entities.Order, error) {

	return vo.orderRepo.GetOrder(orderID)

}
