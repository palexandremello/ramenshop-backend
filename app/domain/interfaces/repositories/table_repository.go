package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// TableRepository interface
type TableRepository interface {
	Add(table *entities.Table) (*entities.Table, error)
	FindByID(id int) (*entities.Table, error)
	Update(table *entities.Table) error
	Remove(id int) error
	List() ([]*entities.Table, error)
}
