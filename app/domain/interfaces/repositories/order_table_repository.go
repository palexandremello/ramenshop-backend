package repositories

type OrderTableRepositry interface {
	CreateAssociation(orderID int, tableID int) error
}
