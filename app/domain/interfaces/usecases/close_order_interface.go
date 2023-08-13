package usecases

// CloseOrder usecase
type CloseOrder interface {
	Execute(orderID int) error
}
