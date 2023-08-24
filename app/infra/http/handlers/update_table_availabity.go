package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
)

type UpdateTableAvailabityHandler struct {
	controller controllers.UpdateTableAvailabityController
}

type UpdateTableRequest struct {
	TableID     int  `json:"table_id"`
	IsAvailable bool `json:"is_available"`
}

func NewUpdateTableAvailabityController(ctrl controllers.UpdateTableAvailabityController) *UpdateTableAvailabityHandler {
	return &UpdateTableAvailabityHandler{controller: ctrl}
}

func (h *UpdateTableAvailabityHandler) GinHandler(ctx *gin.Context) {
	req := &UpdateTableRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(req)

	createdTable, err := h.controller.Execute(req.TableID, req.IsAvailable)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdTable)

}
