package usecases

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// CreateDish UseCase interface
type CreateDish interface {
	Create(name string, description string, photoURL string) (*entities.Dish, error)
	Execute(name string, description string, file []byte, fileName string, price float64, dishType string) (*entities.Dish, error)
}
