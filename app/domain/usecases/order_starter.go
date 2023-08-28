package usecases

import (
	"time"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type orderStarterImpl struct {
	orderRepo      repositories.OrderRepository
	orderItemRepo  repositories.OrderItemRepository
	dishRepo       repositories.DishRepository
	orderTableRepo repositories.OrderTableRepositry
}

var _ usecases.OrderStarter = &orderStarterImpl{}

// NewOrderStarterUseCase creates a new instance of OrderStarter
func NewOrderStarterUseCase(repo repositories.OrderRepository,
	itemRepo repositories.OrderItemRepository,
	repoDish repositories.DishRepository,
	orderTableRepo repositories.OrderTableRepositry) usecases.OrderStarter {
	return &orderStarterImpl{orderRepo: repo, orderItemRepo: itemRepo, dishRepo: repoDish, orderTableRepo: orderTableRepo}
}

func (os *orderStarterImpl) StartOrder(customerName *string, tableID int, dishInputs []usecases.DishOrderInput) (*entities.Order, error) {
	order := &entities.Order{
		CustomerName: customerName,
		TableID:      tableID,
		Status:       entities.InProcess,
		CreatedAt:    time.Now(),
	}

	err := os.orderRepo.Save(order)

	if err != nil {
		return nil, err
	}

	err = os.orderTableRepo.CreateAssociation(order.ID, tableID)

	if err != nil {
		return nil, err
	}

	for _, input := range dishInputs {
		dish, err := os.dishRepo.GetDish(input.DishID)

		if err != nil {
			return nil, err
		}

		orderItem := &entities.OrderItem{
			OrderID: order.ID,
			Dish:    *dish,
			Amount:  input.Amount,
		}

		err = os.orderItemRepo.Save(orderItem)

		if err != nil {
			return nil, err
		}

		order.Items = append(order.Items, *orderItem)
	}
	return order, nil
}
