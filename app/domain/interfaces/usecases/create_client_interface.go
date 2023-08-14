package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// CreateClient usecase interface
type CreateClient interface {
	Create(name string, gender entities.Gender, age int) (*entities.Client, error)
}
