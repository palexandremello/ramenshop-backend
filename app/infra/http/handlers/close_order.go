package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
)

type CloseOrderHandler struct {
	controller controllers.CloseOrderController
}

type CloseOrderRequest struct {
	OrderID int `json:"order_id"`
}

func NewCloseOrderHandler(ctrl controllers.CloseOrderController) *CloseOrderHandler {
	return &CloseOrderHandler{controller: ctrl}
}

func (h *CloseOrderHandler) GinHandler(ctx *gin.Context) {
	req := &CloseOrderRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.controller.Execute(req.OrderID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order closed successfully"})
}
