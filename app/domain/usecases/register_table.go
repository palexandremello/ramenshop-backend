package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type registerTableImpl struct {
	tableRepo repositories.TableRepository
}

var _ usecases.RegisterTable = &registerTableImpl{}

func NewRegisterTable(repo repositories.TableRepository) usecases.RegisterTable {
	return &registerTableImpl{tableRepo: repo}
}

func (rt *registerTableImpl) Execute(capacity int) (*entities.Table, error) {

	if capacity <= 0 {
		return nil, errors.New("table capacity should be more than zero")
	}

	table := &entities.Table{
		Capacity:    capacity,
		IsAvailable: true,
	}

	_, err := rt.tableRepo.Add(table)

	if err != nil {
		return nil, err
	}
	return table, nil
}
