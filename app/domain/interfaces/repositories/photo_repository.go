package repositories

import "github.com/palexandremello/ramenshop-backend/app/domain/entities"

// PhotoRepository Interface
type PhotoRepository interface {
	Save(photo *entities.Photo) error
	// Update(photo *entities.Photo) error
	// Delete(url string) error
}
