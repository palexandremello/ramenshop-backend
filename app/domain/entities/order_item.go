package entities

// OrderItem entity
type OrderItem struct {
	ID      int
	OrderID int
	Dish    Dish
	Amount  int
}
