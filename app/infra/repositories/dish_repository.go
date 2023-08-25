package repositories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

// DishSQLAdapter is a struct that implements the DishRepository interface
type DishSQLAdapter struct {
	DB *sql.DB
}

// NewDishSQLRepository is a factory function that creates a new DishSQLAdapter
func NewDishSQLRepository(database *sql.DB) repositories.DishRepository {
	return &DishSQLAdapter{DB: database}
}

// AddDish is a method that adds a dish to the database
func (da *DishSQLAdapter) AddDish(dish *entities.Dish) error {
	stmt, err := da.DB.Prepare("INSERT INTO dishes (name, description, photo_id, price, available, type) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(dish.Name, dish.Description, dish.Photo.ID, dish.Price, dish.Available, dish.Type).Scan(&dish.ID)
	return err
}

// ListAllDishes is a method that lists all dishes from the database
func (da *DishSQLAdapter) ListAllDishes() ([]*entities.Dish, error) {

	rows, err := da.DB.Query(`
    SELECT d.id, d.name, d.description, p.id, p.url, d.price, d.available, d.type 
    FROM dishes d 
    LEFT JOIN photos p ON d.photo_id = p.id`)

	if err != nil {
		return nil, err
	}

	var dishes []*entities.Dish
	for rows.Next() {
		dish := &entities.Dish{}
		var photoID int
		var photoURL string
		err = rows.Scan(&dish.ID, &dish.Name, &dish.Description, &photoID, &photoURL, &dish.Price, &dish.Available, &dish.Type)
		if err != nil {
			return nil, err
		}
		dish.Photo = &entities.Photo{ID: photoID, URL: photoURL}
		dishes = append(dishes, dish)
	}
	return dishes, rows.Err()

}

// GetDish is a method that gets a dish from the database
func (da *DishSQLAdapter) GetDish(dishID int) (*entities.Dish, error) {
	row := da.DB.QueryRow("SELECT id, name, description, photo_id, price, available, type, created_at FROM dishes WHERE id = $1", dishID)

	dish := &entities.Dish{}
	var photoID int

	err := row.Scan(&dish.ID, &dish.Name, &dish.Description, &photoID, &dish.Price, &dish.Available, &dish.Type)
	if err != nil {
		return nil, err
	}

	dish.Photo = &entities.Photo{ID: photoID}
	return dish, nil
}

// ListDishesByType is a method that lists dishes by type from the database
func (da *DishSQLAdapter) ListDishesByType(dishType string) ([]*entities.Dish, error) {
	rows, err := da.DB.Query("SELECT id, name, description, photo_id, price, available, type, created_at FROM dishes WHERE type = $1", dishType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []*entities.Dish
	for rows.Next() {
		dish := &entities.Dish{}
		var photoID int
		err = rows.Scan(&dish.ID, &dish.Name, &dish.Description, &photoID, &dish.Price, &dish.Available, &dish.Type)
		if err != nil {
			return nil, err
		}
		dish.Photo = &entities.Photo{ID: photoID}
		dishes = append(dishes, dish)
	}
	return dishes, rows.Err()
}

// UpdateDish is a method that updates a dish from the database
func (da *DishSQLAdapter) UpdateDish(dish *entities.Dish) error {
	stmt, err := da.DB.Prepare("UPDATE dishes SET name=$1, description=$2, photo_id=$3, price=$4, available=$5, type=$6 WHERE id=$7")

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(dish.Name, dish.Description, dish.Photo.ID, dish.Price, dish.Available, dish.Type, dish.ID)

	return err
}

// DeleteDish is a method that deletes a dish from the database
func (da *DishSQLAdapter) DeleteDish(dishID int) error {
	stmt, err := da.DB.Prepare("DELETE FROM dishes WHERE id=$1")

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(dishID)
	return err
}
