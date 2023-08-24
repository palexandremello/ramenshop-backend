package entities

// Dish representa o prato real que os clientes podem pedir.
type Dish struct {
	ID          int
	Name        string
	Description string
	Photo       *Photo
	Price       float64
	Available   bool
	Type        string
}
