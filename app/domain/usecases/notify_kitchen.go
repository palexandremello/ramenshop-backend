package usecases

import (
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/services"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type notifyKitchenUseCaseImpl struct {
	notifier services.Notifier
}

var _ usecases.NotifyKitchen = &notifyKitchenUseCaseImpl{}

// NewNotifyKitchenUseCase creates a new instance of NotifyKitchen
func NewNotifyKitchenUseCase(n services.Notifier) usecases.NotifyKitchen {
	return &notifyKitchenUseCaseImpl{notifier: n}
}

func (nk *notifyKitchenUseCaseImpl) Execute(order *entities.Order) error {
	return nk.notifier.NotifyNewOrder(order)
}
