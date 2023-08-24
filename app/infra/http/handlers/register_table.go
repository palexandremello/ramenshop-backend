package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
)

type RegisterTableHandler struct {
	controller controllers.RegisterTableController
}

type TableRequest struct {
	Capacity int `json:"capacity"`
}

func NewRegisterTableController(ctrl controllers.RegisterTableController) *RegisterTableHandler {
	return &RegisterTableHandler{controller: ctrl}
}

func (h *RegisterTableHandler) GinHandler(ctx *gin.Context) {
	req := &TableRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTable, err := h.controller.Execute(req.Capacity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, createdTable)

}
