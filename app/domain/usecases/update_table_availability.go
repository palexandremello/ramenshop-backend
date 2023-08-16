package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type updateTableAvailabiltyImpl struct {
	tableRepo repositories.TableRepository
}

var _ usecases.UpdateTableAvailability = &updateTableAvailabiltyImpl{}

func NewUpdateTableAvailabilty(repo repositories.TableRepository) usecases.UpdateTableAvailability {
	return &updateTableAvailabiltyImpl{
		tableRepo: repo,
	}
}

func (uta *updateTableAvailabiltyImpl) Execute(tableID int, isAvailable bool) error {

	table, err := uta.tableRepo.FindByID(tableID)

	if table == nil || err != nil {
		return errors.New("table does not exists")
	}

	table.IsAvailable = isAvailable

	err = uta.tableRepo.Update(table)

	if err != nil {
		return err
	}

	return nil
}
