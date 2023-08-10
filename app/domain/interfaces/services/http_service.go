package services

// HTTPService interface
type HTTPService interface {
	GetMimeTypeFromURL(url string) (string, error)
}
