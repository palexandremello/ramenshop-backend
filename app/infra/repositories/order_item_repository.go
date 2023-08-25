package repositories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

type OrderItemSQLAdapter struct {
	DB *sql.DB
}

func NewOrderItemSQLRepository(database *sql.DB) repositories.OrderItemRepository {
	return &OrderItemSQLAdapter{DB: database}
}

func (oi *OrderItemSQLAdapter) Save(item *entities.OrderItem) error {

	stmt, err := oi.DB.Prepare("INSERT INTO order_items (order_id, dish_id, quantity) VALUES ($1, $2, $3)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.OrderID, item.Dish.ID, item.Amount)
	return err
}
