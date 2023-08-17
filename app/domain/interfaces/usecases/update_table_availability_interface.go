package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type UpdateTableAvailability interface {
	Execute(tableID int, isAvailable bool) (*entities.Table, error)
}
