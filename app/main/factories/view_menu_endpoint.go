package factories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/application/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/usecases"
	"github.com/palexandremello/ramenshop-backend/app/infra/http/handlers"
	"github.com/palexandremello/ramenshop-backend/app/infra/repositories"
)

type ViewMenuEndpointFactory struct {
	dbConnection *sql.DB
}

// NewViewMenuEndpointFactory is a factory function that creates a new ViewMenuHandler
func NewViewMenuEndpointFactory(dbConnection *sql.DB) *ViewMenuEndpointFactory {
	return &ViewMenuEndpointFactory{dbConnection: dbConnection}
}

// CreateEndpoint is a factory method that creates a new ViewMenuHandler
func (f *ViewMenuEndpointFactory) CreateEndpoint() *handlers.ViewMenuHandler {
	dishRepository := repositories.NewDishSQLRepository(f.dbConnection)
	viewMenuUseCase := usecases.NewViewMenu(dishRepository)
	viewMenuController := controllers.NewViewMenuController(viewMenuUseCase)

	handlerInstance := handlers.NewViewMenuHandler(viewMenuController)

	return handlerInstance
}
