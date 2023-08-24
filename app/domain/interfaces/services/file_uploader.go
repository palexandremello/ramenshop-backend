package services

// FileUploader is an interface that defines the methods for uploading files
type FileUploader interface {
	Upload(file []byte, fileName string) (string, error)
}
