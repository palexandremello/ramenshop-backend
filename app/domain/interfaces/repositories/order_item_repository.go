package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// OrderItemRepository interface
type OrderItemRepository interface {
	Save(item *entities.OrderItem) error
	// Update(item *entities.OrderItem) error
	// Delete(id int) error
	// GetByID(id int) (*entities.OrderItem, error)
	// List() ([]*entities.OrderItem, error)
}
