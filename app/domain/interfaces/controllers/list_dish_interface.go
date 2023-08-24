package controllers

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// ListDishController is a interface that defines the methods of ListDishController
type ListDishController interface {
	Execute(dishType string, dishID int) ([]*entities.Dish, error)
}
