package controllers

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type UpdateTableAvailabityController interface {
	Execute(tableID int, isAvailable bool) (*entities.Table, error)
}
