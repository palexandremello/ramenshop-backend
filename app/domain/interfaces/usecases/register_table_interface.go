package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type RegisterTable interface {
	Execute(capacity int) (*entities.Table, error)
}
