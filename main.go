package main

import (
	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/application/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/infra/db/postegresql"
	"github.com/palexandremello/ramenshop-backend/app/domain/infra/http/handlers"
	"github.com/palexandremello/ramenshop-backend/app/domain/infra/repositories"
	"github.com/palexandremello/ramenshop-backend/app/domain/usecases"
)

func main() {
	r := gin.Default()

	// Database connection
	dbConnection := postegresql.InitializeDatabase()

	// Client setup
	clientRepo := repositories.NewClientSQLRepository(dbConnection)
	clientUseCase := usecases.NewCreateClient(clientRepo)
	clientController := controllers.NewCreateClientController(clientUseCase)
	clientHandler := handlers.NewCreateClientController(clientController)

	r.POST("/clients", clientHandler.GinHandler)

	// Table setup
	tableRepo := repositories.NewRegisterSQLRepository(dbConnection)
	tableUseCase := usecases.NewRegisterTable(tableRepo)
	tableController := controllers.NewRegisterTableController(tableUseCase)
	tableHandler := handlers.NewRegisterTableController(tableController)
	r.POST("/tables", tableHandler.GinHandler)

	updateTableUseCase := usecases.NewUpdateTableAvailabilty(tableRepo)
	updateTableController := controllers.NewUpdateTableAvailabityController(updateTableUseCase)
	updatedTableHandler := handlers.NewUpdateTableAvailabityController(updateTableController)
	r.PATCH("/tables", updatedTableHandler.GinHandler)
	r.Run()
}
