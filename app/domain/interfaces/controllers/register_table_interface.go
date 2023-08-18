package controllers

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

type RegisterTableController interface {
	Execute(capacity int) (*entities.Table, error)
}
