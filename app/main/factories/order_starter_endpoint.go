package factories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/application/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/usecases"
	"github.com/palexandremello/ramenshop-backend/app/infra/http/handlers"
	"github.com/palexandremello/ramenshop-backend/app/infra/repositories"
)

type OrderStarterEndpointFactory struct {
	dbConnection *sql.DB
}

// NewOrderStarterEndpointFactory is a factory function that creates a new OrderStarterEndpointFactory
func NewOrderStarterEndpointFactory(dbConnection *sql.DB) *OrderStarterEndpointFactory {
	return &OrderStarterEndpointFactory{dbConnection: dbConnection}
}

// CreateEndpoint is a factory method that creates a new OrderStarterHandler
func (f *OrderStarterEndpointFactory) CreateEndpoint() *handlers.OrderStarterHandler {
	dishRepository := repositories.NewDishSQLRepository(f.dbConnection)
	orderItemRepository := repositories.NewOrderItemSQLRepository(f.dbConnection)
	orderRepository := repositories.NewOrderSQLRepository(f.dbConnection)
	orderTableRepository := repositories.NewOrderTableSQLRepository(f.dbConnection)
	orderStarterUseCase := usecases.NewOrderStarterUseCase(orderRepository, orderItemRepository, dishRepository, orderTableRepository)
	orderStarterController := controllers.NewOrderStarterController(orderStarterUseCase)

	handlerInstance := handlers.NewOrderStarterHandler(orderStarterController)

	return handlerInstance
}
