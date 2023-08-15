package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type ViewMenu interface {
	Execute() ([]entities.Dish, error)
}
