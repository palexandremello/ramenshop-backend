package services

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/services"
)

// DigitalOceanSpacesFileUploader is a struct that implements the FileUploader interface
type DigitalOceanSpacesFileUploader struct{}

// NewDigitalOceanSpacesFileUploader is a factory function that creates a new DigitalOceanSpacesFileUploader
func NewDigitalOceanSpacesFileUploader() services.FileUploader {
	return &DigitalOceanSpacesFileUploader{}
}

// Upload is a method that uploads a file to DigitalOcean Spaces
func (dsfu *DigitalOceanSpacesFileUploader) Upload(file []byte, fileName string) (string, error) {
	reader := bytes.NewReader(file)

	key := os.Getenv("SPACES_KEY")
	secret := os.Getenv("SPACES_SECRET")

	if key == "" || secret == "" {
		return "", errors.New("SPACES_KEY or SPACES_SECRET are empty or not set")
	}

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String("https://nyc3.digitaloceanspaces.com"),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(false),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	path := "files/" + fileName
	object := s3.PutObjectInput{
		Bucket: aws.String("ramenshop-bucket"),
		Key:    aws.String(path),
		Body:   reader,
		ACL:    aws.String("private"),
	}

	_, err := s3Client.PutObject(&object)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.%s/%s", "ramenshop-bucket", os.Getenv("SPACES_ENDPOINT"), path)

	return url, nil
}
