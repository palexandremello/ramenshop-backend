package controllers

// CloseOrderController is the interface that wraps the basic CloseOrder method.
type CloseOrderController interface {
	Execute(OrderID int) error
}
