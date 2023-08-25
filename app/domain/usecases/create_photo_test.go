package usecases

import (
	"errors"
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileUploader struct {
	mock.Mock
}

func (m *MockFileUploader) Upload(file []byte, fileName string) (string, error) {
	args := m.Called(file, fileName)
	return args.String(0), args.Error(1)
}

type MockPhotoRepository struct {
	mock.Mock
}

func (m *MockPhotoRepository) Save(photo *entities.Photo) error {
	args := m.Called(photo)
	return args.Error(0)
}

type MockHTTPService struct {
	mock.Mock
}

func (m *MockHTTPService) GetMimeTypeFromURL(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

type testResources struct {
	mockRepo        *MockPhotoRepository
	mockHTTPService *MockHTTPService
	uc              usecases.CreatePhoto
}

func setUp() *testResources {
	mockRepo := new(MockPhotoRepository)
	mockHTTPService := new(MockHTTPService)
	mockFileUploader := new(MockFileUploader)
	uc := NewPhotoUseCase(mockRepo, mockHTTPService, mockFileUploader)

	return &testResources{
		mockRepo:        mockRepo,
		mockHTTPService: mockHTTPService,
		uc:              uc,
	}
}

func TestCreatePhoto(t *testing.T) {
	resources := setUp()

	t.Run("should be able to create a new photo with valid URL", func(t *testing.T) {
		url := "https://upload.wikimedia.org/wikipedia/commons/e/ec/Shoyu_ramen%2C_at_Kasukabe_Station_%282014.05.05%29_1.jpg"
		resources.mockRepo.On("Save", mock.AnythingOfType("*entities.Photo")).Return(nil)
		resources.mockHTTPService.On("GetMimeTypeFromURL", url).Return("image/jpeg", nil)

		photo, err := resources.uc.Create(url)

		assert.NoError(t, err)
		assert.NotNil(t, photo)
		assert.Equal(t, url, photo.URL)
		resources.mockRepo.AssertExpectations(t)
		resources.mockHTTPService.AssertExpectations(t)
	})

	t.Run("should return an error for invalid MIME type", func(t *testing.T) {
		url := "https://example.com/image.gif"

		resources.mockHTTPService.On("GetMimeTypeFromURL", url).Return("image/gif", nil)

		photo, err := resources.uc.Create(url)

		assert.Error(t, err)
		assert.Nil(t, photo)
		assert.Equal(t, err, errors.New("Only jpeg and png are supported"))
		resources.mockHTTPService.AssertExpectations(t)

	})
	t.Run("should return an error for empty URL", func(t *testing.T) {
		url := ""

		photo, err := resources.uc.Create(url)

		assert.Error(t, err)
		assert.Nil(t, photo)
		assert.Equal(t, err, errors.New("URL is required"))
	})

	t.Run("should return an error if GetMimeTypeFromURL returns an error", func(t *testing.T) {

		url := "https://example.com/image.jpg"
		resources.mockHTTPService.On("GetMimeTypeFromURL", url).Return("", errors.New("HTTP error"))

		photo, err := resources.uc.Create(url)

		assert.Error(t, err)
		assert.Nil(t, photo)
		assert.Equal(t, err.Error(), "HTTP error")
		resources.mockHTTPService.AssertExpectations(t)

	})

	t.Run("should return an error if Save method fails", func(t *testing.T) {
		resources := setUp()

		url := "https://example.com/image.jpg"
		resources.mockHTTPService.On("GetMimeTypeFromURL", url).Return("image/jpeg", nil)
		resources.mockRepo.On("Save", mock.AnythingOfType("*entities.Photo")).Return(errors.New("DB save error"))

		photo, err := resources.uc.Create(url)

		assert.Error(t, err)
		assert.Nil(t, photo)
		assert.Equal(t, err.Error(), "DB save error")
		resources.mockRepo.AssertExpectations(t)
		resources.mockHTTPService.AssertExpectations(t)

	})

}
