package entities

import (
	"errors"
	"strings"
)

type Gender int

const (
	Female Gender = iota
	Male
)

type Client struct {
	Name   string
	Gender Gender
	Age    int
}

// NewClient creates a new Client, validating that the provided name, gender and age are valid.
func NewClient(name string, gender Gender, age int) (*Client, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name is required")
	}

	if len(name) < 5 {
		return nil, errors.New("name must have at least 5 characters")
	}

	if gender != Male && gender != Female {
		return nil, errors.New("for now we have only two genders, sorry. We Will fix it soon")
	}

	if age < 0 || age > 120 {
		return nil, errors.New("age must be between 0 and 120")
	}

	return &Client{
		Name:   name,
		Gender: gender,
		Age:    age,
	}, nil
}
