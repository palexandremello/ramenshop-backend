package services

// PublisherEvent is an interface that defines the contract for a publisher event
type PublisherEvent interface {
	Execute(channel string, message string) error
}
