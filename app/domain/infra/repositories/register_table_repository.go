package repositories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

type RegisterSQLAdapter struct {
	DB *sql.DB
}

func NewRegisterSQLRepository(database *sql.DB) repositories.TableRepository {
	return &RegisterSQLAdapter{DB: database}
}

func (rs *RegisterSQLAdapter) Add(table *entities.Table) (*entities.Table, error) {

	stmt, err := rs.DB.Prepare("INSERT INTO tables (capacity, is_available) VALUES ($1, $2) RETURNING id")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var tableID int
	err = stmt.QueryRow(table.Capacity, table.IsAvailable).Scan(&tableID)
	if err != nil {
		return nil, err
	}

	table.ID = tableID
	return table, nil
}

func (rs *RegisterSQLAdapter) FindByID(id int) (*entities.Table, error) {
	row := rs.DB.QueryRow("SELECT id, capacity, is_available FROM tables WHERE id = $1", id)

	var table entities.Table
	err := row.Scan(&table.ID, &table.Capacity, &table.IsAvailable)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &table, nil
}

func (rs *RegisterSQLAdapter) Update(table *entities.Table) error {
	stmt, err := rs.DB.Prepare("UPDATE tables SET capacity = $1, is_available = $2 WHERE id = $3")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(table.Capacity, table.IsAvailable, table.ID)
	return err
}

func (rs *RegisterSQLAdapter) Remove(id int) error {
	stmt, err := rs.DB.Prepare("DELETE FROM tables WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (rs *RegisterSQLAdapter) List() ([]*entities.Table, error) {
	rows, err := rs.DB.Query("SELECT id, capacity, is_available FROM tables")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tables []*entities.Table
	for rows.Next() {
		var table entities.Table
		err = rows.Scan(&table.ID, &table.Capacity, &table.IsAvailable)

		if err != nil {
			return nil, err
		}

		tables = append(tables, &table)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}
