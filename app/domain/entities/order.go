package entities

import "time"

// Order represents an order
type Order struct {
	ID           int
	CustomerName *string
	PhoneNumber  *string
	TableID      int
	Items        []OrderItem
	CreatedAt    time.Time
	ClosedAt     *time.Time
}
