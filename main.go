package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/palexandremello/ramenshop-backend/app/application/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/usecases"
	"github.com/palexandremello/ramenshop-backend/app/infra/db/postegresql"
	"github.com/palexandremello/ramenshop-backend/app/infra/http/handlers"
	"github.com/palexandremello/ramenshop-backend/app/infra/repositories"
	"github.com/palexandremello/ramenshop-backend/app/main/factories"
)

func main() {
	godotenv.Load()

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

	// Dishes Setup

	dishEndpointFactory := factories.NewDishEndpointFactory(dbConnection)
	createDishHandler := dishEndpointFactory.CreateEndpoint()
	r.POST("/dish", createDishHandler.GinHandler)

	dishRepository := repositories.NewDishSQLRepository(dbConnection)

	listDishUseCase := usecases.NewListDishUseCase(dishRepository)
	listDishController := controllers.NewListDishController(listDishUseCase)

	listDishHandler := handlers.NewListDishHandler(listDishController)
	r.GET("/dish", listDishHandler.GinHandler)

	orderStarterEndpoint := factories.NewOrderStarterEndpointFactory(dbConnection)
	orderStarterHandler := orderStarterEndpoint.CreateEndpoint()
	r.POST("/orders", orderStarterHandler.GinHandler)

	viewMenuEndpoint := factories.NewViewMenuEndpointFactory(dbConnection)
	viewMenuHandler := viewMenuEndpoint.CreateEndpoint()

	r.GET("/menu", viewMenuHandler.GinHandler)

	// Close Order Setup
	orderRepository := repositories.NewOrderSQLRepository(dbConnection)
	closeOrderUseCase := usecases.NewCloseOrder(orderRepository)
	closeOrderController := controllers.NewCloseOrderController(closeOrderUseCase)
	closeOrderHandler := handlers.NewCloseOrderHandler(closeOrderController)
	r.PATCH("/orders", closeOrderHandler.GinHandler)

	r.Run()
}
