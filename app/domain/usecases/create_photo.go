package usecases

import (
	"errors"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/services"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type createPhotoImpl struct {
	photoRepo    repositories.PhotoRepository
	httpService  services.HTTPService
	fileUploader services.FileUploader
}

var _ usecases.CreatePhoto = &createPhotoImpl{}

// NewPhotoUseCase factory
func NewPhotoUseCase(repo repositories.PhotoRepository, httpService services.HTTPService, fileUploader services.FileUploader) usecases.CreatePhoto {
	return &createPhotoImpl{photoRepo: repo, httpService: httpService,
		fileUploader: fileUploader}
}

func (pu *createPhotoImpl) Create(url string) (*entities.Photo, error) {
	if url == "" {
		return nil, errors.New("URL is required")
	}

	mimeType, err := pu.httpService.GetMimeTypeFromURL(url)

	if err != nil {
		return nil, err
	}

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return nil, errors.New("Only jpeg and png are supported")
	}

	photo := &entities.Photo{
		URL: url,
	}

	err = pu.photoRepo.Save(photo)

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (pu *createPhotoImpl) Upload(file []byte, fileName string) (*entities.Photo, error) {

	uploadedURL, err := pu.fileUploader.Upload(file, fileName)

	if err != nil {
		return nil, err
	}

	photo := &entities.Photo{
		URL: uploadedURL,
	}

	err = pu.photoRepo.Save(photo)

	if err != nil {
		return nil, err
	}

	return photo, nil
}
