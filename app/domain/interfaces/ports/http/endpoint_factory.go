package http

import "github.com/palexandremello/ramenshop-backend/app/domain/interfaces/ports"

// EndpointFactory is a interface that defines a method to create a endpoint
type EndpointFactory interface {
	CreateEndpoint() ports.Handler
}
