package entities

import "time"

// OrderStatus represents the possible states an order can be in.
type OrderStatus string

const (
	Open      OrderStatus = "ABERTO"
	InProcess OrderStatus = "EM PREPARO"
	Closed    OrderStatus = "FECHADO"
)

// Order represents an order
type Order struct {
	ID           int
	CustomerName *string
	PhoneNumber  *string
	TableID      int
	Items        []OrderItem
	Status       OrderStatus
	CreatedAt    time.Time
	ClosedAt     *time.Time
}
