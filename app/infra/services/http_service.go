package services

import (
	"net/http"
	"strings"

	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/services"
)

// HTTPServiceImpl is a struct that implements the HTTPService interface
type HTTPServiceImpl struct{}

// NewHTTPService is a factory function that creates a new HTTPServiceImpl
func NewHTTPService() services.HTTPService {
	return &HTTPServiceImpl{}
}

// GetMimeTypeFromURL is a method that returns the mime type of a given url
func (h *HTTPServiceImpl) GetMimeTypeFromURL(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", err
	}

	mimeType := strings.Split(contentType, ";")[0]

	return mimeType, nil
}
