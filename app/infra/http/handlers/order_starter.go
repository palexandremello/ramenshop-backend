package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/usecases"
)

type OrderStarterHandler struct {
	controller controllers.OrderStarterController
}

type OrderStarterRequest struct {
	CustomerName string `json:"customer_name"`
	TableID      int    `json:"table_id"`
	OrderItems   []struct {
		DishID int `json:"dish_id"`
		Amount int `json:"amount"`
	} `json:"orders_item"`
}

// NewOrderStarterHandler creates a new instance of OrderStarterHandler
func NewOrderStarterHandler(controller controllers.OrderStarterController) *OrderStarterHandler {
	return &OrderStarterHandler{controller: controller}
}

// GinHandler handles the request
func (h *OrderStarterHandler) GinHandler(ctx *gin.Context) {
	req := &OrderStarterRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se OrderItems está vazio
	if len(req.OrderItems) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "OrderItems should not be empty"})
		return
	}

	dishInputs := make([]usecases.DishOrderInput, len(req.OrderItems))

	for i, item := range req.OrderItems {
		dishInputs[i] = usecases.DishOrderInput{
			DishID: item.DishID,
			Amount: item.Amount,
		}
	}

	order, err := h.controller.Execute(&req.CustomerName, req.TableID, dishInputs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}
