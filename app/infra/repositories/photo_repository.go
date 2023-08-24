package repositories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
)

// PhotoSQLAdapter is a struct that implements the PhotoRepository interface
type PhotoSQLAdapter struct {
	DB *sql.DB
}

// NewPhotoSQLRepository is a factory function that creates a new PhotoSQLAdapter
func NewPhotoSQLRepository(database *sql.DB) repositories.PhotoRepository {
	return &PhotoSQLAdapter{DB: database}
}

// Save is a method that saves a new photo to the database
func (pa *PhotoSQLAdapter) Save(photo *entities.Photo) error {
	stmt, err := pa.DB.Prepare("INSERT INTO photos (url) VALUES ($1) RETURNING id")

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(photo.URL).Scan(&photo.ID)
	return err
}
