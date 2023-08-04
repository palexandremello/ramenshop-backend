package entities

import "time"

// Order represents an order
type Order struct {
	ID        int
	Client    Client
	Items     []OrderItem
	CreatedAt time.Time
}
