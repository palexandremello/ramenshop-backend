package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type OrderRepository interface {
	Save(order *entities.Order) error
	List() ([]entities.Order, error)
}
