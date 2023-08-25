package repositories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

type OrderTableAdapter struct {
	DB *sql.DB
}

func NewOrderTableSQLRepository(database *sql.DB) repositories.OrderTableRepositry {
	return &OrderTableAdapter{DB: database}
}

func (ot *OrderTableAdapter) CreateAssociation(orderID int, tableID int) error {

	stmt, err := ot.DB.Prepare("INSERT INTO order_table (order_id, table_id) VALUES ($1, $2)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(orderID, tableID)
	return err
}
