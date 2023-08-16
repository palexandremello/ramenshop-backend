package usecases

type UpdateTableAvailability interface {
	Execute(tableID int, isAvailable bool) error
}
