package services

// SubscriberEvent is an interface that defines the contract for a subscriber event
type SubscriberEvent interface {
	Subscribe(channel string) (<-chan string, error)
	Unsubscribe(channel string) error
}
