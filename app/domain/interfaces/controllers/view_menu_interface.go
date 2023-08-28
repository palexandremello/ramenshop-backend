package controllers

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// ViewMenuController is an interface that defines the Controller for the ViewMenu use case
type ViewMenuController interface {
	Execute() ([]*entities.Dish, error)
}
