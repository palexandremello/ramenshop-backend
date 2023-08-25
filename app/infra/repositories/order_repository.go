package repositories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

type OrderSQLAdapter struct {
	DB *sql.DB
}

func NewOrderSQLRepository(database *sql.DB) repositories.OrderRepository {
	return &OrderSQLAdapter{DB: database}
}

func (oa *OrderSQLAdapter) Save(order *entities.Order) error {
	stmt, err := oa.DB.Prepare("INSERT INTO orders (customer_name, phone_number, table_id) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(order.CustomerName, order.PhoneNumber, order.TableID).Scan(&order.ID)
	return err
}

func (oa *OrderSQLAdapter) AddOrderItem(orderItem *entities.OrderItem) error {
	stmt, err := oa.DB.Prepare("INSERT INTO order_items (order_id, dish_id, amount) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(orderItem.OrderID, orderItem.Dish.ID, orderItem.Amount)
	return err
}

func (oa *OrderSQLAdapter) List() ([]entities.Order, error) {
	rows, err := oa.DB.Query("SELECT id, customer_name, phone_number, table_id, created_at, closed_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		err := rows.Scan(&order.ID, &order.CustomerName, &order.PhoneNumber, &order.TableID, &order.CreatedAt, &order.ClosedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// GetOrder retrieves a single order by ID from the database
func (oa *OrderSQLAdapter) GetOrder(orderID int) (*entities.Order, error) {
	row := oa.DB.QueryRow("SELECT id, customer_name, phone_number, table_id, created_at, closed_at FROM orders WHERE id = $1", orderID)

	var order entities.Order
	err := row.Scan(&order.ID, &order.CustomerName, &order.PhoneNumber, &order.TableID, &order.CreatedAt, &order.ClosedAt)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// Update updates the details of an existing order
func (oa *OrderSQLAdapter) Update(order *entities.Order) error {
	stmt, err := oa.DB.Prepare("UPDATE orders SET customer_name = $1, phone_number = $2, table_id = $3, closed_at = $4 WHERE id = $5")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.CustomerName, order.PhoneNumber, order.TableID, order.ClosedAt, order.ID)
	return err
}
