package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
)

type ViewMenuHandler struct {
	controller controllers.ViewMenuController
}

// NewViewMenuHandler is a factory function that creates a new ViewMenuHandler
func NewViewMenuHandler(controller controllers.ViewMenuController) *ViewMenuHandler {
	return &ViewMenuHandler{controller: controller}
}

// GinHandler is a function that handles the request and returns a response
func (h *ViewMenuHandler) GinHandler(ctx *gin.Context) {
	order, err := h.controller.Execute()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}
