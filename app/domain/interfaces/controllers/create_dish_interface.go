package controllers

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// CreateDishController is a interface that defines the methods of CreateDishController
type CreateDishController interface {
	AddDish(name string, description string, photoURL string) (*entities.Dish, error)
	Execute(name string, description string, file []byte, fileName string, price float64, dishType string) (*entities.Dish, error)
}
