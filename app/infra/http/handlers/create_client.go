package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/entities"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
)

type CreateClientHandler struct {
	controller controllers.CreateClientController
}

func NewCreateClientController(ctrl controllers.CreateClientController) *CreateClientHandler {
	return &CreateClientHandler{controller: ctrl}
}

func (h *CreateClientHandler) GinHandler(ctx *gin.Context) {
	client := &entities.Client{}
	fmt.Println(client)
	if err := ctx.ShouldBindJSON(client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdClient, err := h.controller.Create(client)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdClient)
}
