package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
)

// ListDishHandler is a struct that defines the ListDishHandler
type ListDishHandler struct {
	controller controllers.ListDishController
}

// NewListDishHandler is a factory function that creates a new ListDishHandler
func NewListDishHandler(ctrl controllers.ListDishController) *ListDishHandler {
	return &ListDishHandler{controller: ctrl}
}

// GinHandler is a function that handles the request and returns the response
func (h *ListDishHandler) GinHandler(ctx *gin.Context) {
	// Extrair parâmetros de consulta
	dishType := ctx.DefaultQuery("dishType", "") // retorna uma string vazia se o parâmetro dishType não estiver presente
	dishIDStr := ctx.DefaultQuery("dishID", "-1")
	dishID, _ := strconv.Atoi(dishIDStr)

	// Chame o controller com os parâmetros
	dishes, err := h.controller.Execute(dishType, dishID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Se tudo estiver bem, envie a lista de pratos como resposta
	ctx.JSON(http.StatusOK, dishes)
}
