package factories

import (
	"database/sql"

	"github.com/palexandremello/ramenshop-backend/app/application/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/usecases"
	"github.com/palexandremello/ramenshop-backend/app/infra/http/handlers"
	"github.com/palexandremello/ramenshop-backend/app/infra/repositories"
	"github.com/palexandremello/ramenshop-backend/app/infra/services"
)

// DishEndpointFactory is a factory that creates a new CreateDishHandler
type DishEndpointFactory struct {
	dbConnection *sql.DB
}

// NewDishEndpointFactory is a factory function that creates a new DishEndpointFactory
func NewDishEndpointFactory(dbConnection *sql.DB) *DishEndpointFactory {
	return &DishEndpointFactory{dbConnection: dbConnection}
}

// CreateEndpoint is a factory method that creates a new CreateDishHandler
func (f *DishEndpointFactory) CreateEndpoint() *handlers.CreateDishHandler {

	dishRepository := repositories.NewDishSQLRepository(f.dbConnection)
	photoRepository := repositories.NewPhotoSQLRepository(f.dbConnection)
	digitalOceanSpaceFileUploader := services.NewDigitalOceanSpacesFileUploader()
	httpService := services.NewHTTPService()
	photoUseCase := usecases.NewPhotoUseCase(photoRepository, httpService, digitalOceanSpaceFileUploader)
	dishesUseCase := usecases.NewDishUseCase(dishRepository, photoUseCase)

	createDishController := controllers.NewCreateDishController(dishesUseCase)

	handlerInstance := handlers.NewCreateDishHandler(createDishController)

	return handlerInstance

}
