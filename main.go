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

	dbConnection := postegresql.InitializeDatabase()
	clientRepo := repositories.NewClientSQLRepository(dbConnection)
	clientUseCase := usecases.NewCreateClient(clientRepo)
	clientController := controllers.NewCreateClientController(clientUseCase)
	clientHandler := handlers.NewCreateClientController(clientController)
	r.POST("/clients", clientHandler.GinHandler)

	r.Run()
}
