package controllers

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type CreateDishController interface {
	AddDish(name string, description string, photoURL string) (*entities.Dish, error)
}
