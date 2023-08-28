package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type NotifyKitchen interface {
	Execute(order *entities.Order) error
}
