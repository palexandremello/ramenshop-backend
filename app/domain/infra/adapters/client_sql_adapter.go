package adapters

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

type ClientSQLAdapter struct {
	DB *sql.DB
}

func NewClientSQLRepository(database *sql.DB) repositories.ClientRepository {
	return &ClientSQLAdapter{DB: database}
}

func (ca *ClientSQLAdapter) Save(client *entities.Client) error {

	stmt, err := ca.DB.Prepare("INSERT INTO clients (id, name, gender, age) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(client.ID, client.Name, client.Gender, client.Age)
	return err
}

func (ca *ClientSQLAdapter) List() ([]entities.Client, error) {

	rows, err := ca.DB.Query("SELECT id, name, gender, age FROM clients")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		err = rows.Scan(&client.ID, &client.Name, &client.Gender, &client.Age)

		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}
