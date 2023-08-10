package usecases

import (
	"testing"

	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
	uc := NewPhotoUseCase(mockRepo, mockHTTPService)

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
}
